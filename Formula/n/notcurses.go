
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// notcursesFormula represents a formula in Go.
type notcursesFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg notcursesFormula) Print() {
    fmt.Printf("Name: notcurses\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := notcursesFormula{
        Description:  "Blingful character graphics/TUI library",
        Homepage:     "https://nick-black.com/dankwiki/index.php/Notcurses",
        URL:          "https://github.com/dankamongmen/notcurses/commit/441d66a063c7fc86436ed7ff73984050434c9142.patch?full_index=1",
        Sha256:       "aee69211bf5280bb773360a0f206e79f825ae86dbb7e05117d69acfa12917c13",
        Dependencies: []string{"cmake", "doctest", "pandoc", "pkg-config", "ffmpeg", "libdeflate", "libunistring", "ncurses"},
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

    fmt.Printf("Installing formula: %s\n", "notcurses")
    if err := pkg.Installnotcurses(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg notcursesFormula) Installnotcurses() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "441d66a063c7fc86436ed7ff73984050434c9142.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "441d66a063c7fc86436ed7ff73984050434c9142"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
