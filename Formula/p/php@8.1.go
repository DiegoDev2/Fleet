
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// php@8.1Formula represents a formula in Go.
type php@8.1Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg php@8.1Formula) Print() {
    fmt.Printf("Name: php@8.1\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := php@8.1Formula{
        Description:  "General-purpose scripting language",
        Homepage:     "https://www.php.net/",
        URL:          "https://www.php.net/downloads",
        Sha256:       "f9b8c400b17157ad739167f0a79fafd5fdae79fad6971a13c3fac2b25cf5e945",
        Dependencies: []string{"httpd", "pkg-config", "apr", "apr-util", "argon2", "aspell", "autoconf", "curl", "freetds", "gd", "gettext", "gmp", "icu4c", "krb5", "libpq", "libsodium", "libzip", "oniguruma", "openldap", "openssl@3", "pcre2", "sqlite", "tidy-html5", "unixodbc"},
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

    fmt.Printf("Installing formula: %s\n", "php@8.1")
    if err := pkg.Installphp@8.1(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg php@8.1Formula) Installphp@8.1() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "downloads"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "downloads"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
