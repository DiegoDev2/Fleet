package main

// GitXargs - CLI for making updates across multiple Github repositories with a single command
// Homepage: https://github.com/gruntwork-io/git-xargs

import (
	"fmt"
	
	"os/exec"
)

func installGitXargs() {
	// Método 1: Descargar y extraer .tar.gz
	gitxargs_tar_url := "https://github.com/gruntwork-io/git-xargs/archive/refs/tags/v0.1.11.tar.gz"
	gitxargs_cmd_tar := exec.Command("curl", "-L", gitxargs_tar_url, "-o", "package.tar.gz")
	err := gitxargs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitxargs_zip_url := "https://github.com/gruntwork-io/git-xargs/archive/refs/tags/v0.1.11.zip"
	gitxargs_cmd_zip := exec.Command("curl", "-L", gitxargs_zip_url, "-o", "package.zip")
	err = gitxargs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitxargs_bin_url := "https://github.com/gruntwork-io/git-xargs/archive/refs/tags/v0.1.11.bin"
	gitxargs_cmd_bin := exec.Command("curl", "-L", gitxargs_bin_url, "-o", "binary.bin")
	err = gitxargs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitxargs_src_url := "https://github.com/gruntwork-io/git-xargs/archive/refs/tags/v0.1.11.src.tar.gz"
	gitxargs_cmd_src := exec.Command("curl", "-L", gitxargs_src_url, "-o", "source.tar.gz")
	err = gitxargs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitxargs_cmd_direct := exec.Command("./binary")
	err = gitxargs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
