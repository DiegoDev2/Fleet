
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// xorg-serverFormula represents a formula in Go.
type xorg-serverFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg xorg-serverFormula) Print() {
    fmt.Printf("Name: xorg-server\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := xorg-serverFormula{
        Description:  "X Window System display server",
        Homepage:     "https://www.x.org",
        URL:          "https://salsa.debian.org/xorg-team/xserver/xorg-server/-/raw/xorg-server-2_21.1.4-1/debian/local/xvfb-run.1",
        Sha256:       "08f14f55e14e52e5d98713c4d8f25ae68d67e2ee188dc0247770c6ada6e27c05",
        Dependencies: []string{"font-util", "libxkbfile", "meson", "ninja", "pkg-config", "util-macros", "xorgproto", "xtrans", "libx11", "libxau", "libxcb", "libxdmcp", "libxext", "libxfixes", "libxfont2", "mesa", "pixman", "xauth", "xcb-util", "xcb-util-image", "xcb-util-keysyms", "xcb-util-renderutil", "xcb-util-wm", "xkbcomp", "xkeyboardconfig", "libapplewm", "dbus", "libdrm", "libepoxy", "libpciaccess", "libtirpc", "libxcvt", "libxshmfence", "nettle", "systemd"},
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

    fmt.Printf("Installing formula: %s\n", "xorg-server")
    if err := pkg.Installxorg-server(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg xorg-serverFormula) Installxorg-server() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "xvfb-run.1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "xvfb-run"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
