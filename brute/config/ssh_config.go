package brute

import "flag"

// FtpConfig holds the configuration for FTP brute force attack
type SSHConfig struct {
	UserWordlist string // Path to the wordlist file
	PassWordlist string
	Target       string // SSH server address (e.g., hostname)
	BrtOption 	 string
	HHelp     	 bool
}

// FtpLoadConfig loads the configuration values from command-line flags
func SshConfig() *SSHConfig {
	var SSHConfig SSHConfig
	flag.StringVar(&SSHConfig.UserWordlist, "U", "", "Path to the users wordlist file")
	flag.StringVar(&SSHConfig.PassWordlist, "P", "", "Path to the password wordlist file")
	flag.StringVar(&SSHConfig.Target, "T", "", "SSH target server address (e.g., target:22)")
	flag.StringVar(&SSHConfig.BrtOption, "pr", "", "Protocol (ssh or ftp)")
	flag.BoolVar(&SSHConfig.HHelp, "h", false, "Help command")
	flag.Parse()
	// Return the populated FtpConfig struct
	return &SSHConfig
}
