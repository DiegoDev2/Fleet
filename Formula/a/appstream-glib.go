
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// appstream-glibFormula represents a formula in Go.
type appstream-glibFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg appstream-glibFormula) Print() {
    fmt.Printf("Name: appstream-glib\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := appstream-glibFormula{
        Description:  "Helper library for reading and writing AppStream metadata",
        Homepage:     "https://github.com/hughsie/appstream-glib",
        URL:          "https://github.com/hughsie/appstream-glib/archive/refs/tags/appstream_glib_0_8_3.tar.gz",
        Sha256:       "e3817d099c3f3d7bcdb0aec2920ff5b92e76454398d4efda5acb72c6d2ac4a1f",
        Dependencies: []string{"docbook", "docbook-xsl", "gobject-introspection", "meson", "ninja", "pkg-config", "gdk-pixbuf", "glib", "json-glib", "libarchive", "gettext", "util-linux"},
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

    fmt.Printf("Installing formula: %s\n", "appstream-glib")
    if err := pkg.Installappstream-glib(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg appstream-glibFormula) Installappstream-glib() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "appstream_glib_0_8_3.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "appstream_glib_0_8_3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
