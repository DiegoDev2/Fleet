
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// percona-xtrabackupFormula represents a formula in Go.
type percona-xtrabackupFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg percona-xtrabackupFormula) Print() {
    fmt.Printf("Name: percona-xtrabackup\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := percona-xtrabackupFormula{
        Description:  "Open source hot backup tool for InnoDB and XtraDB databases",
        Homepage:     "https://www.percona.com/software/mysql-database/percona-xtrabackup",
        URL:          "https://boostorg.jfrog.io/artifactory/main/release/1.77.0/source/boost_1_77_0.tar.bz2",
        Sha256:       "fc9f85fc030e233142908241af7a846e60630aa7388de9a5fafb1f3a26840854",
        Dependencies: []string{"bison", "cmake", "pkg-config", "sphinx-doc", "icu4c", "libev", "libevent", "libfido2", "libgcrypt", "lz4", "mysql-client", "openssl@3", "protobuf@21", "zlib", "zstd", "libgpg-error", "patchelf", "libaio", "procps"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installpercona-xtrabackup(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg percona-xtrabackupFormula) Installpercona-xtrabackup() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "boost_1_77_0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "boost_1_77_0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
