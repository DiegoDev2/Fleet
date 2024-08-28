
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// jpeg-xlFormula represents a formula in Go.
type jpeg-xlFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg jpeg-xlFormula) Print() {
    fmt.Printf("Name: jpeg-xl\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := jpeg-xlFormula{
        Description:  "New file format for still image compression",
        Homepage:     "https://jpeg.org/jpegxl/index.html",
        URL:          "https://github.com/webmproject/sjpeg.git",
        Sha256:       "64b61d8a8d21bed6842b83b3e2810fe12ad1d77bdad9781850edef818a1355e7",
        Dependencies: []string{"asciidoc", "cmake", "docbook-xsl", "doxygen", "pkg-config", "sphinx-doc", "pkg-config", "brotli", "giflib", "highway", "imath", "jpeg-turbo", "libpng", "little-cms2", "openexr", "webp"},
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

    fmt.Printf("Installing formula: %s\n", "jpeg-xl")
    if err := pkg.Installjpeg-xl(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg jpeg-xlFormula) Installjpeg-xl() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "sjpeg.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "sjpeg"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
