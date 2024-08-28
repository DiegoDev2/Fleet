
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pkcs11-toolsFormula represents a formula in Go.
type pkcs11-toolsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pkcs11-toolsFormula) Print() {
    fmt.Printf("Name: pkcs11-tools\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pkcs11-toolsFormula{
        Description:  "Tools to manage objects on PKCS#11 crypotographic tokens",
        Homepage:     "https://github.com/Mastercard/pkcs11-tools",
        URL:          "https://git.savannah.gnu.org/cgit/gnulib.git/patch/lib?id=cc91160a1ea5e18fcb2ccadb32e857d365581f53",
        Sha256:       "204de485eee7fdc9c63d60924bf2a2559bcddb2b13badbd60f97c8fcbd6ab4c3",
        Dependencies: []string{"pkg-config", "softhsm", "openssl@3"},
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

    fmt.Printf("Installing formula: %s\n", "pkcs11-tools")
    if err := pkg.Installpkcs11-tools(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pkcs11-toolsFormula) Installpkcs11-tools() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "lib?id=cc91160a1ea5e18fcb2ccadb32e857d365581f53"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "lib?id=cc91160a1ea5e18fcb2ccadb32e857d365581f53"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
