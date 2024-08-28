
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mkvtoolnixFormula represents a formula in Go.
type mkvtoolnixFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mkvtoolnixFormula) Print() {
    fmt.Printf("Name: mkvtoolnix\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mkvtoolnixFormula{
        Description:  "Matroska media files manipulation tools",
        Homepage:     "https://mkvtoolnix.download/",
        URL:          "https://gitlab.com/mbunkus/mkvtoolnix.git",
        Sha256:       "21d2ed6afb62cb9bfae54e58c0eaa4ee3755a181dbe09ee8022ab686931e7dbd",
        Dependencies: []string{"autoconf", "automake", "libtool", "docbook-xsl", "pkg-config", "boost", "flac", "fmt", "gettext", "gmp", "libebml", "libmatroska", "libogg", "libvorbis", "nlohmann-json", "pugixml", "qt", "utf8cpp"},
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

    fmt.Printf("Installing formula: %s\n", "mkvtoolnix")
    if err := pkg.Installmkvtoolnix(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mkvtoolnixFormula) Installmkvtoolnix() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "mkvtoolnix.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "mkvtoolnix"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
