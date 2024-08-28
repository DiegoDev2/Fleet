
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gnumericFormula represents a formula in Go.
type gnumericFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gnumericFormula) Print() {
    fmt.Printf("Name: gnumeric\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gnumericFormula{
        Description:  "GNOME Spreadsheet Application",
        Homepage:     "https://projects.gnome.org/gnumeric/",
        URL:          "https://download.gnome.org/sources/gnumeric/1.12/gnumeric-1.12.57.tar.xz",
        Sha256:       "22083c21842d8e41c7a37d5af2bb683d6fb1fb52347db917785fd02024a6a341",
        Dependencies: []string{"gettext", "intltool", "itstool", "pkg-config", "adwaita-icon-theme", "at-spi2-core", "cairo", "gdk-pixbuf", "glib", "goffice", "gtk+3", "libgsf", "libxml2", "pango", "gettext", "harfbuzz", "perl-xml-parser"},
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

    fmt.Printf("Installing formula: %s\n", "gnumeric")
    if err := pkg.Installgnumeric(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gnumericFormula) Installgnumeric() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gnumeric-1.12.57.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gnumeric-1.12.57.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
