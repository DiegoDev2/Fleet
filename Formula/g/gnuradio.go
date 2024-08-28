
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gnuradioFormula represents a formula in Go.
type gnuradioFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gnuradioFormula) Print() {
    fmt.Printf("Name: gnuradio\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gnuradioFormula{
        Description:  "SDK for signal processing blocks to implement software radios",
        Homepage:     "https://gnuradio.org/",
        URL:          "https://github.com/gnuradio/gnuradio/commit/ca344658756dab10a762571c51acf92c00c066c1.patch?full_index=1",
        Sha256:       "7e16ca70d07ce61bc16870f756acc194eb893e22703c53ed2826f5cf90dc7f4e",
        Dependencies: []string{"cmake", "doxygen", "pkg-config", "pybind11", "rust", "adwaita-icon-theme", "boost", "cppzmq", "fftw", "fmt", "gmp", "gsl", "gtk+3", "jack", "libsndfile", "libyaml", "log4cpp", "numpy", "portaudio", "pygobject3", "pyqt@5", "python@3.12", "qt@5", "qwt-qt5", "soapyrtlsdr", "soapysdr", "spdlog", "uhd", "volk", "zeromq", "alsa-lib", "llvm"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgnuradio(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gnuradioFormula) Installgnuradio() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "ca344658756dab10a762571c51acf92c00c066c1.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "ca344658756dab10a762571c51acf92c00c066c1"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
