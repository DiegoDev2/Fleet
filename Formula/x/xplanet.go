
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// xplanetFormula represents a formula in Go.
type xplanetFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg xplanetFormula) Print() {
    fmt.Printf("Name: xplanet\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := xplanetFormula{
        Description:  "Create HQ wallpapers of planet Earth",
        Homepage:     "https://xplanet.sourceforge.net/",
        URL:          "https://raw.githubusercontent.com/archlinux/svntogit-community/040965e32860345ca2d744239b6e257da33460a2/trunk/xplanet-c%2B%2B11.patch",
        Sha256:       "e651c7081c43ea48090186580b5a2a5d5039ab3ffbf34f7dd970037a16081454",
        Dependencies: []string{"pkg-config", "freetype", "giflib", "jpeg-turbo", "libpng", "libtiff", "automake"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installxplanet(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg xplanetFormula) Installxplanet() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "xplanet-c%2B%2B11.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "xplanet-c%2B%2B11"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
