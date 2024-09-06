package main

// Gitea - Painless self-hosted all-in-one software development service
// Homepage: https://about.gitea.com/

import (
	"fmt"
	
	"os/exec"
)

func installGitea() {
	// Método 1: Descargar y extraer .tar.gz
	gitea_tar_url := "https://dl.gitea.com/gitea/1.22.2/gitea-src-1.22.2.tar.gz"
	gitea_cmd_tar := exec.Command("curl", "-L", gitea_tar_url, "-o", "package.tar.gz")
	err := gitea_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitea_zip_url := "https://dl.gitea.com/gitea/1.22.2/gitea-src-1.22.2.zip"
	gitea_cmd_zip := exec.Command("curl", "-L", gitea_zip_url, "-o", "package.zip")
	err = gitea_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitea_bin_url := "https://dl.gitea.com/gitea/1.22.2/gitea-src-1.22.2.bin"
	gitea_cmd_bin := exec.Command("curl", "-L", gitea_bin_url, "-o", "binary.bin")
	err = gitea_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitea_src_url := "https://dl.gitea.com/gitea/1.22.2/gitea-src-1.22.2.src.tar.gz"
	gitea_cmd_src := exec.Command("curl", "-L", gitea_src_url, "-o", "source.tar.gz")
	err = gitea_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitea_cmd_direct := exec.Command("./binary")
	err = gitea_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: yarn")
exec.Command("latte", "install", "yarn")
}
