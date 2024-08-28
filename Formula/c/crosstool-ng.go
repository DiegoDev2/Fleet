
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// crosstool-ngFormula represents a formula in Go.
type crosstool-ngFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg crosstool-ngFormula) Print() {
    fmt.Printf("Name: crosstool-ng\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := crosstool-ngFormula{
        Description:  "Tool for building toolchains",
        Homepage:     "https://crosstool-ng.github.io/",
        URL:          "https://crosstool-ng.github.io/download/",
        Sha256:       "a309a9662fc4b2f991e21fde07dd0e4b552829d36add35543991cf77ea7d6b3d",
        Dependencies: []string{"help2man", "autoconf", "automake", "binutils", "bison", "flex", "gettext", "libtool", "lzip", "m4", "ncurses", "python@3.12", "xz", "bash", "coreutils", "gawk", "gnu-sed", "grep", "make", "texinfo"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installcrosstool-ng(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg crosstool-ngFormula) Installcrosstool-ng() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := ""
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := ""
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
