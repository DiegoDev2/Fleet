
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// popplerFormula represents a formula in Go.
type popplerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg popplerFormula) Print() {
    fmt.Printf("Name: poppler\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := popplerFormula{
        Description:  "PDF rendering library (based on the xpdf-3.0 code base)",
        Homepage:     "https://poppler.freedesktop.org/",
        URL:          "https://poppler.freedesktop.org/poppler-data-0.4.12.tar.gz",
        Sha256:       "c835b640a40ce357e1b83666aabd95edffa24ddddd49b8daff63adb851cdab74",
        Dependencies: []string{"cmake", "gobject-introspection", "pkg-config", "cairo", "fontconfig", "freetype", "gettext", "glib", "gpgme", "jpeg-turbo", "libpng", "libtiff", "little-cms2", "nspr", "nss", "openjpeg", "libassuan"},
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

    fmt.Printf("Installing formula: %s\n", "poppler")
    if err := pkg.Installpoppler(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg popplerFormula) Installpoppler() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "poppler-data-0.4.12.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "poppler-data-0.4.12.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
