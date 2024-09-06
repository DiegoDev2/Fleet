package main

// Gitbackup - Tool to backup your Bitbucket, GitHub and GitLab repositories
// Homepage: https://github.com/amitsaha/gitbackup

import (
	"fmt"
	
	"os/exec"
)

func installGitbackup() {
	// Método 1: Descargar y extraer .tar.gz
	gitbackup_tar_url := "https://github.com/amitsaha/gitbackup/archive/refs/tags/v0.9.0.tar.gz"
	gitbackup_cmd_tar := exec.Command("curl", "-L", gitbackup_tar_url, "-o", "package.tar.gz")
	err := gitbackup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitbackup_zip_url := "https://github.com/amitsaha/gitbackup/archive/refs/tags/v0.9.0.zip"
	gitbackup_cmd_zip := exec.Command("curl", "-L", gitbackup_zip_url, "-o", "package.zip")
	err = gitbackup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitbackup_bin_url := "https://github.com/amitsaha/gitbackup/archive/refs/tags/v0.9.0.bin"
	gitbackup_cmd_bin := exec.Command("curl", "-L", gitbackup_bin_url, "-o", "binary.bin")
	err = gitbackup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitbackup_src_url := "https://github.com/amitsaha/gitbackup/archive/refs/tags/v0.9.0.src.tar.gz"
	gitbackup_cmd_src := exec.Command("curl", "-L", gitbackup_src_url, "-o", "source.tar.gz")
	err = gitbackup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitbackup_cmd_direct := exec.Command("./binary")
	err = gitbackup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
