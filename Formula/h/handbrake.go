
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// handbrakeFormula represents a formula in Go.
type handbrakeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg handbrakeFormula) Print() {
    fmt.Printf("Name: handbrake\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := handbrakeFormula{
        Description:  "Open-source video transcoder available for Linux, Mac, and Windows",
        Homepage:     "https://handbrake.fr/",
        URL:          "https://github.com/HandBrake/HandBrake/releases/download/1.7.3/HandBrake-1.7.3-source.tar.bz2",
        Sha256:       "99e46b6be919867b917eb6db71151c02b601712d7e4403ea60aaa22d154d2db3",
        Dependencies: []string{"autoconf", "automake", "cmake", "libtool", "meson", "nasm", "ninja", "pkg-config", "10.3", "yasm", "jansson", "jpeg-turbo", "lame", "libass", "libvorbis", "libvpx", "numactl", "opus", "speex", "theora", "x264", "xz"},
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

    fmt.Printf("Installing formula: %s\n", "handbrake")
    if err := pkg.Installhandbrake(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg handbrakeFormula) Installhandbrake() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "HandBrake-1.7.3-source.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "HandBrake-1.7.3-source.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
