package main

// GoJsonnet - Go implementation of configuration language for defining JSON data
// Homepage: https://jsonnet.org/

import (
	"fmt"
	
	"os/exec"
)

func installGoJsonnet() {
	// Método 1: Descargar y extraer .tar.gz
	gojsonnet_tar_url := "https://github.com/google/go-jsonnet/archive/refs/tags/v0.20.0.tar.gz"
	gojsonnet_cmd_tar := exec.Command("curl", "-L", gojsonnet_tar_url, "-o", "package.tar.gz")
	err := gojsonnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gojsonnet_zip_url := "https://github.com/google/go-jsonnet/archive/refs/tags/v0.20.0.zip"
	gojsonnet_cmd_zip := exec.Command("curl", "-L", gojsonnet_zip_url, "-o", "package.zip")
	err = gojsonnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gojsonnet_bin_url := "https://github.com/google/go-jsonnet/archive/refs/tags/v0.20.0.bin"
	gojsonnet_cmd_bin := exec.Command("curl", "-L", gojsonnet_bin_url, "-o", "binary.bin")
	err = gojsonnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gojsonnet_src_url := "https://github.com/google/go-jsonnet/archive/refs/tags/v0.20.0.src.tar.gz"
	gojsonnet_cmd_src := exec.Command("curl", "-L", gojsonnet_src_url, "-o", "source.tar.gz")
	err = gojsonnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gojsonnet_cmd_direct := exec.Command("./binary")
	err = gojsonnet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
