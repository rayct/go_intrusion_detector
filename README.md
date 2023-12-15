#  Basic Log Intrusion Detector in Go

This simple Go program reads a log file and detects potential intrusions based on login attempts.

## Overview

The program scans a log file for failed login attempts and identifies potential intruders by tracking the number of login failures within a specified time window for each IP address.

## Usage

### Setup

1. Ensure you have Go installed on your system.
2. Clone or download this repository.

### Running the Program

1. Open a terminal and navigate to the project directory.
2. Modify the `sample_log.txt` file or provide your log file with the relevant data.
3. Run the program using the following command:
   ```bash
   go run main.go


```markdown
# Intrusion Detection Program

This Go program analyzes log files for potential intrusions based on failed login attempts within a specified time window.

## Overview

The program `intrusion_detection.go` reads a log file (`sample_log.txt` by default) containing login attempts and identifies potential intrusions based on failed login attempts from the same IP address within a specified time window.

## Usage

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readLogFile(logFile string) ([]string, error) {
	file, err := os.Open(logFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logEntries []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logEntries = append(logEntries, scanner.Text())
	}
	return logEntries, scanner.Err()
}

func detectIntrusion(logEntries []string, maxAttempts int, timeWindow time.Duration) (bool, string) {
	ipAttempts := make(map[string][]time.Time)

	for _, entry := range logEntries {
		parts := strings.Fields(entry)
		if len(parts) >= 5 && parts[3] == "failed" && parts[4] == "login" && parts[5] == "attempt" {
			timestampStr := parts[0] + " " + parts[1]
			timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
			if err != nil {
				continue
			}

			ipAddress := parts[2]
			if attempts, ok := ipAttempts[ipAddress]; ok {
				if time.Since(attempts[len(attempts)-1]) <= timeWindow {
					ipAttempts[ipAddress] = append(ipAttempts[ipAddress], timestamp)
					if len(ipAttempts[ipAddress]) >= maxAttempts {
						return true, ipAddress
					}
				} else {
					ipAttempts[ipAddress] = []time.Time{timestamp}
				}
			} else {
				ipAttempts[ipAddress] = []time.Time{timestamp}
			}
		}
	}

	return false, ""
}

func main() {
	logFile := "sample_log.txt"
	logEntries, err := readLogFile(logFile)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	intrusionDetected, suspiciousIP := detectIntrusion(logEntries, 3, 60*time.Second)

	if intrusionDetected {
		fmt.Printf("Intrusion detected from %s.\n", suspiciousIP)
	} else {
		fmt.Println("No intrusion detected.")
	}
}

```

```markdown
### Prerequisites

- Go programming language installed on your system.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/intrusion-detection.git
   ```

2. Navigate to the project directory:
   ```bash
   cd intrusion-detection
   ```

3. Run the program:
   ```bash
   go run main.go
   ```
   This command executes the program with the default log file `sample_log.txt`.

### Customization

- To use a different log file, modify the `logFile` variable in the `main()` function.
- Adjust the `maxAttempts` and `timeWindow` parameters in the `detectIntrusion()` function to change the threshold for identifying intrusions.

## File Structure

- `intrusion_detection.go`: Contains the main program logic for intrusion detection.
- `sample_log.txt`: Sample log file for demonstration purposes.

## Sample Code Explanation

The program consists of the following key functionalities:

- **readLogFile**: Reads the contents of a log file into a slice of strings.
- **detectIntrusion**: Analyzes log entries to detect potential intrusions based on failed login attempts within a specified time window.
- **main**: Entry point of the program, reads the log file, and calls `detectIntrusion` to identify potential intrusions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This program is a simple illustration of log file analysis for intrusion detection using Go.

---

**Documentation By:** Raymond C. TURNER

**Revision:** Friday 15th December 2023

codestak.io
