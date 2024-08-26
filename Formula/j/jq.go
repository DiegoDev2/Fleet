package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Formula struct {
	Name        string
	Description string
	Homepage    string
	URL         string
	Sha256      string
	License     string
	Install     func() error
	Test        func() error
}

func (f *Formula) InstallPackage() error {
	fmt.Printf("Installing %s...\n", f.Name)
	if err := f.Install(); err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}
	return nil
}

func (f *Formula) TestPackage() error {
	fmt.Printf("Testing %s...\n", f.Name)
	if err := f.Test(); err != nil {
		return fmt.Errorf("testing failed: %v", err)
	}
	return nil
}

var kubernetes = &Formula{
	Name:        "kubernetes",
	Description: "Kubernetes container orchestration platform",
	Homepage:    "https://kubernetes.io/",
	URL:         "https://dl.k8s.io/v1.26.0/kubernetes-client-linux-amd64.tar.gz",
	Sha256:      "b8ffbb97e0c8a7a0d2c1f0ff3c6a6d97b5768a0c3f1d12d51ac1f55aeb1c5a8d",
	License:     "Apache-2.0",
	Install: func() error {
		fmt.Println("Downloading Kubernetes client...")
		cmd := exec.Command("curl", "-LO", "https://dl.k8s.io/v1.26.0/kubernetes-client-linux-amd64.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Kubernetes client...")
		cmd = exec.Command("tar", "-xzf", "kubernetes-client-linux-amd64.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Kubernetes client...")
		cmd = exec.Command("mv", "kubernetes/client/bin/kubectl", "/usr/local/bin/kubectl")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Kubernetes client...")
		cmd := exec.Command("kubectl", "version", "--client")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := kubernetes.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := kubernetes.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Kubernetes client installed and tested successfully!")
}
