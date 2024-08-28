
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// jsonrpc-glibFormula represents a formula in Go.
type jsonrpc-glibFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg jsonrpc-glibFormula) Print() {
    fmt.Printf("Name: jsonrpc-glib\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := jsonrpc-glibFormula{
        Description:  "GNOME library to communicate with JSON-RPC based peers",
        Homepage:     "https://gitlab.gnome.org/GNOME/jsonrpc-glib",
        URL:          "https://download.gnome.org/sources/jsonrpc-glib/3.44/jsonrpc-glib-3.44.1.tar.xz",
        Sha256:       "a4403578f986fa53a5f8c41444aeb7ba65f34e2e5685ab58c0424457c2f2eddd",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "vala", "glib", "json-glib", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installjsonrpc-glib(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg jsonrpc-glibFormula) Installjsonrpc-glib() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "jsonrpc-glib-3.44.1.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "jsonrpc-glib-3.44.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
