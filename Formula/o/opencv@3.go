
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// opencv@3Formula represents a formula in Go.
type opencv@3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg opencv@3Formula) Print() {
    fmt.Printf("Name: opencv@3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := opencv@3Formula{
        Description:  "Open source computer vision library",
        Homepage:     "https://opencv.org/",
        URL:          "https://github.com/opencv/opencv_contrib/archive/refs/tags/3.4.20.tar.gz",
        Sha256:       "b0bb3fa7ae4ac00926b83d4d95c6500c2f7af542f8ec78d0f01b2961a690d5dc",
        Dependencies: []string{"cmake", "pkg-config", "python-setuptools", "ceres-solver", "eigen", "ffmpeg@4", "gflags", "glog", "jpeg-turbo", "libpng", "libtiff", "numpy", "openexr", "protobuf@21", "python@3.12", "tbb", "webp"},
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

    fmt.Printf("Installing formula: %s\n", "opencv@3")
    if err := pkg.Installopencv@3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg opencv@3Formula) Installopencv@3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "3.4.20.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "3.4.20.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
