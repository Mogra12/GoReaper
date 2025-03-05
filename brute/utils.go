package brute

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Red   = "\033[31m" // Color code for red
	Green = "\033[32m" // Color code for green
	Reset = "\033[0m"  // Reset color code
)

// isEmpty checks if a file is empty
func isEmpty(path string) (bool, error) {
	// Get the file information
	info, err := os.Stat(path)
	if err != nil {
		return false, err // Return error if file access fails
	}
	// Check if the file size is 0, indicating it's empty
	return info.Size() == 0, nil
}

// WlLoader loads the wordlist from the file at the given path
func WlLoader(wordlistPath string) ([]string, bool) {
	// Open the file specified by wordlistPath
	file, err := os.Open(wordlistPath)
	if err != nil {
		log.Fatal("File does not exist or invalid path!") // Log fatal error if file doesn't exist
	}
	defer file.Close() // Ensure the file is closed after reading

	// Declare a slice to store the wordlist
	var Wordlist []string
	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line from the file to the Wordlist slice
		Wordlist = append(Wordlist, scanner.Text())
	}

	// Check if the file is empty by calling isEmpty
	empty, err := isEmpty(wordlistPath)
	if err != nil {
		fmt.Printf("Error checking file: %v\n", err) // Print error if isEmpty fails
	}

	// Return the wordlist and whether the file was empty
	return Wordlist, empty
}

func Help() {
	fmt.Printf(" <----- %sFTP Commands%s -----> \n", Green, Reset)
	fmt.Printf("%s  -Cn     %sMaximum number of concurrent connections\n", Green, Reset)
	fmt.Printf("%s  -tls    %sUses FTPS (FTP over TLS)\n", Green, Reset)
	fmt.Printf("%s  -w      %sPath to the wordlist file\n", Green, Reset)
	fmt.Printf("%s  -t      %sFTP target server address (e.g., hostname:21)\n", Green, Reset)
	fmt.Printf("%s  -time   %sNumber of seconds to sleep between login attempts when lwr is true\n", Green, Reset)
	
	fmt.Printf(" <----- %sSSH Commands%s -----> \n", Green, Reset)
	fmt.Printf("%s  -U      %sPath to the users wordlist file\n", Green, Reset)
	fmt.Printf("%s  -P      %sPath to the password wordlist file\n", Green, Reset)
	fmt.Printf("%s  -T      %sSSH target server address (e.g., target:22)\n", Green, Reset)
}
