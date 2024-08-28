
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// syslog-ngFormula represents a formula in Go.
type syslog-ngFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg syslog-ngFormula) Print() {
    fmt.Printf("Name: syslog-ng\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := syslog-ngFormula{
        Description:  "Log daemon with advanced processing pipeline and a wide range of I/O methods",
        Homepage:     "https://www.syslog-ng.com",
        URL:          "https://github.com/syslog-ng/syslog-ng/releases/download/syslog-ng-4.8.0/syslog-ng-4.8.0.tar.gz",
        Sha256:       "95613dc27d4a988e69e01c9d0fcd36fca5150a0829673ac1718dac59e2e21571",
        Dependencies: []string{"pkg-config", "abseil", "glib", "grpc", "hiredis", "ivykis", "json-c", "libdbi", "libmaxminddb", "libnet", "libpaho-mqtt", "librdkafka", "mongo-c-driver", "net-snmp", "openssl@3", "pcre2", "protobuf", "python@3.12", "rabbitmq-c", "riemann-client", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "syslog-ng")
    if err := pkg.Installsyslog-ng(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg syslog-ngFormula) Installsyslog-ng() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "syslog-ng-4.8.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "syslog-ng-4.8.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
