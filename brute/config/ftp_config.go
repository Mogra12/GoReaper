package brute

import "flag"

// FtpConfig holds the configuration for FTP brute force attack
type FtpConfig struct {
	MaxConcurrent int    // Maximum number of concurrent connections
	SleepDuration int    // Duration in seconds to sleep between login attempts (if lwr is true)
	WordlistPath  string // Path to the wordlist file
	Target        string   // SSH server address (e.g., hostname:21)
	TlsFlag       bool   // Flag to indicate if FTPS (FTP over TLS) should be used
}

// FtpLoadConfig loads the configuration values from command-line flags
func FtpLoadConfig() *FtpConfig {
	var FtpConfig FtpConfig
	
	// Bind command-line flags to the FtpConfig struct fields
	flag.IntVar(&FtpConfig.MaxConcurrent, "Cn", 5, "Maximum number of concurrent connections") // Number of simultaneous connections
	flag.BoolVar(&FtpConfig.TlsFlag, "tls", false, "Uses FTPS (FTP over TLS)")            // Whether to use FTPS
	flag.StringVar(&FtpConfig.WordlistPath, "w", "/wordlists/rockyou.txt", "Path to the wordlist file")                  // Path to the wordlist file
	flag.StringVar(&FtpConfig.Target, "t", "", "FTP target server address (e.g., hostname:21)")          // FTP server address
	flag.IntVar(&FtpConfig.SleepDuration, "time", 1, "Number of seconds to sleep between login attempts when lwr is true") // Sleep duration between login attempts
	flag.Parse() // Parse the command-line flags
	
	// Return the populated FtpConfig struct
	return &FtpConfig
}
