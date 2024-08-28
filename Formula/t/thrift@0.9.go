
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// thrift@0.9Formula represents a formula in Go.
type thrift@0.9Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg thrift@0.9Formula) Print() {
    fmt.Printf("Name: thrift@0.9\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := thrift@0.9Formula{
        Description:  "Framework for scalable cross-language services development",
        Homepage:     "https://thrift.apache.org",
        URL:          "https://github.com/apache/thrift/commit/b819260c653f6fd9602419ee2541060ecb930c4c.patch?full_index=1",
        Sha256:       "5934555674b67fb7a9fad04ffe0bd46fdbe3eca5e8f98dd072efa4bb342c9bfa",
        Dependencies: []string{"autoconf", "automake", "bison", "libtool", "openssl@3", "pkg-config", "boost"},
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

    fmt.Printf("Installing formula: %s\n", "thrift@0.9")
    if err := pkg.Installthrift@0.9(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg thrift@0.9Formula) Installthrift@0.9() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "b819260c653f6fd9602419ee2541060ecb930c4c.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "b819260c653f6fd9602419ee2541060ecb930c4c"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
