package main

// GitlabGem - Ruby client and CLI for GitLab API
// Homepage: https://narkoz.github.io/gitlab/

import (
	"fmt"
	
	"os/exec"
)

func installGitlabGem() {
	// Método 1: Descargar y extraer .tar.gz
	gitlabgem_tar_url := "https://github.com/NARKOZ/gitlab/archive/refs/tags/v5.0.0.tar.gz"
	gitlabgem_cmd_tar := exec.Command("curl", "-L", gitlabgem_tar_url, "-o", "package.tar.gz")
	err := gitlabgem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitlabgem_zip_url := "https://github.com/NARKOZ/gitlab/archive/refs/tags/v5.0.0.zip"
	gitlabgem_cmd_zip := exec.Command("curl", "-L", gitlabgem_zip_url, "-o", "package.zip")
	err = gitlabgem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitlabgem_bin_url := "https://github.com/NARKOZ/gitlab/archive/refs/tags/v5.0.0.bin"
	gitlabgem_cmd_bin := exec.Command("curl", "-L", gitlabgem_bin_url, "-o", "binary.bin")
	err = gitlabgem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitlabgem_src_url := "https://github.com/NARKOZ/gitlab/archive/refs/tags/v5.0.0.src.tar.gz"
	gitlabgem_cmd_src := exec.Command("curl", "-L", gitlabgem_src_url, "-o", "source.tar.gz")
	err = gitlabgem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitlabgem_cmd_direct := exec.Command("./binary")
	err = gitlabgem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
