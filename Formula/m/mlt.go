
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mltFormula represents a formula in Go.
type mltFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mltFormula) Print() {
    fmt.Printf("Name: mlt\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mltFormula{
        Description:  "Author, manage, and run multitrack audio/video compositions",
        Homepage:     "https://www.mltframework.org/",
        URL:          "https://github.com/mltframework/mlt/releases/download/v7.26.0/mlt-7.26.0.tar.gz",
        Sha256:       "f8696cbec063d6ba3f1db543e29e09c40566ccafd286c5d8bfee9f6decee3dcf",
        Dependencies: []string{"cmake", "pkg-config", "ffmpeg", "fftw", "fontconfig", "frei0r", "gdk-pixbuf", "glib", "libdv", "libexif", "libsamplerate", "libvidstab", "libvorbis", "opencv", "pango", "qt", "rubberband", "sdl2", "sox", "freetype", "gettext", "harfbuzz", "alsa-lib", "pulseaudio"},
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

    fmt.Printf("Installing formula: %s\n", "mlt")
    if err := pkg.Installmlt(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mltFormula) Installmlt() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "mlt-7.26.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "mlt-7.26.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
