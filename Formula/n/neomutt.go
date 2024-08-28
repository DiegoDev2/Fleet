
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// neomuttFormula represents a formula in Go.
type neomuttFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg neomuttFormula) Print() {
    fmt.Printf("Name: neomutt\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := neomuttFormula{
        Description:  "E-mail reader with support for Notmuch, NNTP and much more",
        Homepage:     "https://neomutt.org/",
        URL:          "https://github.com/neomutt/neomutt/archive/refs/tags/20240425.tar.gz",
        Sha256:       "e7647714a64ec9e65bcc014557938039aea02bcb435b24344bf30577840412b6",
        Dependencies: []string{"docbook-xsl", "pkg-config", "tcl-tk", "gettext", "gpgme", "libidn2", "lmdb", "lua", "ncurses", "notmuch", "openssl@3", "pcre2", "sqlite", "tokyo-cabinet", "libgpg-error", "libiconv"},
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

    fmt.Printf("Installing formula: %s\n", "neomutt")
    if err := pkg.Installneomutt(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg neomuttFormula) Installneomutt() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "20240425.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "20240425.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
