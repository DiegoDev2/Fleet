package main

// GitVendor - Command for managing git vendored dependencies
// Homepage: https://brettlangdon.github.io/git-vendor

import (
	"fmt"
	
	"os/exec"
)

func installGitVendor() {
	// Método 1: Descargar y extraer .tar.gz
	gitvendor_tar_url := "https://github.com/brettlangdon/git-vendor/archive/refs/tags/v1.3.0.tar.gz"
	gitvendor_cmd_tar := exec.Command("curl", "-L", gitvendor_tar_url, "-o", "package.tar.gz")
	err := gitvendor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitvendor_zip_url := "https://github.com/brettlangdon/git-vendor/archive/refs/tags/v1.3.0.zip"
	gitvendor_cmd_zip := exec.Command("curl", "-L", gitvendor_zip_url, "-o", "package.zip")
	err = gitvendor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitvendor_bin_url := "https://github.com/brettlangdon/git-vendor/archive/refs/tags/v1.3.0.bin"
	gitvendor_cmd_bin := exec.Command("curl", "-L", gitvendor_bin_url, "-o", "binary.bin")
	err = gitvendor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitvendor_src_url := "https://github.com/brettlangdon/git-vendor/archive/refs/tags/v1.3.0.src.tar.gz"
	gitvendor_cmd_src := exec.Command("curl", "-L", gitvendor_src_url, "-o", "source.tar.gz")
	err = gitvendor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitvendor_cmd_direct := exec.Command("./binary")
	err = gitvendor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
