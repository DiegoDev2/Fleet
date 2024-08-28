
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// etcd-cpp-apiv3Formula represents a formula in Go.
type etcd-cpp-apiv3Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg etcd-cpp-apiv3Formula) Print() {
    fmt.Printf("Name: etcd-cpp-apiv3\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := etcd-cpp-apiv3Formula{
        Description:  "C++ implementation for etcd's v3 client API, i.e., ETCDCTL_API=3",
        Homepage:     "https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3",
        URL:          "https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3/archive/refs/tags/v0.15.4.tar.gz",
        Sha256:       "fcb362d41bcd81d856636f1f5673a24501f3401e434121e72e0d2769b826b889",
        Dependencies: []string{"cmake", "etcd", "abseil", "boost", "c-ares", "cpprestsdk", "grpc", "openssl@3", "protobuf", "re2"},
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

    fmt.Printf("Installing formula: %s\n", "etcd-cpp-apiv3")
    if err := pkg.Installetcd-cpp-apiv3(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg etcd-cpp-apiv3Formula) Installetcd-cpp-apiv3() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v0.15.4.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v0.15.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
