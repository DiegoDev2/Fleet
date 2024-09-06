package main

// Plog - Portable, simple and extensible C++ logging library
// Homepage: https://github.com/SergiusTheBest/plog

import (
	"fmt"
	
	"os/exec"
)

func installPlog() {
	// Método 1: Descargar y extraer .tar.gz
	plog_tar_url := "https://github.com/SergiusTheBest/plog/archive/refs/tags/1.1.10.tar.gz"
	plog_cmd_tar := exec.Command("curl", "-L", plog_tar_url, "-o", "package.tar.gz")
	err := plog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plog_zip_url := "https://github.com/SergiusTheBest/plog/archive/refs/tags/1.1.10.zip"
	plog_cmd_zip := exec.Command("curl", "-L", plog_zip_url, "-o", "package.zip")
	err = plog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plog_bin_url := "https://github.com/SergiusTheBest/plog/archive/refs/tags/1.1.10.bin"
	plog_cmd_bin := exec.Command("curl", "-L", plog_bin_url, "-o", "binary.bin")
	err = plog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plog_src_url := "https://github.com/SergiusTheBest/plog/archive/refs/tags/1.1.10.src.tar.gz"
	plog_cmd_src := exec.Command("curl", "-L", plog_src_url, "-o", "source.tar.gz")
	err = plog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plog_cmd_direct := exec.Command("./binary")
	err = plog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
