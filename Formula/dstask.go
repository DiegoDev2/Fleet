package main

// Dstask - Git-powered personal task tracker
// Homepage: https://github.com/naggie/dstask

import (
	"fmt"
	
	"os/exec"
)

func installDstask() {
	// Método 1: Descargar y extraer .tar.gz
	dstask_tar_url := "https://github.com/naggie/dstask/archive/refs/tags/v0.26.tar.gz"
	dstask_cmd_tar := exec.Command("curl", "-L", dstask_tar_url, "-o", "package.tar.gz")
	err := dstask_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dstask_zip_url := "https://github.com/naggie/dstask/archive/refs/tags/v0.26.zip"
	dstask_cmd_zip := exec.Command("curl", "-L", dstask_zip_url, "-o", "package.zip")
	err = dstask_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dstask_bin_url := "https://github.com/naggie/dstask/archive/refs/tags/v0.26.bin"
	dstask_cmd_bin := exec.Command("curl", "-L", dstask_bin_url, "-o", "binary.bin")
	err = dstask_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dstask_src_url := "https://github.com/naggie/dstask/archive/refs/tags/v0.26.src.tar.gz"
	dstask_cmd_src := exec.Command("curl", "-L", dstask_src_url, "-o", "source.tar.gz")
	err = dstask_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dstask_cmd_direct := exec.Command("./binary")
	err = dstask_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
