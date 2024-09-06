package main

// Roapi - Full-fledged APIs for static datasets without writing a single line of code
// Homepage: https://roapi.github.io/docs

import (
	"fmt"
	
	"os/exec"
)

func installRoapi() {
	// Método 1: Descargar y extraer .tar.gz
	roapi_tar_url := "https://github.com/roapi/roapi/archive/refs/tags/roapi-v0.12.0.tar.gz"
	roapi_cmd_tar := exec.Command("curl", "-L", roapi_tar_url, "-o", "package.tar.gz")
	err := roapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	roapi_zip_url := "https://github.com/roapi/roapi/archive/refs/tags/roapi-v0.12.0.zip"
	roapi_cmd_zip := exec.Command("curl", "-L", roapi_zip_url, "-o", "package.zip")
	err = roapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	roapi_bin_url := "https://github.com/roapi/roapi/archive/refs/tags/roapi-v0.12.0.bin"
	roapi_cmd_bin := exec.Command("curl", "-L", roapi_bin_url, "-o", "binary.bin")
	err = roapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	roapi_src_url := "https://github.com/roapi/roapi/archive/refs/tags/roapi-v0.12.0.src.tar.gz"
	roapi_cmd_src := exec.Command("curl", "-L", roapi_src_url, "-o", "source.tar.gz")
	err = roapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	roapi_cmd_direct := exec.Command("./binary")
	err = roapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
