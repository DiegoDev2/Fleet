
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtksourceview5Formula represents a formula in Go.
type gtksourceview5Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtksourceview5Formula) Print() {
    fmt.Printf("Name: gtksourceview5\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtksourceview5Formula{
        Description:  "Text view with syntax, undo/redo, and text marks",
        Homepage:     "https://projects.gnome.org/gtksourceview/",
        URL:          "https://download.gnome.org/sources/gtksourceview/5.12/gtksourceview-5.12.1.tar.xz",
        Sha256:       "96f6ff7d81b628f00bace8723469e5bb58033e10f8a685dac576a045bc700a02",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "vala", "cairo", "fontconfig", "fribidi", "gdk-pixbuf", "glib", "gtk4", "pango", "pcre2", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "gtksourceview5")
    if err := pkg.Installgtksourceview5(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtksourceview5Formula) Installgtksourceview5() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtksourceview-5.12.1.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtksourceview-5.12.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
