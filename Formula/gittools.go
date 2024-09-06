package main

// GitTools - Assorted git-related scripts and tools
// Homepage: https://github.com/MestreLion/git-tools

import (
	"fmt"
	
	"os/exec"
)

func installGitTools() {
	// Método 1: Descargar y extraer .tar.gz
	gittools_tar_url := "https://github.com/MestreLion/git-tools/archive/refs/tags/v2022.12.tar.gz"
	gittools_cmd_tar := exec.Command("curl", "-L", gittools_tar_url, "-o", "package.tar.gz")
	err := gittools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gittools_zip_url := "https://github.com/MestreLion/git-tools/archive/refs/tags/v2022.12.zip"
	gittools_cmd_zip := exec.Command("curl", "-L", gittools_zip_url, "-o", "package.zip")
	err = gittools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gittools_bin_url := "https://github.com/MestreLion/git-tools/archive/refs/tags/v2022.12.bin"
	gittools_cmd_bin := exec.Command("curl", "-L", gittools_bin_url, "-o", "binary.bin")
	err = gittools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gittools_src_url := "https://github.com/MestreLion/git-tools/archive/refs/tags/v2022.12.src.tar.gz"
	gittools_cmd_src := exec.Command("curl", "-L", gittools_src_url, "-o", "source.tar.gz")
	err = gittools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gittools_cmd_direct := exec.Command("./binary")
	err = gittools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
