
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// arm-linux-gnueabihf-binutilsFormula represents a formula in Go.
type arm-linux-gnueabihf-binutilsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg arm-linux-gnueabihf-binutilsFormula) Print() {
    fmt.Printf("Name: arm-linux-gnueabihf-binutils\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := arm-linux-gnueabihf-binutilsFormula{
        Description:  "FSF/GNU binutils for cross-compiling to arm-linux",
        Homepage:     "https://www.gnu.org/software/binutils/binutils.html",
        URL:          "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2",
        Sha256:       "6dca015b099d6221e09b9acfba7514099d538d5bf57fa5b40940a7bf1bb256be",
        Dependencies: []string{"pkg-config", "zstd", "texinfo"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installarm-linux-gnueabihf-binutils(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg arm-linux-gnueabihf-binutilsFormula) Installarm-linux-gnueabihf-binutils() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "binutils-2.43.1.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "binutils-2.43.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
