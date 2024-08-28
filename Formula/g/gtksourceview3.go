
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtksourceview3Formula represents a formula in Go.
type gtksourceview3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtksourceview3Formula) Print() {
    fmt.Printf("Name: gtksourceview3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtksourceview3Formula{
        Description:  "Text view with syntax, undo/redo, and text marks",
        Homepage:     "https://projects.gnome.org/gtksourceview/",
        URL:          "https://download.gnome.org/sources/gtksourceview/3.24/gtksourceview-3.24.11.tar.xz",
        Sha256:       "a29fc106c186da13c48399358ec0152907ca4bc3be2545632400409aaffcce2e",
        Dependencies: []string{"gobject-introspection", "pkg-config", "vala", "at-spi2-core", "cairo", "gdk-pixbuf", "glib", "gtk+3", "pango", "autoconf", "automake", "gtk-doc", "libtool", "gettext", "harfbuzz"},
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

    fmt.Printf("Installing formula: %s\n", "gtksourceview3")
    if err := pkg.Installgtksourceview3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtksourceview3Formula) Installgtksourceview3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtksourceview-3.24.11.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtksourceview-3.24.11.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
