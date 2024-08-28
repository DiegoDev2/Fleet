
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// llvm@17Formula represents a formula in Go.
type llvm@17Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg llvm@17Formula) Print() {
    fmt.Printf("Name: llvm@17\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := llvm@17Formula{
        Description:  "Next-gen compiler infrastructure",
        Homepage:     "https://llvm.org/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/23704400c86976aaa4f421f56928484a270ac79c/llvm/17.x-arm64-opt.patch",
        Sha256:       "0e312207fd9474bd26f4a283ee23d94b334d3ec8732086d30bce95f7c8dc2201",
        Dependencies: []string{"cmake", "ninja", "python@3.12", "zstd", "pkg-config", "binutils", "elfutils"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installllvm@17(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg llvm@17Formula) Installllvm@17() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "17.x-arm64-opt.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "17.x-arm64-opt"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
