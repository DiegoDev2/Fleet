
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pulseaudioFormula represents a formula in Go.
type pulseaudioFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pulseaudioFormula) Print() {
    fmt.Printf("Name: pulseaudio\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pulseaudioFormula{
        Description:  "Sound system for POSIX OSes",
        Homepage:     "https://wiki.freedesktop.org/www/Software/PulseAudio/",
        URL:          "https://www.freedesktop.org/software/pulseaudio/releases/pulseaudio-17.0.tar.xz",
        Sha256:       "504035dfda3bffabae42f352d0e7c0a90c6b8c1f6b925fe7e17502124d5d6529",
        Dependencies: []string{"gettext", "meson", "ninja", "pkg-config", "glib", "libsndfile", "libsoxr", "libtool", "openssl@3", "orc", "speexdsp", "gettext", "perl-xml-parser", "alsa-lib", "dbus", "libcap"},
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

    fmt.Printf("Installing formula: %s\n", "pulseaudio")
    if err := pkg.Installpulseaudio(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pulseaudioFormula) Installpulseaudio() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "pulseaudio-17.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "pulseaudio-17.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
