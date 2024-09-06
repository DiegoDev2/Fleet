package main

// Jsonnet - Domain specific configuration language for defining JSON data
// Homepage: https://jsonnet.org/

import (
	"fmt"
	
	"os/exec"
)

func installJsonnet() {
	// Método 1: Descargar y extraer .tar.gz
	jsonnet_tar_url := "https://github.com/google/jsonnet/archive/refs/tags/v0.20.0.tar.gz"
	jsonnet_cmd_tar := exec.Command("curl", "-L", jsonnet_tar_url, "-o", "package.tar.gz")
	err := jsonnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonnet_zip_url := "https://github.com/google/jsonnet/archive/refs/tags/v0.20.0.zip"
	jsonnet_cmd_zip := exec.Command("curl", "-L", jsonnet_zip_url, "-o", "package.zip")
	err = jsonnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonnet_bin_url := "https://github.com/google/jsonnet/archive/refs/tags/v0.20.0.bin"
	jsonnet_cmd_bin := exec.Command("curl", "-L", jsonnet_bin_url, "-o", "binary.bin")
	err = jsonnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonnet_src_url := "https://github.com/google/jsonnet/archive/refs/tags/v0.20.0.src.tar.gz"
	jsonnet_cmd_src := exec.Command("curl", "-L", jsonnet_src_url, "-o", "source.tar.gz")
	err = jsonnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonnet_cmd_direct := exec.Command("./binary")
	err = jsonnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
