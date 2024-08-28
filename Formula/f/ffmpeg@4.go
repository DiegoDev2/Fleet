
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ffmpeg@4Formula represents a formula in Go.
type ffmpeg@4Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ffmpeg@4Formula) Print() {
    fmt.Printf("Name: ffmpeg@4\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ffmpeg@4Formula{
        Description:  "Play, record, convert, and stream audio and video",
        Homepage:     "https://ffmpeg.org/",
        URL:          "https://ffmpeg.org/download.html",
        Sha256:       "b290118efecb94b72600d14677ca5ec69f660cddbbf53e92eb7162cd1b74fd81",
        Dependencies: []string{"nasm", "pkg-config", "aom", "dav1d", "fontconfig", "freetype", "frei0r", "gnutls", "lame", "libass", "libbluray", "librist", "libsoxr", "libvidstab", "libvorbis", "libvpx", "libxcb", "opencore-amr", "openjpeg", "opus", "rav1e", "rubberband", "sdl2", "snappy", "speex", "srt", "tesseract", "theora", "webp", "x264", "x265", "xvid", "xz", "zeromq", "zimg", "libarchive", "libogg", "libsamplerate", "libvmaf", "alsa-lib", "libx11", "libxext", "libxv"},
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

    fmt.Printf("Installing formula: %s\n", "ffmpeg@4")
    if err := pkg.Installffmpeg@4(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg ffmpeg@4Formula) Installffmpeg@4() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "download.html"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "download"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
