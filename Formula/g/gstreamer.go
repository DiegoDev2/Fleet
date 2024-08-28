package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gstreamerFormula represents a formula in Go.
type gstreamerFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gstreamerFormula) Print() {
	fmt.Printf("Name: gstreamer\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gstreamerFormula{
		Description:  "Development framework for multimedia applications",
		Homepage:     "https://gstreamer.freedesktop.org/",
		URL:          "https://gitlab.freedesktop.org/gstreamer/gst-plugins-rs.git",
		Sha256:       "6c8884c463bc4e9cdb5cc0e3cadbd2695ea24df7beeddd65a7d4284213c0230e",
		Dependencies: []string{"bison", "cargo-c", "gettext", "gobject-introspection", "meson", "nasm", "ninja", "pkg-config", "rust", "aom", "cairo", "dav1d", "faac", "faad2", "fdk-aac", "ffmpeg", "flac", "gdk-pixbuf", "glib", "graphene", "gtk+3", "gtk4", "imath", "jpeg-turbo", "json-glib", "lame", "libass", "libogg", "libpng", "libshout", "libsndfile", "libsodium", "libsoup", "libusrsctp", "libvorbis", "libvpx", "libx11", "libxcb", "libxext", "libxfixes", "libxi", "libxtst", "little-cms2", "mpg123", "nettle", "opencore-amr", "openexr", "openjpeg", "openssl@3", "opus", "orc", "pango", "pygobject3", "python@3.12", "rtmpdump", "speex", "srt", "srtp", "svt-av1", "taglib", "theora", "webp", "x264", "x265", "gettext", "harfbuzz", "musepack", "alsa-lib", "libdrm", "libva", "libxdamage", "libxv", "mesa", "pulseaudio", "wayland"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgstreamer(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gstreamerFormula) Installgstreamer() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "gst-plugins-rs.git"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "gst-plugins-rs"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
