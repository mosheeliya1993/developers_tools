package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "darwin":
		setupMac()
	case "windows":
		setupWindows()
	case "linux":
		setupLinux()
	default:
		fmt.Println("Unsupported operating system")
	}
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
}

func setupMac() {
	// Homebrew installation
	runCommand("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")

	// Install dependencies using brew
	runCommand("brew", "install", "kubectl")
	runCommand("brew", "install", "minikube")
	runCommand("brew", "install", "helm")
}

// setupWindows sets up the necessary dependencies for running Kubernetes on a Windows machine.
func setupWindows() {
	// define a slice of commands to be executed
	commands := []struct {
		cmd  string
		args []string
	}{
		// set the execution policy to bypass for the current process
		{"Set-ExecutionPolicy", []string{"Bypass", "-Scope", "Process", "-Force"}},
		// execute the PowerShell script to install Chocolatey package manager
		{"iex", []string{"((Invoke-WebRequest -useb https://chocolatey.org/install.ps1).Content)"}},
		// install the Kubernetes CLI using Chocolatey
		{"choco", []string{"install", "kubernetes-cli"}},
		// install Minikube using Chocolatey
		{"choco", []string{"install", "minikube"}},
		// install Helm using Chocolatey
		{"choco", []string{"install", "helm"}},
	}
	// iterate over each command and run it
	for _, command := range commands {
		runCommand(command.cmd, command.args...)
	}
}
func setupLinux() {
	// Update and install snap (for Ubuntu)
	runCommand("sudo", "apt", "update")
	runCommand("sudo", "apt", "install", "snapd")

	// Install dependencies
	runCommand("snap", "install", "kubectl", "--classic")
	runCommand("snap", "install", "minikube", "--classic")
	runCommand("snap", "install", "helm", "--classic")
}
