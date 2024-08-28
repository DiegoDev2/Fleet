
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtkmm3Formula represents a formula in Go.
type gtkmm3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtkmm3Formula) Print() {
    fmt.Printf("Name: gtkmm3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtkmm3Formula{
        Description:  "C++ interfaces for GTK+ and GNOME",
        Homepage:     "https://www.gtkmm.org/",
        URL:          "https://download.gnome.org/sources/gtkmm/3.24/gtkmm-3.24.9.tar.xz",
        Sha256:       "548ac36e8284ef35d9bb0700f63a50deb54ebfb67b27de3e72b0ab45793b6232",
        Dependencies: []string{"meson", "ninja", "pkg-config", "atkmm@2.28", "cairomm@1.14", "gdk-pixbuf", "glib", "glibmm@2.66", "gtk+3", "libsigc++@2", "pangomm@2.46"},
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

    fmt.Printf("Installing formula: %s\n", "gtkmm3")
    if err := pkg.Installgtkmm3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtkmm3Formula) Installgtkmm3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtkmm-3.24.9.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtkmm-3.24.9.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
