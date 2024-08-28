
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// xcb-util-cursorFormula represents a formula in Go.
type xcb-util-cursorFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg xcb-util-cursorFormula) Print() {
    fmt.Printf("Name: xcb-util-cursor\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := xcb-util-cursorFormula{
        Description:  "XCB cursor library (replacement for libXcursor)",
        Homepage:     "https://xcb.freedesktop.org",
        URL:          "https://gitlab.freedesktop.org/xorg/lib/libxcb-cursor.git",
        Sha256:       "33e1dc7a4ba2663c341960faf21cb4e2f664e877f096dcb29ff3036f15dba730",
        Dependencies: []string{"autoconf", "automake", "libtool", "util-macros", "pkg-config", "libxcb", "xcb-util", "xcb-util-image", "xcb-util-renderutil"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installxcb-util-cursor(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg xcb-util-cursorFormula) Installxcb-util-cursor() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libxcb-cursor.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libxcb-cursor"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
