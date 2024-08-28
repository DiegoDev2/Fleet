
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// dwarfsFormula represents a formula in Go.
type dwarfsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg dwarfsFormula) Print() {
    fmt.Printf("Name: dwarfs\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := dwarfsFormula{
        Description:  "Fast high compression read-only file system for Linux, Windows, and macOS",
        Homepage:     "https://github.com/mhx/dwarfs",
        URL:          "https://github.com/mhx/dwarfs/releases/download/v0.10.1/dwarfs-0.10.1.tar.xz",
        Sha256:       "9cafeff56db381b4ec58754b7d33326f174757845338e138a1c14f4673be3247",
        Dependencies: []string{"cmake", "googletest", "pkg-config", "boost", "brotli", "double-conversion", "flac", "fmt", "gflags", "glog", "howard-hinnant-date", "libarchive", "libevent", "libsodium", "lz4", "nlohmann-json", "openssl@3", "parallel-hashmap", "range-v3", "utf8cpp", "xxhash", "xz", "zstd", "llvm", "jemalloc"},
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

    fmt.Printf("Installing formula: %s\n", "dwarfs")
    if err := pkg.Installdwarfs(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg dwarfsFormula) Installdwarfs() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "dwarfs-0.10.1.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "dwarfs-0.10.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
