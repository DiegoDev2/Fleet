
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libxaw3dFormula represents a formula in Go.
type libxaw3dFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libxaw3dFormula) Print() {
    fmt.Printf("Name: libxaw3d\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libxaw3dFormula{
        Description:  "X.Org: 3D Athena widget set based on the Xt library",
        Homepage:     "https://www.x.org",
        URL:          "https://gitlab.freedesktop.org/xorg/lib/libxaw3d/-/commit/b2365950e5314b0453dd7cf2a552aa30ec19c046.diff",
        Sha256:       "18919b30bfafd2642895a4f9497f54b8d263e4eb593f192a923340732ed4afa8",
        Dependencies: []string{"pkg-config", "util-macros", "libx11", "libxext", "libxmu", "libxpm", "libxt"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installlibxaw3d(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libxaw3dFormula) Installlibxaw3d() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "b2365950e5314b0453dd7cf2a552aa30ec19c046.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "b2365950e5314b0453dd7cf2a552aa30ec19c046"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
