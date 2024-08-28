
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mariadb@10.8Formula represents a formula in Go.
type mariadb@10.8Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mariadb@10.8Formula) Print() {
    fmt.Printf("Name: mariadb@10.8\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mariadb@10.8Formula{
        Description:  "Drop-in replacement for MySQL",
        Homepage:     "https://mariadb.org/",
        URL:          "https://downloads.mariadb.com/MariaDB/mariadb-10.8.8/source/mariadb-10.8.8.tar.gz",
        Sha256:       "82bc7398f14cb5ff8943c114f4049e67ecb6f0aa9ff45973b115686f6b924643",
        Dependencies: []string{"bison", "cmake", "fmt", "pkg-config", "groonga", "openssl@3", "pcre2", "zstd", "linux-pam", "readline"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installmariadb@10.8(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg mariadb@10.8Formula) Installmariadb@10.8() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "mariadb-10.8.8.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "mariadb-10.8.8.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
