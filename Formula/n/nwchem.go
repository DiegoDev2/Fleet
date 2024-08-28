
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// nwchemFormula represents a formula in Go.
type nwchemFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg nwchemFormula) Print() {
    fmt.Printf("Name: nwchem\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := nwchemFormula{
        Description:  "High-performance computational chemistry tools",
        Homepage:     "https://nwchemgit.github.io",
        URL:          "https://github.com/nwchemgit/nwchem/commit/bc18d20d90ba1fd6efc894558bef2fdacaac28a8.patch?full_index=1",
        Sha256:       "5432e8b0af47e80efb22f11774738e578919f5f857a7a3e46138a173910269d7",
        Dependencies: []string{"gcc", "hwloc", "libxc", "open-mpi", "openblas", "pkg-config", "python@3.12", "scalapack"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installnwchem(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg nwchemFormula) Installnwchem() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "bc18d20d90ba1fd6efc894558bef2fdacaac28a8.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "bc18d20d90ba1fd6efc894558bef2fdacaac28a8"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
