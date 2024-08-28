
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// qtFormula represents a formula in Go.
type qtFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg qtFormula) Print() {
    fmt.Printf("Name: qt\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := qtFormula{
        Description:  "Cross-platform application and UI framework",
        Homepage:     "https://www.qt.io/",
        URL:          "https://github.com/qt/qtwebengine-chromium/commit/a766045f65f934df3b5f1aa63bc86fbb3e003a09.patch?full_index=1",
        Sha256:       "4d097145bf99a61704a99bf3e4b449b9bf613ae1f06efdcf44b881a045c5230c",
        Dependencies: []string{"cmake", "ninja", "node", "pkg-config", "python@3.12", "vulkan-headers", "vulkan-loader", "assimp", "brotli", "dbus", "double-conversion", "freetype", "glib", "harfbuzz", "hunspell", "icu4c", "jasper", "jpeg-turbo", "libb2", "libmng", "libpng", "libtiff", "md4c", "openssl@3", "pcre2", "sqlite", "webp", "zstd", "molten-vk", "alsa-lib", "at-spi2-core", "bluez", "ffmpeg", "fontconfig", "gstreamer", "gypsy", "libdrm", "libevent", "libice", "libsm", "libvpx", "libxcomposite", "libxkbcommon", "libxkbfile", "libxrandr", "libxtst", "little-cms2", "mesa", "minizip", "nss", "opus", "pulseaudio", "sdl2", "snappy", "systemd", "wayland", "xcb-util", "xcb-util-cursor", "xcb-util-image", "xcb-util-keysyms", "xcb-util-renderutil", "xcb-util-wm"},
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

    if err := pkg.Installqt(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg qtFormula) Installqt() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "a766045f65f934df3b5f1aa63bc86fbb3e003a09.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "a766045f65f934df3b5f1aa63bc86fbb3e003a09"
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
