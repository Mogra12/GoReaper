# GoReaper

GoReaper is a brute-force tool developed in Go to test FTP and SSH credentials. It supports concurrent attacks and can use TLS for secure connections.

## Features

- Support for brute-force attacks on FTP and SSH.
- Concurrent execution with configurable connection limits.
- Option for FTPS (FTP over TLS) connection.
- Flexible configuration via command line.
- Support for separate user and password lists in SSH attack.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/Mogra12/GoReaper.git
   cd GoReaper
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Compile the program:
   ```bash
   go build -o goreaper main.go
   ```

## Usage

### FTP Attack

```bash
goreaper -Cn 10 -tls -w wordlist.txt -t target:21 -time 2
```

#### FTP Parameters:
- `-Cn`   : Set the maximum number of simultaneous connections.
- `-tls`  : Use FTPS (FTP over TLS).
- `-w`    : Path to the wordlist file.
- `-t`    : FTP server address (e.g., hostname:21).
- `-time` : Timeout in seconds between login attempts (when `lwr` is active).

### SSH Attack

```bash
goreaper -U users.txt -P passwords.txt -t target:22
```

#### SSH Parameters:
- `-U` : Path to the user wordlist.
- `-P` : Path to the password wordlist.
- `-T` : SSH server address (e.g., target:22).

## Contribution

Feel free to open issues and submit PRs for improvements!

## Legal Disclaimer

**This tool should only be used for security testing in authorized environments. Misuse may result in legal penalties.**
