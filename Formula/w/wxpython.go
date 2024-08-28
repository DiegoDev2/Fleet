
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// wxpythonFormula represents a formula in Go.
type wxpythonFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg wxpythonFormula) Print() {
    fmt.Printf("Name: wxpython\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := wxpythonFormula{
        Description:  "Python bindings for wxWidgets",
        Homepage:     "https://www.wxpython.org/",
        URL:          "https://github.com/wxWidgets/Phoenix/commit/0b21230ee21e5e5d0212871b96a6d2fefd281038.patch?full_index=1",
        Sha256:       "befd2a9594a2fa41f926edf412476479f2f311b4088c4738a867c5e7ca6c0f82",
        Dependencies: []string{"doxygen", "python-setuptools", "sip", "numpy", "pillow", "python@3.12", "six", "wxwidgets", "pkg-config", "gtk+3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installwxpython(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg wxpythonFormula) Installwxpython() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "0b21230ee21e5e5d0212871b96a6d2fefd281038.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "0b21230ee21e5e5d0212871b96a6d2fefd281038"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
