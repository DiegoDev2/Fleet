package main

// GitFlowAvh - AVH edition of git-flow
// Homepage: https://github.com/petervanderdoes/gitflow-avh

import (
	"fmt"
	
	"os/exec"
)

func installGitFlowAvh() {
	// Método 1: Descargar y extraer .tar.gz
	gitflowavh_tar_url := "https://github.com/petervanderdoes/gitflow-avh/archive/refs/tags/1.12.3.tar.gz"
	gitflowavh_cmd_tar := exec.Command("curl", "-L", gitflowavh_tar_url, "-o", "package.tar.gz")
	err := gitflowavh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitflowavh_zip_url := "https://github.com/petervanderdoes/gitflow-avh/archive/refs/tags/1.12.3.zip"
	gitflowavh_cmd_zip := exec.Command("curl", "-L", gitflowavh_zip_url, "-o", "package.zip")
	err = gitflowavh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitflowavh_bin_url := "https://github.com/petervanderdoes/gitflow-avh/archive/refs/tags/1.12.3.bin"
	gitflowavh_cmd_bin := exec.Command("curl", "-L", gitflowavh_bin_url, "-o", "binary.bin")
	err = gitflowavh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitflowavh_src_url := "https://github.com/petervanderdoes/gitflow-avh/archive/refs/tags/1.12.3.src.tar.gz"
	gitflowavh_cmd_src := exec.Command("curl", "-L", gitflowavh_src_url, "-o", "source.tar.gz")
	err = gitflowavh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitflowavh_cmd_direct := exec.Command("./binary")
	err = gitflowavh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-getopt")
exec.Command("latte", "install", "gnu-getopt")
}
