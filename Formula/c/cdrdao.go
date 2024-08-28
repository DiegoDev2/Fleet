
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// cdrdaoFormula represents a formula in Go.
type cdrdaoFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg cdrdaoFormula) Print() {
    fmt.Printf("Name: cdrdao\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := cdrdaoFormula{
        Description:  "Record CDs in Disk-At-Once mode",
        Homepage:     "https://cdrdao.sourceforge.net/",
        URL:          "https://github.com/cdrdao/cdrdao/commit/105d72a61f510e3c47626476f9bbc9516f824ede.patch?full_index=1",
        Sha256:       "0e235c0c34abaad56edb03a2526b3792f6f7ea12a8144cee48998cf1326894eb",
        Dependencies: []string{"autoconf", "automake", "pkg-config", "lame", "libao", "libvorbis", "mad"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installcdrdao(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg cdrdaoFormula) Installcdrdao() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "105d72a61f510e3c47626476f9bbc9516f824ede.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "105d72a61f510e3c47626476f9bbc9516f824ede"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
