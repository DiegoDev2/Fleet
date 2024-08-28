
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// audaciousFormula represents a formula in Go.
type audaciousFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg audaciousFormula) Print() {
    fmt.Printf("Name: audacious\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := audaciousFormula{
        Description:  "Lightweight and versatile audio player",
        Homepage:     "https://audacious-media-player.org/",
        URL:          "https://github.com/audacious-media-player/audacious-plugins.git",
        Sha256:       "695f2ae08a9e8b2e3e34358b64eaa566d39121334abf57c913323ccbeb9e5483",
        Dependencies: []string{"gettext", "meson", "ninja", "pkg-config", "faad2", "ffmpeg", "flac", "fluid-synth", "gdk-pixbuf", "glib", "lame", "libbs2b", "libcue", "libmodplug", "libnotify", "libogg", "libopenmpt", "libsamplerate", "libsndfile", "libsoxr", "libvorbis", "mpg123", "neon", "opusfile", "qt", "sdl2", "wavpack", "gettext", "opus", "alsa-lib", "jack", "libx11", "libxml2", "pulseaudio"},
    }

    pkg.Print()

    // Instalar dependencias si no están instaladas
    for _, dep := range pkg.Dependencies {
        if !isDependencyInstalled(dep) {
            fmt.Printf("🛠️ Dependency %s not found. Installing...
", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("✅ Dependency %s is already installed.
", dep)
        }
    }

    if err := pkg.Installaudacious(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg audaciousFormula) Installaudacious() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "audacious-plugins.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "audacious-plugins"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}

func isDependencyInstalled(dep string) bool {
    cmd := exec.Command("brew", "list", dep)
    output, err := cmd.CombinedOutput()
    return err == nil && strings.TrimSpace(string(output)) != ""
}
