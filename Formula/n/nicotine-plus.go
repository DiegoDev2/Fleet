
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// nicotine-plusFormula represents a formula in Go.
type nicotine-plusFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg nicotine-plusFormula) Print() {
    fmt.Printf("Name: nicotine-plus\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := nicotine-plusFormula{
        Description:  "Graphical client for the Soulseek peer-to-peer network",
        Homepage:     "https://nicotine-plus.org",
        URL:          "https://files.pythonhosted.org/packages/63/1c/73f765da20b5b7e3579f6099490a9c4ac93e7c6341f97cf51d53ea0df49f/nicotine_plus-3.3.4.tar.gz",
        Sha256:       "603be73b6c65bf171d66df59924bb95264b65836e32a69e0a8c523dd5d076207",
        Dependencies: []string{"adwaita-icon-theme", "gtk4", "libadwaita", "py3cairo", "pygobject3", "python@3.12", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installnicotine-plus(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg nicotine-plusFormula) Installnicotine-plus() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "nicotine_plus-3.3.4.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "nicotine_plus-3.3.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
