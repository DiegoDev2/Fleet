
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// octaveFormula represents a formula in Go.
type octaveFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg octaveFormula) Print() {
    fmt.Printf("Name: octave\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := octaveFormula{
        Description:  "High-level interpreted language for numerical computing",
        Homepage:     "https://octave.org/index.html",
        URL:          "https://hg.savannah.gnu.org/hgweb/octave",
        Sha256:       "6287c397454f4d4c2c6e1871d0cbcd92e363fbb26f605242f264572c1be4c7f6",
        Dependencies: []string{"autoconf", "automake", "bison", "icoutils", "librsvg", "gnu-sed", "openjdk", "pkg-config", "arpack", "epstool", "fftw", "fig2dev", "fltk", "fontconfig", "freetype", "gcc", "ghostscript", "gl2ps", "glpk", "graphicsmagick", "hdf5", "libsndfile", "libtool", "openblas", "pcre2", "portaudio", "pstoedit", "qhull", "qrupdate", "qscintilla2", "qt", "rapidjson", "readline", "suite-sparse", "sundials", "texinfo", "little-cms2", "autoconf", "automake", "mesa", "mesa-glu"},
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

    fmt.Printf("Installing formula: %s\n", "octave")
    if err := pkg.Installoctave(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg octaveFormula) Installoctave() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "octave"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "octave"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
