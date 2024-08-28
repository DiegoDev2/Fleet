
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// eflFormula represents a formula in Go.
type eflFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg eflFormula) Print() {
    fmt.Printf("Name: efl\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := eflFormula{
        Description:  "Enlightenment Foundation Libraries",
        Homepage:     "https://www.enlightenment.org",
        URL:          "https://download.enlightenment.org/rel/libs/efl/",
        Sha256:       "1d265348e67f8f8691f74219ba1b7b03a6ad1fcc8b99663225cca5a49b3d5ecc",
        Dependencies: []string{"meson", "ninja", "pkg-config", "bullet", "cairo", "dbus", "fontconfig", "freetype", "fribidi", "gettext", "giflib", "glib", "gstreamer", "harfbuzz", "jpeg-turbo", "libpng", "libraw", "librsvg", "libsndfile", "libspectre", "libtiff", "luajit", "lz4", "openjpeg", "openssl@3", "poppler", "pulseaudio", "shared-mime-info", "webp", "gdk-pixbuf", "little-cms2", "mesa"},
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

    fmt.Printf("Installing formula: %s\n", "efl")
    if err := pkg.Installefl(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg eflFormula) Installefl() error {
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
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
