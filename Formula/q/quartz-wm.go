
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// quartz-wmFormula represents a formula in Go.
type quartz-wmFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg quartz-wmFormula) Print() {
    fmt.Printf("Name: quartz-wm\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := quartz-wmFormula{
        Description:  "XQuartz window-manager",
        Homepage:     "https://gitlab.freedesktop.org/xorg/app/quartz-wm",
        URL:          "https://gitlab.freedesktop.org/xorg/app/quartz-wm/-/archive/babff9d70f61239c46c53a3e41ce10c7ca1419ce/quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar.bz2",
        Sha256:       "4d0abe6b48b2ce5ee2c6c44c3c7af67201f0df90b2fe7d75e9273238df19fdce",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "util-macros", "xorg-server", "libapplewm", "libx11", "libxext", "libxinerama", "libxrandr", "pixman"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installquartz-wm(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg quartz-wmFormula) Installquartz-wm() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "quartz-wm-babff9d70f61239c46c53a3e41ce10c7ca1419ce.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
