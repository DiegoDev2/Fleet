package main

// GitAppraise - Distributed code review system for Git repos
// Homepage: https://github.com/google/git-appraise

import (
	"fmt"
	
	"os/exec"
)

func installGitAppraise() {
	// Método 1: Descargar y extraer .tar.gz
	gitappraise_tar_url := "https://github.com/google/git-appraise/archive/refs/tags/v0.7.tar.gz"
	gitappraise_cmd_tar := exec.Command("curl", "-L", gitappraise_tar_url, "-o", "package.tar.gz")
	err := gitappraise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitappraise_zip_url := "https://github.com/google/git-appraise/archive/refs/tags/v0.7.zip"
	gitappraise_cmd_zip := exec.Command("curl", "-L", gitappraise_zip_url, "-o", "package.zip")
	err = gitappraise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitappraise_bin_url := "https://github.com/google/git-appraise/archive/refs/tags/v0.7.bin"
	gitappraise_cmd_bin := exec.Command("curl", "-L", gitappraise_bin_url, "-o", "binary.bin")
	err = gitappraise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitappraise_src_url := "https://github.com/google/git-appraise/archive/refs/tags/v0.7.src.tar.gz"
	gitappraise_cmd_src := exec.Command("curl", "-L", gitappraise_src_url, "-o", "source.tar.gz")
	err = gitappraise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitappraise_cmd_direct := exec.Command("./binary")
	err = gitappraise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
