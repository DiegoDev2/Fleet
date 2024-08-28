
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// musikcubeFormula represents a formula in Go.
type musikcubeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg musikcubeFormula) Print() {
    fmt.Printf("Name: musikcube\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := musikcubeFormula{
        Description:  "Terminal-based audio engine, library, player and server",
        Homepage:     "https://musikcube.com",
        URL:          "https://github.com/clangen/musikcube/archive/refs/tags/3.0.4.tar.gz",
        Sha256:       "e2ad36a2565ba7e15e751e5acd0bf92375dde852c83c7869b163dd66713012f6",
        Dependencies: []string{"asio", "cmake", "pkg-config", "ffmpeg", "game-music-emu", "gnutls", "lame", "libev", "libmicrohttpd", "libogg", "libopenmpt", "libvorbis", "ncurses", "openssl@3", "taglib", "mpg123", "portaudio"},
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

    fmt.Printf("Installing formula: %s\n", "musikcube")
    if err := pkg.Installmusikcube(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg musikcubeFormula) Installmusikcube() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "3.0.4.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "3.0.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
