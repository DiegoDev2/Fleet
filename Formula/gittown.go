package main

// GitTown - High-level command-line interface for Git
// Homepage: https://www.git-town.com/

import (
	"fmt"
	
	"os/exec"
)

func installGitTown() {
	// Método 1: Descargar y extraer .tar.gz
	gittown_tar_url := "https://github.com/git-town/git-town/archive/refs/tags/v16.0.0.tar.gz"
	gittown_cmd_tar := exec.Command("curl", "-L", gittown_tar_url, "-o", "package.tar.gz")
	err := gittown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gittown_zip_url := "https://github.com/git-town/git-town/archive/refs/tags/v16.0.0.zip"
	gittown_cmd_zip := exec.Command("curl", "-L", gittown_zip_url, "-o", "package.zip")
	err = gittown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gittown_bin_url := "https://github.com/git-town/git-town/archive/refs/tags/v16.0.0.bin"
	gittown_cmd_bin := exec.Command("curl", "-L", gittown_bin_url, "-o", "binary.bin")
	err = gittown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gittown_src_url := "https://github.com/git-town/git-town/archive/refs/tags/v16.0.0.src.tar.gz"
	gittown_cmd_src := exec.Command("curl", "-L", gittown_src_url, "-o", "source.tar.gz")
	err = gittown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gittown_cmd_direct := exec.Command("./binary")
	err = gittown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
