package main

import (
    "fmt"
    "os/exec"
    "log"
)

func getSocketStatistics() (string, error) {
    // Check if ss command is available
    _, err := exec.LookPath("ss")
    if err != nil {
        return "", fmt.Errorf("ss command not found: %v", err)
    }

    cmd := exec.Command("ss", "-s")
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to execute ss command: %v", err)
    }
    return string(output), nil
}

func printSocketStatistics() error {
    socketStats, err := getSocketStatistics()
    if err != nil {
        return err // Error will be logged in main
    }

    fmt.Println("Socket Statistics:")
    fmt.Println(socketStats)
    return nil
}

func main() {
    if err := printSocketStatistics(); err != nil {
        log.Printf("Failed to get socket statistics: %v\n", err)
    }
}
