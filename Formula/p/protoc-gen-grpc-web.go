
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// protoc-gen-grpc-webFormula represents a formula in Go.
type protoc-gen-grpc-webFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg protoc-gen-grpc-webFormula) Print() {
    fmt.Printf("Name: protoc-gen-grpc-web\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := protoc-gen-grpc-webFormula{
        Description:  "Protoc plugin that generates code for gRPC-Web clients",
        Homepage:     "https://github.com/grpc/grpc-web",
        URL:          "https://github.com/grpc/grpc-web/archive/refs/tags/1.5.0.tar.gz",
        Sha256:       "066c224c130c3a628c49551eac976a0bd4eea89314bc9195b31b13d3892f1160",
        Dependencies: []string{"cmake", "node", "typescript", "protobuf@21", "protoc-gen-js"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installprotoc-gen-grpc-web(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg protoc-gen-grpc-webFormula) Installprotoc-gen-grpc-web() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "1.5.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "1.5.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
