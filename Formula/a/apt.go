
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// aptFormula represents a formula in Go.
type aptFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg aptFormula) Print() {
    fmt.Printf("Name: apt\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := aptFormula{
        Description:  "Advanced Package Tool",
        Homepage:     "https://wiki.debian.org/Apt",
        URL:          "https://cpan.metacpan.org/authors/id/P/PE/PEVANS/Syntax-Keyword-Try-0.29.tar.gz",
        Sha256:       "cc320719d3608daa9514743a43dac2be99cb8ccd989b1fefa285290cb1d59d8f",
        Dependencies: []string{"cmake", "docbook", "docbook-xsl", "doxygen", "googletest", "libxslt", "po4a", "w3m", "berkeley-db@5", "bzip2", "dpkg", "gettext", "gnupg", "gnutls", "libgcrypt", "lz4", "perl", "systemd", "xxhash", "xz", "zlib", "zstd"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installapt(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg aptFormula) Installapt() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "Syntax-Keyword-Try-0.29.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "Syntax-Keyword-Try-0.29.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
