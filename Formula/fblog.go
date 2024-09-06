package main

// Fblog - Small command-line JSON log viewer
// Homepage: https://github.com/brocode/fblog

import (
	"fmt"
	
	"os/exec"
)

func installFblog() {
	// Método 1: Descargar y extraer .tar.gz
	fblog_tar_url := "https://github.com/brocode/fblog/archive/refs/tags/v4.10.0.tar.gz"
	fblog_cmd_tar := exec.Command("curl", "-L", fblog_tar_url, "-o", "package.tar.gz")
	err := fblog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fblog_zip_url := "https://github.com/brocode/fblog/archive/refs/tags/v4.10.0.zip"
	fblog_cmd_zip := exec.Command("curl", "-L", fblog_zip_url, "-o", "package.zip")
	err = fblog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fblog_bin_url := "https://github.com/brocode/fblog/archive/refs/tags/v4.10.0.bin"
	fblog_cmd_bin := exec.Command("curl", "-L", fblog_bin_url, "-o", "binary.bin")
	err = fblog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fblog_src_url := "https://github.com/brocode/fblog/archive/refs/tags/v4.10.0.src.tar.gz"
	fblog_cmd_src := exec.Command("curl", "-L", fblog_src_url, "-o", "source.tar.gz")
	err = fblog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fblog_cmd_direct := exec.Command("./binary")
	err = fblog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
