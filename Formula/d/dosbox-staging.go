
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// dosbox-stagingFormula represents a formula in Go.
type dosbox-stagingFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg dosbox-stagingFormula) Print() {
    fmt.Printf("Name: dosbox-staging\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := dosbox-stagingFormula{
        Description:  "Modernized DOSBox soft-fork",
        Homepage:     "https://dosbox-staging.github.io/",
        URL:          "https://github.com/dosbox-staging/dosbox-staging/commit/9f0fc1dc762010e5f7471d01c504d817a066cae3.patch?full_index=1",
        Sha256:       "dda24c081c9437844203c3624888ab07298f74c0b29f5dca250c9e071b764a7f",
        Dependencies: []string{"meson", "ninja", "pkg-config", "fluid-synth", "glib", "iir1", "libpng", "libslirp", "mt32emu", "opusfile", "sdl2", "sdl2_image", "sdl2_net", "speexdsp", "alsa-lib", "mesa", "mesa-glu"},
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

    fmt.Printf("Installing formula: %s\n", "dosbox-staging")
    if err := pkg.Installdosbox-staging(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg dosbox-stagingFormula) Installdosbox-staging() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "9f0fc1dc762010e5f7471d01c504d817a066cae3.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "9f0fc1dc762010e5f7471d01c504d817a066cae3"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
