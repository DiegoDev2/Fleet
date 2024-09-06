package main

// Copa - Tool to directly patch container images given the vulnerability scanning results
// Homepage: https://github.com/project-copacetic/copacetic

import (
	"fmt"
	
	"os/exec"
)

func installCopa() {
	// Método 1: Descargar y extraer .tar.gz
	copa_tar_url := "https://github.com/project-copacetic/copacetic/archive/refs/tags/v0.7.0.tar.gz"
	copa_cmd_tar := exec.Command("curl", "-L", copa_tar_url, "-o", "package.tar.gz")
	err := copa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	copa_zip_url := "https://github.com/project-copacetic/copacetic/archive/refs/tags/v0.7.0.zip"
	copa_cmd_zip := exec.Command("curl", "-L", copa_zip_url, "-o", "package.zip")
	err = copa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	copa_bin_url := "https://github.com/project-copacetic/copacetic/archive/refs/tags/v0.7.0.bin"
	copa_cmd_bin := exec.Command("curl", "-L", copa_bin_url, "-o", "binary.bin")
	err = copa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	copa_src_url := "https://github.com/project-copacetic/copacetic/archive/refs/tags/v0.7.0.src.tar.gz"
	copa_cmd_src := exec.Command("curl", "-L", copa_src_url, "-o", "source.tar.gz")
	err = copa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	copa_cmd_direct := exec.Command("./binary")
	err = copa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
