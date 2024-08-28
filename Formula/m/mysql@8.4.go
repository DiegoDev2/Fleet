
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mysql@8.4Formula represents a formula in Go.
type mysql@8.4Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mysql@8.4Formula) Print() {
    fmt.Printf("Name: mysql@8.4\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mysql@8.4Formula{
        Description:  "Open source relational database management system",
        Homepage:     "https://dev.mysql.com/doc/refman/8.4/en/",
        URL:          "https://dev.mysql.com/downloads/mysql/8.4.html?tpl=files&os=src&version=8.4",
        Sha256:       "efc13eb9812d9c206b60e4732638395a2166e01c81f4a22e774839e63a3542ae",
        Dependencies: []string{"bison", "cmake", "pkg-config", "abseil", "icu4c", "libfido2", "lz4", "openssl@3", "protobuf", "zlib", "zstd", "llvm", "patchelf", "libtirpc"},
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

    fmt.Printf("Installing formula: %s\n", "mysql@8.4")
    if err := pkg.Installmysql@8.4(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mysql@8.4Formula) Installmysql@8.4() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "8.4.html?tpl=files&os=src&version=8.4"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "8.4.html?tpl=files&os=src&version=8"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
