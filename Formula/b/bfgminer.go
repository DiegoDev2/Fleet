
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// bfgminerFormula represents a formula in Go.
type bfgminerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg bfgminerFormula) Print() {
    fmt.Printf("Name: bfgminer\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := bfgminerFormula{
        Description:  "Modular CPU/GPU/ASIC/FPGA miner written in C",
        Homepage:     "https://web.archive.org/web/20221230131107/http://bfgminer.org/",
        URL:          "https://web.archive.org/web/20190824104403/http://bfgminer.org/files/latest/bfgminer-5.5.0.txz",
        Sha256:       "9f81c9ed9c2d32a0296fb74aa46b592b02e60fc08224d103c9c6ba73bedff86b",
        Dependencies: []string{"hidapi", "libgcrypt", "libscrypt", "libtool", "pkg-config", "uthash", "jansson", "libevent", "libusb"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installbfgminer(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg bfgminerFormula) Installbfgminer() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "bfgminer-5.5.0.txz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "bfgminer-5.5.0"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
