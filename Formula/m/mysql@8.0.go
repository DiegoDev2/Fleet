
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mysql@8.0Formula represents a formula in Go.
type mysql@8.0Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mysql@8.0Formula) Print() {
    fmt.Printf("Name: mysql@8.0\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mysql@8.0Formula{
        Description:  "Open source relational database management system",
        Homepage:     "https://dev.mysql.com/doc/refman/8.0/en/",
        URL:          "https://dev.mysql.com/downloads/mysql/8.0.html?tpl=files&os=src&version=8.0",
        Sha256:       "53dd31fd7b09579c160001b9981b8f6f54267c1fcbd79eb35732e2fcebb82074",
        Dependencies: []string{"bison", "cmake", "pkg-config", "abseil", "icu4c", "libevent", "libfido2", "lz4", "openssl@3", "protobuf", "zlib", "zstd", "patchelf", "libtirpc"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installmysql@8.0(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg mysql@8.0Formula) Installmysql@8.0() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "8.0.html?tpl=files&os=src&version=8.0"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "8.0.html?tpl=files&os=src&version=8"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
