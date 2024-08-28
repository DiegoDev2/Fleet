
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// x86_64-linux-gnu-binutilsFormula represents a formula in Go.
type x86_64-linux-gnu-binutilsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg x86_64-linux-gnu-binutilsFormula) Print() {
    fmt.Printf("Name: x86_64-linux-gnu-binutils\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := x86_64-linux-gnu-binutilsFormula{
        Description:  "GNU Binutils for x86_64-linux-gnu cross development",
        Homepage:     "https://www.gnu.org/software/binutils/binutils.html",
        URL:          "https://commondatastorage.googleapis.com/chrome-linux-sysroot/toolchain/2028cdaf24259d23adcff95393b8cc4f0eef714b/debian_bullseye_amd64_sysroot.tar.xz",
        Sha256:       "1be60e7c456abc590a613c64fab4eac7632c81ec6f22734a61b53669a4407346",
        Dependencies: []string{"pkg-config", "zstd", "texinfo"},
    }

    pkg.Print()

    // Instalar dependencias si no están instaladas
    for _, dep := range pkg.Dependencies {
        if !isDependencyInstalled(dep) {
            fmt.Printf("🛠️ Dependency %s not found. Installing...
", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("✅ Dependency %s is already installed.
", dep)
        }
    }

    if err := pkg.Installx86_64-linux-gnu-binutils(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg x86_64-linux-gnu-binutilsFormula) Installx86_64-linux-gnu-binutils() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "debian_bullseye_amd64_sysroot.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "debian_bullseye_amd64_sysroot.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}

func isDependencyInstalled(dep string) bool {
    cmd := exec.Command("brew", "list", dep)
    output, err := cmd.CombinedOutput()
    return err == nil && strings.TrimSpace(string(output)) != ""
}
