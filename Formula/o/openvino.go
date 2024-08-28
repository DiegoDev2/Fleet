
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// openvinoFormula represents a formula in Go.
type openvinoFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg openvinoFormula) Print() {
    fmt.Printf("Name: openvino\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := openvinoFormula{
        Description:  "Open Visual Inference And Optimization toolkit for AI inference",
        Homepage:     "https://docs.openvino.ai",
        URL:          "https://github.com/openvinotoolkit/openvino/commit/e653ebc7c8c11508c7e5fd4f797174d21e4382bc.patch?full_index=1",
        Sha256:       "d4b6eb705decaf9d8f7319a8cce69b64f9c719536138b510aa4b499b983b016c",
        Dependencies: []string{"cmake", "flatbuffers", "pkg-config", "protobuf@21", "pybind11", "python-setuptools", "python@3.12", "numpy", "pugixml", "snappy", "tbb", "opencl-clhpp-headers", "opencl-headers", "rapidjson", "opencl-icd-loader", "scons", "xbyak"},
    }

    pkg.Print()

    // Instalar dependencias si no están instaladas
    for _, dep := range pkg.Dependencies {
        if !isDependencyInstalled(dep) {
            fmt.Printf("🛠️ Dependency %s not found. Installing...
", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("✅ Dependency %s is already installed.
", dep)
        }
    }

    if err := pkg.Installopenvino(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg openvinoFormula) Installopenvino() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "e653ebc7c8c11508c7e5fd4f797174d21e4382bc.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "e653ebc7c8c11508c7e5fd4f797174d21e4382bc"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}

func isDependencyInstalled(dep string) bool {
    cmd := exec.Command("brew", "list", dep)
    output, err := cmd.CombinedOutput()
    return err == nil && strings.TrimSpace(string(output)) != ""
}
