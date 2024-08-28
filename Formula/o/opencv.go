
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// opencvFormula represents a formula in Go.
type opencvFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg opencvFormula) Print() {
    fmt.Printf("Name: opencv\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := opencvFormula{
        Description:  "Open source computer vision library",
        Homepage:     "https://opencv.org/",
        URL:          "https://github.com/opencv/opencv_contrib/archive/refs/tags/4.10.0.tar.gz",
        Sha256:       "65597f8fb8dc2b876c1b45b928bbcc5f772ddbaf97539bf1b737623d0604cba1",
        Dependencies: []string{"cmake", "pkg-config", "python-setuptools", "abseil", "ceres-solver", "eigen", "ffmpeg", "freetype", "gflags", "glog", "harfbuzz", "jpeg-turbo", "libpng", "libtiff", "numpy", "openblas", "openexr", "openjpeg", "openvino", "protobuf", "python@3.12", "tbb", "tesseract", "vtk", "webp", "glew", "imath", "jsoncpp", "libarchive", "cairo", "gdk-pixbuf", "glib", "gtk+3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installopencv(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg opencvFormula) Installopencv() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "4.10.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "4.10.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
