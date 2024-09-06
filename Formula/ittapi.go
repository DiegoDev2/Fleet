package main

// Ittapi - Intel Instrumentation and Tracing Technology (ITT) and Just-In-Time (JIT) API
// Homepage: https://github.com/intel/ittapi

import (
	"fmt"
	
	"os/exec"
)

func installIttapi() {
	// Método 1: Descargar y extraer .tar.gz
	ittapi_tar_url := "https://github.com/intel/ittapi/archive/refs/tags/v3.25.2.tar.gz"
	ittapi_cmd_tar := exec.Command("curl", "-L", ittapi_tar_url, "-o", "package.tar.gz")
	err := ittapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ittapi_zip_url := "https://github.com/intel/ittapi/archive/refs/tags/v3.25.2.zip"
	ittapi_cmd_zip := exec.Command("curl", "-L", ittapi_zip_url, "-o", "package.zip")
	err = ittapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ittapi_bin_url := "https://github.com/intel/ittapi/archive/refs/tags/v3.25.2.bin"
	ittapi_cmd_bin := exec.Command("curl", "-L", ittapi_bin_url, "-o", "binary.bin")
	err = ittapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ittapi_src_url := "https://github.com/intel/ittapi/archive/refs/tags/v3.25.2.src.tar.gz"
	ittapi_cmd_src := exec.Command("curl", "-L", ittapi_src_url, "-o", "source.tar.gz")
	err = ittapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ittapi_cmd_direct := exec.Command("./binary")
	err = ittapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
