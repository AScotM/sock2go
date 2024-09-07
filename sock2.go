package main

import (
	"fmt"
	"os/exec"
	"log"
)

func getSocketStatistics() (string, error) {
	cmd := exec.Command("ss", "-s")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func printSocketStatistics() error {
	socketStats, err := getSocketStatistics()
	if err != nil {
		return err
	}

	// Print statistics
	fmt.Println("Socket Statistics:")
	fmt.Println(socketStats)
	return nil
}

func main() {
	err := printSocketStatistics()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

