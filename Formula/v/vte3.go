
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vte3Formula represents a formula in Go.
type vte3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vte3Formula) Print() {
    fmt.Printf("Name: vte3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vte3Formula{
        Description:  "Terminal emulator widget used by GNOME terminal",
        Homepage:     "https://wiki.gnome.org/Apps/Terminal/VTE",
        URL:          "https://download.gnome.org/sources/vte/0.76/vte-0.76.3.tar.xz",
        Sha256:       "7e3d98d527fad56e94e8026b7bd655d460a0e0e720a5e7141948d5361148e1a6",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "vala", "at-spi2-core", "cairo", "fribidi", "gdk-pixbuf", "glib", "gnutls", "gtk+3", "gtk4", "icu4c", "lz4", "pango", "pcre2", "llvm", "gettext", "llvm", "linux-headers@5.15", "systemd"},
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

    fmt.Printf("Installing formula: %s\n", "vte3")
    if err := pkg.Installvte3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vte3Formula) Installvte3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "vte-0.76.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "vte-0.76.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
