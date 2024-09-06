package main

// JsonSpirit - C++ JSON parser/generator
// Homepage: https://www.codeproject.com/Articles/20027/JSON-Spirit-A-C-JSON-Parser-Generator-Implemented

import (
	"fmt"
	
	"os/exec"
)

func installJsonSpirit() {
	// Método 1: Descargar y extraer .tar.gz
	jsonspirit_tar_url := "https://github.com/png85/json_spirit/archive/refs/tags/json_spirit-4.0.8.tar.gz"
	jsonspirit_cmd_tar := exec.Command("curl", "-L", jsonspirit_tar_url, "-o", "package.tar.gz")
	err := jsonspirit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonspirit_zip_url := "https://github.com/png85/json_spirit/archive/refs/tags/json_spirit-4.0.8.zip"
	jsonspirit_cmd_zip := exec.Command("curl", "-L", jsonspirit_zip_url, "-o", "package.zip")
	err = jsonspirit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonspirit_bin_url := "https://github.com/png85/json_spirit/archive/refs/tags/json_spirit-4.0.8.bin"
	jsonspirit_cmd_bin := exec.Command("curl", "-L", jsonspirit_bin_url, "-o", "binary.bin")
	err = jsonspirit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonspirit_src_url := "https://github.com/png85/json_spirit/archive/refs/tags/json_spirit-4.0.8.src.tar.gz"
	jsonspirit_cmd_src := exec.Command("curl", "-L", jsonspirit_src_url, "-o", "source.tar.gz")
	err = jsonspirit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonspirit_cmd_direct := exec.Command("./binary")
	err = jsonspirit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
