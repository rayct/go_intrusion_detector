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
