
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// phpFormula represents a formula in Go.
type phpFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg phpFormula) Print() {
    fmt.Printf("Name: php\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := phpFormula{
        Description:  "General-purpose scripting language",
        Homepage:     "https://www.php.net/",
        URL:          "https://github.com/php/php-src.git",
        Sha256:       "111e4a51d201e6d99f94256cf0bf4da6adeb88754616a677e77ada74170c19b7",
        Dependencies: []string{"bison", "re2c", "httpd", "pkg-config", "apr", "apr-util", "argon2", "aspell", "autoconf", "curl", "freetds", "gd", "gettext", "gmp", "icu4c", "krb5", "libpq", "libsodium", "libzip", "oniguruma", "openldap", "openssl@3", "pcre2", "sqlite", "tidy-html5", "unixodbc"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installphp(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg phpFormula) Installphp() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "php-src.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "php-src"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
