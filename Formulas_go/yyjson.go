package main

// Yyjson - High performance JSON library written in ANSI C
// Homepage: https://github.com/ibireme/yyjson

import (
	"fmt"
	
	"os/exec"
)

func installYyjson() {
	// Método 1: Descargar y extraer .tar.gz
	yyjson_tar_url := "https://github.com/ibireme/yyjson/archive/refs/tags/0.10.0.tar.gz"
	yyjson_cmd_tar := exec.Command("curl", "-L", yyjson_tar_url, "-o", "package.tar.gz")
	err := yyjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yyjson_zip_url := "https://github.com/ibireme/yyjson/archive/refs/tags/0.10.0.zip"
	yyjson_cmd_zip := exec.Command("curl", "-L", yyjson_zip_url, "-o", "package.zip")
	err = yyjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yyjson_bin_url := "https://github.com/ibireme/yyjson/archive/refs/tags/0.10.0.bin"
	yyjson_cmd_bin := exec.Command("curl", "-L", yyjson_bin_url, "-o", "binary.bin")
	err = yyjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yyjson_src_url := "https://github.com/ibireme/yyjson/archive/refs/tags/0.10.0.src.tar.gz"
	yyjson_cmd_src := exec.Command("curl", "-L", yyjson_src_url, "-o", "source.tar.gz")
	err = yyjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yyjson_cmd_direct := exec.Command("./binary")
	err = yyjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
