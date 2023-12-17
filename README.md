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
```

# Intrusion Detection Program

This Go program analyzes log files for potential intrusions based on failed login attempts within a specified time window.

## Overview

The program `intrusion_detection` reads a log file (`sample_log.txt` by default) containing login attempts and identifies potential intrusions based on failed login attempts from the same IP address within a specified time window.

## Usage

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

This project is licensed under the Apache-2.0 License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This program is a simple illustration of log file analysis for intrusion detection using Go.

---

**Documentation By:** Raymond C. TURNER

**Revision:** Sunday 17th December 2023

codestak.io
