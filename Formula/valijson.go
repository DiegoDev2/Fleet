package main

// Valijson - Header-only C++ library for JSON Schema validation
// Homepage: https://github.com/tristanpenman/valijson

import (
	"fmt"
	
	"os/exec"
)

func installValijson() {
	// Método 1: Descargar y extraer .tar.gz
	valijson_tar_url := "https://github.com/tristanpenman/valijson/archive/refs/tags/v1.0.3.tar.gz"
	valijson_cmd_tar := exec.Command("curl", "-L", valijson_tar_url, "-o", "package.tar.gz")
	err := valijson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	valijson_zip_url := "https://github.com/tristanpenman/valijson/archive/refs/tags/v1.0.3.zip"
	valijson_cmd_zip := exec.Command("curl", "-L", valijson_zip_url, "-o", "package.zip")
	err = valijson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	valijson_bin_url := "https://github.com/tristanpenman/valijson/archive/refs/tags/v1.0.3.bin"
	valijson_cmd_bin := exec.Command("curl", "-L", valijson_bin_url, "-o", "binary.bin")
	err = valijson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	valijson_src_url := "https://github.com/tristanpenman/valijson/archive/refs/tags/v1.0.3.src.tar.gz"
	valijson_cmd_src := exec.Command("curl", "-L", valijson_src_url, "-o", "source.tar.gz")
	err = valijson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	valijson_cmd_direct := exec.Command("./binary")
	err = valijson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: jsoncpp")
	exec.Command("latte", "install", "jsoncpp").Run()
}
