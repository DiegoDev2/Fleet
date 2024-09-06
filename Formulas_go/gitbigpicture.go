package main

// GitBigPicture - Visualization tool for Git repositories
// Homepage: https://github.com/git-big-picture/git-big-picture

import (
	"fmt"
	
	"os/exec"
)

func installGitBigPicture() {
	// Método 1: Descargar y extraer .tar.gz
	gitbigpicture_tar_url := "https://github.com/git-big-picture/git-big-picture/archive/refs/tags/v1.3.0.tar.gz"
	gitbigpicture_cmd_tar := exec.Command("curl", "-L", gitbigpicture_tar_url, "-o", "package.tar.gz")
	err := gitbigpicture_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitbigpicture_zip_url := "https://github.com/git-big-picture/git-big-picture/archive/refs/tags/v1.3.0.zip"
	gitbigpicture_cmd_zip := exec.Command("curl", "-L", gitbigpicture_zip_url, "-o", "package.zip")
	err = gitbigpicture_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitbigpicture_bin_url := "https://github.com/git-big-picture/git-big-picture/archive/refs/tags/v1.3.0.bin"
	gitbigpicture_cmd_bin := exec.Command("curl", "-L", gitbigpicture_bin_url, "-o", "binary.bin")
	err = gitbigpicture_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitbigpicture_src_url := "https://github.com/git-big-picture/git-big-picture/archive/refs/tags/v1.3.0.src.tar.gz"
	gitbigpicture_cmd_src := exec.Command("curl", "-L", gitbigpicture_src_url, "-o", "source.tar.gz")
	err = gitbigpicture_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitbigpicture_cmd_direct := exec.Command("./binary")
	err = gitbigpicture_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
