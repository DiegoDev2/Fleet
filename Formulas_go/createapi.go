package main

// CreateApi - Delightful code generator for OpenAPI specs
// Homepage: https://github.com/CreateAPI/CreateAPI

import (
	"fmt"
	
	"os/exec"
)

func installCreateApi() {
	// Método 1: Descargar y extraer .tar.gz
	createapi_tar_url := "https://github.com/CreateAPI/CreateAPI/archive/refs/tags/0.2.0.tar.gz"
	createapi_cmd_tar := exec.Command("curl", "-L", createapi_tar_url, "-o", "package.tar.gz")
	err := createapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	createapi_zip_url := "https://github.com/CreateAPI/CreateAPI/archive/refs/tags/0.2.0.zip"
	createapi_cmd_zip := exec.Command("curl", "-L", createapi_zip_url, "-o", "package.zip")
	err = createapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	createapi_bin_url := "https://github.com/CreateAPI/CreateAPI/archive/refs/tags/0.2.0.bin"
	createapi_cmd_bin := exec.Command("curl", "-L", createapi_bin_url, "-o", "binary.bin")
	err = createapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	createapi_src_url := "https://github.com/CreateAPI/CreateAPI/archive/refs/tags/0.2.0.src.tar.gz"
	createapi_cmd_src := exec.Command("curl", "-L", createapi_src_url, "-o", "source.tar.gz")
	err = createapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	createapi_cmd_direct := exec.Command("./binary")
	err = createapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
