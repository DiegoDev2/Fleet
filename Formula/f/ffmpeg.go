package main

import (
	"fmt"
	"log"
	"os/exec"
)

// ffmpegFormula represents a formula in Go.
type ffmpegFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg ffmpegFormula) Print() {
	fmt.Printf("Name: ffmpeg\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := ffmpegFormula{
		Description:  "Play, record, convert, and stream audio and video",
		Homepage:     "https://ffmpeg.org/",
		URL:          "https://gitlab.archlinux.org/archlinux/packaging/packages/ffmpeg/-/raw/5670ccd86d3b816f49ebc18cab878125eca2f81f/add-av_stream_get_first_dts-for-chromium.patch",
		Sha256:       "57e26caced5a1382cb639235f9555fc50e45e7bf8333f7c9ae3d49b3241d3f77",
		Dependencies: []string{"pkg-config", "aom", "aribb24", "dav1d", "fontconfig", "freetype", "frei0r", "gnutls", "harfbuzz", "jpeg-xl", "lame", "libass", "libbluray", "librist", "libsoxr", "libssh", "libvidstab", "libvmaf", "libvorbis", "libvpx", "libx11", "libxcb", "opencore-amr", "openjpeg", "opus", "rav1e", "rubberband", "sdl2", "snappy", "speex", "srt", "svt-av1", "tesseract", "theora", "webp", "x264", "x265", "xvid", "xz", "zeromq", "zimg", "libarchive", "libogg", "libsamplerate", "alsa-lib", "libxext", "libxv", "nasm"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installffmpeg(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg ffmpegFormula) Installffmpeg() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "add-av_stream_get_first_dts-for-chromium.patch"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "add-av_stream_get_first_dts-for-chromium"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
