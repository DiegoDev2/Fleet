package main

// GitExtras - Small git utilities
// Homepage: https://github.com/tj/git-extras

import (
	"fmt"
	
	"os/exec"
)

func installGitExtras() {
	// Método 1: Descargar y extraer .tar.gz
	gitextras_tar_url := "https://github.com/tj/git-extras/archive/refs/tags/7.2.0.tar.gz"
	gitextras_cmd_tar := exec.Command("curl", "-L", gitextras_tar_url, "-o", "package.tar.gz")
	err := gitextras_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitextras_zip_url := "https://github.com/tj/git-extras/archive/refs/tags/7.2.0.zip"
	gitextras_cmd_zip := exec.Command("curl", "-L", gitextras_zip_url, "-o", "package.zip")
	err = gitextras_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitextras_bin_url := "https://github.com/tj/git-extras/archive/refs/tags/7.2.0.bin"
	gitextras_cmd_bin := exec.Command("curl", "-L", gitextras_bin_url, "-o", "binary.bin")
	err = gitextras_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitextras_src_url := "https://github.com/tj/git-extras/archive/refs/tags/7.2.0.src.tar.gz"
	gitextras_cmd_src := exec.Command("curl", "-L", gitextras_src_url, "-o", "source.tar.gz")
	err = gitextras_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitextras_cmd_direct := exec.Command("./binary")
	err = gitextras_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
