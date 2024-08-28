
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// qt@5Formula represents a formula in Go.
type qt@5Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg qt@5Formula) Print() {
    fmt.Printf("Name: qt@5\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := qt@5Formula{
        Description:  "Cross-platform application and UI framework",
        Homepage:     "https://www.qt.io/",
        URL:          "https://download.qt.io/official_releases/qt/5.15/0002-CVE-2023-51714-qtbase-5.15.diff",
        Sha256:       "99d5d32527e767d6ab081ee090d92e0b11f27702619a4af8966b711db4f23e42",
        Dependencies: []string{"node", "pkg-config", "python@3.11", "freetype", "glib", "jpeg-turbo", "libpng", "pcre2", "webp", "alsa-lib", "at-spi2-core", "fontconfig", "harfbuzz", "icu4c", "libdrm", "libevent", "libice", "libproxy", "libsm", "libvpx", "libxcomposite", "libxdamage", "libxkbcommon", "libxkbfile", "libxrandr", "libxtst", "mesa", "minizip", "nss", "opus", "pulseaudio", "sdl2", "snappy", "systemd", "wayland", "xcb-util", "xcb-util-image", "xcb-util-keysyms", "xcb-util-renderutil", "xcb-util-wm", "zstd"},
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

    fmt.Printf("Installing formula: %s\n", "qt@5")
    if err := pkg.Installqt@5(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg qt@5Formula) Installqt@5() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "0002-CVE-2023-51714-qtbase-5.15.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "0002-CVE-2023-51714-qtbase-5.15"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
