package main

// GitHooksGo - Git hooks manager
// Homepage: https://git-hooks.github.io/git-hooks

import (
	"fmt"
	
	"os/exec"
)

func installGitHooksGo() {
	// Método 1: Descargar y extraer .tar.gz
	githooksgo_tar_url := "https://github.com/git-hooks/git-hooks/archive/refs/tags/v1.3.1.tar.gz"
	githooksgo_cmd_tar := exec.Command("curl", "-L", githooksgo_tar_url, "-o", "package.tar.gz")
	err := githooksgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	githooksgo_zip_url := "https://github.com/git-hooks/git-hooks/archive/refs/tags/v1.3.1.zip"
	githooksgo_cmd_zip := exec.Command("curl", "-L", githooksgo_zip_url, "-o", "package.zip")
	err = githooksgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	githooksgo_bin_url := "https://github.com/git-hooks/git-hooks/archive/refs/tags/v1.3.1.bin"
	githooksgo_cmd_bin := exec.Command("curl", "-L", githooksgo_bin_url, "-o", "binary.bin")
	err = githooksgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	githooksgo_src_url := "https://github.com/git-hooks/git-hooks/archive/refs/tags/v1.3.1.src.tar.gz"
	githooksgo_cmd_src := exec.Command("curl", "-L", githooksgo_src_url, "-o", "source.tar.gz")
	err = githooksgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	githooksgo_cmd_direct := exec.Command("./binary")
	err = githooksgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
