package main

// Simdjson - SIMD-accelerated C++ JSON parser
// Homepage: https://simdjson.org

import (
	"fmt"
	
	"os/exec"
)

func installSimdjson() {
	// Método 1: Descargar y extraer .tar.gz
	simdjson_tar_url := "https://github.com/simdjson/simdjson/archive/refs/tags/v3.10.1.tar.gz"
	simdjson_cmd_tar := exec.Command("curl", "-L", simdjson_tar_url, "-o", "package.tar.gz")
	err := simdjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simdjson_zip_url := "https://github.com/simdjson/simdjson/archive/refs/tags/v3.10.1.zip"
	simdjson_cmd_zip := exec.Command("curl", "-L", simdjson_zip_url, "-o", "package.zip")
	err = simdjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simdjson_bin_url := "https://github.com/simdjson/simdjson/archive/refs/tags/v3.10.1.bin"
	simdjson_cmd_bin := exec.Command("curl", "-L", simdjson_bin_url, "-o", "binary.bin")
	err = simdjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simdjson_src_url := "https://github.com/simdjson/simdjson/archive/refs/tags/v3.10.1.src.tar.gz"
	simdjson_cmd_src := exec.Command("curl", "-L", simdjson_src_url, "-o", "source.tar.gz")
	err = simdjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simdjson_cmd_direct := exec.Command("./binary")
	err = simdjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
