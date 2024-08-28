
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// dartsimFormula represents a formula in Go.
type dartsimFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg dartsimFormula) Print() {
    fmt.Printf("Name: dartsim\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := dartsimFormula{
        Description:  "Dynamic Animation and Robotics Toolkit",
        Homepage:     "https://dartsim.github.io/",
        URL:          "https://github.com/dartsim/dart/archive/refs/tags/v6.14.4.tar.gz",
        Sha256:       "97c73b3716e432237f9fd0152a81b18ef58026b35a02690dcffbbda886751b4e",
        Dependencies: []string{"cmake", "pkg-config", "assimp", "bullet", "eigen", "fcl", "flann", "fmt", "ipopt", "libccd", "nlopt", "octomap", "ode", "open-scene-graph", "spdlog", "tinyxml2", "urdfdom", "mesa"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installdartsim(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg dartsimFormula) Installdartsim() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v6.14.4.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v6.14.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
