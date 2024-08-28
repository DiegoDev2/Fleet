
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mpdFormula represents a formula in Go.
type mpdFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mpdFormula) Print() {
    fmt.Printf("Name: mpd\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mpdFormula{
        Description:  "Music Player Daemon",
        Homepage:     "https://github.com/MusicPlayerDaemon/MPD",
        URL:          "https://github.com/MusicPlayerDaemon/MPD/commit/e380ae90ebb6325d1820b6f34e10bf3474710899.patch?full_index=1",
        Sha256:       "75d0d90d316c773424592fec34f88647029259f8f71b7b3547481bfa0250c2b3",
        Dependencies: []string{"boost", "meson", "ninja", "pkg-config", "chromaprint", "expat", "faad2", "ffmpeg", "flac", "fluid-synth", "fmt", "glib", "icu4c", "lame", "libao", "libgcrypt", "libid3tag", "libmpdclient", "libnfs", "libogg", "libsamplerate", "libshout", "libsndfile", "libsoxr", "libupnp", "libvorbis", "mpg123", "opus", "pcre2", "sqlite", "wavpack", "systemd", "alsa-lib", "dbus", "jack", "pulseaudio", "systemd"},
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

    fmt.Printf("Installing formula: %s\n", "mpd")
    if err := pkg.Installmpd(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mpdFormula) Installmpd() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "e380ae90ebb6325d1820b6f34e10bf3474710899.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "e380ae90ebb6325d1820b6f34e10bf3474710899"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
