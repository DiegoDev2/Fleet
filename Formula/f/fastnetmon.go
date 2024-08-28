
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// fastnetmonFormula represents a formula in Go.
type fastnetmonFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg fastnetmonFormula) Print() {
    fmt.Printf("Name: fastnetmon\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := fastnetmonFormula{
        Description:  "DDoS detection tool with sFlow, Netflow, IPFIX and port mirror support",
        Homepage:     "https://github.com/pavel-odintsov/fastnetmon/",
        URL:          "https://github.com/pavel-odintsov/fastnetmon/commit/fad8757b8986226024d549a6dfb40abbab01643e.patch?full_index=1",
        Sha256:       "2da8dbdf9dc63df9f17067aef20d198123ce1338559d394f29461761e6b85f85",
        Dependencies: []string{"cmake", "abseil", "boost", "capnp", "grpc", "hiredis", "log4cpp", "mongo-c-driver", "openssl@3", "protobuf", "elfutils", "libbpf", "libpcap"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installfastnetmon(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg fastnetmonFormula) Installfastnetmon() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "fad8757b8986226024d549a6dfb40abbab01643e.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "fad8757b8986226024d549a6dfb40abbab01643e"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
