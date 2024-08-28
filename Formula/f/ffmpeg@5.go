
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ffmpeg@5Formula represents a formula in Go.
type ffmpeg@5Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ffmpeg@5Formula) Print() {
    fmt.Printf("Name: ffmpeg@5\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ffmpeg@5Formula{
        Description:  "Play, record, convert, and stream audio and video",
        Homepage:     "https://ffmpeg.org/",
        URL:          "https://ffmpeg.org/download.html",
        Sha256:       "53cd9c93078f6a47f56e932956b6987be3bcf759a905723832a312d486a5d686",
        Dependencies: []string{"pkg-config", "aom", "aribb24", "dav1d", "fontconfig", "freetype", "frei0r", "gnutls", "lame", "libass", "libbluray", "librist", "libsoxr", "libvidstab", "libvmaf", "libvorbis", "libvpx", "libx11", "libxcb", "opencore-amr", "openjpeg", "opus", "rav1e", "rubberband", "sdl2", "snappy", "speex", "srt", "svt-av1", "tesseract", "theora", "webp", "x264", "x265", "xvid", "xz", "zeromq", "zimg", "libarchive", "libogg", "libsamplerate", "alsa-lib", "libxext", "libxv", "nasm"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installffmpeg@5(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg ffmpeg@5Formula) Installffmpeg@5() error {
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
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
