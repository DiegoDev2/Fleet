
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sync_gatewayFormula represents a formula in Go.
type sync_gatewayFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sync_gatewayFormula) Print() {
    fmt.Printf("Name: sync_gateway\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sync_gatewayFormula{
        Description:  "Make Couchbase Server a replication endpoint for Couchbase Lite",
        Homepage:     "https://docs.couchbase.com/sync-gateway/current/index.html",
        URL:          "https://github.com/couchbase/sync_gateway/commit/97279d5ff172c1535bd4df764fbc51029d003b53.patch?full_index=1",
        Sha256:       "0e32ee39b28d597b25f950b27cfb14b306db52fd752f855cab27d0a994271277",
        Dependencies: []string{"gnupg", "go", "python@3.11", "repo"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        if !isFormulaInstalled(dep) {
            fmt.Printf("Installing dependency: %s\n", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("Dependency %s is already installed.\n", dep)
        }
    }

    fmt.Printf("Installing formula: %s\n", "sync_gateway")
    if err := pkg.Installsync_gateway(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg sync_gatewayFormula) Installsync_gateway() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "97279d5ff172c1535bd4df764fbc51029d003b53.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "97279d5ff172c1535bd4df764fbc51029d003b53"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
