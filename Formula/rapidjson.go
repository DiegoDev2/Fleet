package main

// Rapidjson - JSON parser/generator for C++ with SAX and DOM style APIs
// Homepage: https://rapidjson.org/

import (
	"fmt"
	
	"os/exec"
)

func installRapidjson() {
	// Método 1: Descargar y extraer .tar.gz
	rapidjson_tar_url := "https://github.com/Tencent/rapidjson/archive/refs/tags/v1.1.0.tar.gz"
	rapidjson_cmd_tar := exec.Command("curl", "-L", rapidjson_tar_url, "-o", "package.tar.gz")
	err := rapidjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rapidjson_zip_url := "https://github.com/Tencent/rapidjson/archive/refs/tags/v1.1.0.zip"
	rapidjson_cmd_zip := exec.Command("curl", "-L", rapidjson_zip_url, "-o", "package.zip")
	err = rapidjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rapidjson_bin_url := "https://github.com/Tencent/rapidjson/archive/refs/tags/v1.1.0.bin"
	rapidjson_cmd_bin := exec.Command("curl", "-L", rapidjson_bin_url, "-o", "binary.bin")
	err = rapidjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rapidjson_src_url := "https://github.com/Tencent/rapidjson/archive/refs/tags/v1.1.0.src.tar.gz"
	rapidjson_cmd_src := exec.Command("curl", "-L", rapidjson_src_url, "-o", "source.tar.gz")
	err = rapidjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rapidjson_cmd_direct := exec.Command("./binary")
	err = rapidjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
