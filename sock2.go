package main

import (
    "fmt"
    "log"
    "os/exec"
)

func getSocketStatistics() (string, error) {
    // Check if the "ss" command is available
    if _, err := exec.LookPath("ss"); err != nil {
        return "", fmt.Errorf("ss command not found: %v", err)
    }

    // Execute the "ss -s" command
    cmd := exec.Command("ss", "-s")
    output, err := cmd.CombinedOutput() // Captures both stdout and stderr
    if err != nil {
        return "", fmt.Errorf("failed to execute ss command: %v, output: %s", err, string(output))
    }

    return string(output), nil
}

func printSocketStatistics() error {
    socketStats, err := getSocketStatistics()
    if err != nil {
        return err // Logging will be handled in main
    }

    fmt.Println("Socket Statistics:\n" + socketStats)
    return nil
}

func main() {
    if err := printSocketStatistics(); err != nil {
        log.Fatalf("Failed to get socket statistics: %v", err) // Terminates on failure
    }
}
