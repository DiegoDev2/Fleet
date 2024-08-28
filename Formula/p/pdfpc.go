
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pdfpcFormula represents a formula in Go.
type pdfpcFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pdfpcFormula) Print() {
    fmt.Printf("Name: pdfpc\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pdfpcFormula{
        Description:  "Presenter console with multi-monitor support for PDF files",
        Homepage:     "https://pdfpc.github.io/",
        URL:          "https://github.com/pdfpc/pdfpc/commit/18beaecbbcc066e0d4c889b3aa3ecaa7351f7768.patch?full_index=1",
        Sha256:       "e516f2299f19c3e6ab4f200e4fcef85508655fe7890331e741a7303f9bc1a528",
        Dependencies: []string{"cmake", "vala", "cairo", "discount", "gdk-pixbuf", "glib", "gstreamer", "gtk+3", "json-glib", "libgee", "librsvg", "libx11", "pango", "poppler", "at-spi2-core", "gettext", "harfbuzz", "webkitgtk"},
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

    fmt.Printf("Installing formula: %s\n", "pdfpc")
    if err := pkg.Installpdfpc(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pdfpcFormula) Installpdfpc() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "18beaecbbcc066e0d4c889b3aa3ecaa7351f7768.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "18beaecbbcc066e0d4c889b3aa3ecaa7351f7768"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
