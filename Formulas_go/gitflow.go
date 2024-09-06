package main

// GitFlow - Extensions to follow Vincent Driessen's branching model
// Homepage: https://github.com/nvie/gitflow

import (
	"fmt"
	
	"os/exec"
)

func installGitFlow() {
	// Método 1: Descargar y extraer .tar.gz
	gitflow_tar_url := "https://github.com/nvie/gitflow.git"
	gitflow_cmd_tar := exec.Command("curl", "-L", gitflow_tar_url, "-o", "package.tar.gz")
	err := gitflow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitflow_zip_url := "https://github.com/nvie/gitflow.git"
	gitflow_cmd_zip := exec.Command("curl", "-L", gitflow_zip_url, "-o", "package.zip")
	err = gitflow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitflow_bin_url := "https://github.com/nvie/gitflow.git"
	gitflow_cmd_bin := exec.Command("curl", "-L", gitflow_bin_url, "-o", "binary.bin")
	err = gitflow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitflow_src_url := "https://github.com/nvie/gitflow.git"
	gitflow_cmd_src := exec.Command("curl", "-L", gitflow_src_url, "-o", "source.tar.gz")
	err = gitflow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitflow_cmd_direct := exec.Command("./binary")
	err = gitflow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
