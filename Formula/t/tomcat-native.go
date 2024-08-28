
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// tomcat-nativeFormula represents a formula in Go.
type tomcat-nativeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg tomcat-nativeFormula) Print() {
    fmt.Printf("Name: tomcat-native\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := tomcat-nativeFormula{
        Description:  "Lets Tomcat use some native resources for performance",
        Homepage:     "https://tomcat.apache.org/native-doc/",
        URL:          "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-connectors/native/2.0.8/source/tomcat-native-2.0.8-src.tar.gz",
        Sha256:       "3b6871bc6de4e15f504981b8930ccd64e3b955ae2b73ec5642376f9dbf3d5efc",
        Dependencies: []string{"tomcat", "apr", "openjdk", "openssl@3"},
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

    fmt.Printf("Installing formula: %s\n", "tomcat-native")
    if err := pkg.Installtomcat-native(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg tomcat-nativeFormula) Installtomcat-native() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "tomcat-native-2.0.8-src.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "tomcat-native-2.0.8-src.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
