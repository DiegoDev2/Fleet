
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// freeswitchFormula represents a formula in Go.
type freeswitchFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg freeswitchFormula) Print() {
    fmt.Printf("Name: freeswitch\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := freeswitchFormula{
        Description:  "Telephony platform to route various communication protocols",
        Homepage:     "https://freeswitch.org",
        URL:          "https://github.com/signalwire/signalwire-c.git",
        Sha256:       "980591a853fbf763818eb77132ea7e3ed876f8c4701e85070d612e1ebba09ae9",
        Dependencies: []string{"autoconf", "automake", "cmake", "libtool", "pkg-config", "yasm", "ffmpeg@5", "freetype", "jpeg-turbo", "ldns", "libpng", "libpq", "libsndfile", "libtiff", "lua", "opencore-amr", "openssl@3", "opus", "pcre", "sofia-sip", "speex", "speexdsp", "sqlite", "util-linux"},
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

    fmt.Printf("Installing formula: %s\n", "freeswitch")
    if err := pkg.Installfreeswitch(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg freeswitchFormula) Installfreeswitch() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "signalwire-c.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "signalwire-c"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
