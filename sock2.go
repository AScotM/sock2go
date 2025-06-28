package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

// CommandRunner interface for testability
type CommandRunner interface {
	LookPath(string) (string, error)
	CombinedOutput(string, ...string) ([]byte, error)
}

// RealCmdRunner implements CommandRunner for actual system commands
type RealCmdRunner struct{}

func (r RealCmdRunner) LookPath(cmd string) (string, error) {
	return exec.LookPath(cmd)
}

func (r RealCmdRunner) CombinedOutput(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}

// getSocketStatistics gets socket stats using the CommandRunner interface
func getSocketStatistics(runner CommandRunner, timeout time.Duration) (string, error) {
	// Check if command exists
	if _, err := runner.LookPath("ss"); err != nil {
		return "", fmt.Errorf("ss command not found (Linux-only): %w", err)
	}

	// Set up timeout context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Execute command with context
	output, err := runner.CombinedOutput("ss", "-tanl")
	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("ss command timed out after %v", timeout)
	}
	if err != nil {
		return "", fmt.Errorf("command execution failed: %w\nOutput: %s", err, string(output))
	}

	// Validate output
	if len(output) == 0 {
		return "", fmt.Errorf("empty output from ss command")
	}

	return string(output), nil
}

// printSocketStatistics handles the display logic
func printSocketStatistics(runner CommandRunner, timeout time.Duration) error {
	stats, err := getSocketStatistics(runner, timeout)
	if err != nil {
		return fmt.Errorf("failed to get socket statistics: %w", err)
	}

	fmt.Printf("Socket Statistics (collected at %s):\n%s", 
		time.Now().Format(time.RFC3339), stats)
	return nil
}

func main() {
	const defaultTimeout = 5 * time.Second
	
	// Initialize with real command runner
	runner := RealCmdRunner{}
	
	// Execute and handle errors
	if err := printSocketStatistics(runner, defaultTimeout); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
