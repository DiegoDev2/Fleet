
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// csoundFormula represents a formula in Go.
type csoundFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg csoundFormula) Print() {
    fmt.Printf("Name: csound\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := csoundFormula{
        Description:  "Sound and music computing system",
        Homepage:     "https://csound.com",
        URL:          "https://download.savannah.gnu.org/releases/getfem/stable/getfem-5.4.2.tar.gz",
        Sha256:       "80b625d5892fe9959c3b316340f326e3ece4e98325eb0a81dd5b9ddae563b1d1",
        Dependencies: []string{"asio", "cmake", "eigen", "swig", "faust", "fltk", "fluid-synth", "gettext", "hdf5", "jack", "lame", "liblo", "libpng", "libsamplerate", "libsndfile", "libwebsockets", "numpy", "openjdk", "openssl@3", "portaudio", "portmidi", "python@3.12", "stk", "wiiuse", "alsa-lib", "libx11"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installcsound(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg csoundFormula) Installcsound() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "getfem-5.4.2.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "getfem-5.4.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
