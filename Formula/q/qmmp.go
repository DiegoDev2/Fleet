
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// qmmpFormula represents a formula in Go.
type qmmpFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg qmmpFormula) Print() {
    fmt.Printf("Name: qmmp\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := qmmpFormula{
        Description:  "Qt-based Multimedia Player",
        Homepage:     "https://qmmp.ylsoftware.com/",
        URL:          "https://qmmp.ylsoftware.com/files/qmmp-plugin-pack/2.1/qmmp-plugin-pack-2.1.2.tar.bz2",
        Sha256:       "fb5b7381a7f11a31e686fb7c76213d42dfa5df1ec61ac2c7afccd8697730c84b",
        Dependencies: []string{"cmake", "pkg-config", "faad2", "ffmpeg", "flac", "game-music-emu", "jack", "libarchive", "libbs2b", "libcddb", "libcdio", "libmms", "libmodplug", "libogg", "libsamplerate", "libshout", "libsndfile", "libsoxr", "libvorbis", "libxcb", "libxmp", "mad", "mpg123", "mplayer", "opus", "opusfile", "projectm", "pulseaudio", "qt", "taglib", "wavpack", "wildmidi", "gettext", "glib", "musepack", "alsa-lib", "libx11", "mesa"},
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

    fmt.Printf("Installing formula: %s\n", "qmmp")
    if err := pkg.Installqmmp(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg qmmpFormula) Installqmmp() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "qmmp-plugin-pack-2.1.2.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "qmmp-plugin-pack-2.1.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
