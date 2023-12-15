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

---

**Documentation By:** Raymond C. TURNER

**Revision:** Friday 15th December 2023

codestak.io
