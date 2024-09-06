package main

// Jsonschema2pojo - Generates Java types from JSON Schema (or example JSON)
// Homepage: https://www.jsonschema2pojo.org/

import (
	"fmt"
	
	"os/exec"
)

func installJsonschema2pojo() {
	// Método 1: Descargar y extraer .tar.gz
	jsonschema2pojo_tar_url := "https://github.com/joelittlejohn/jsonschema2pojo/releases/download/jsonschema2pojo-1.2.1/jsonschema2pojo-1.2.1.tar.gz"
	jsonschema2pojo_cmd_tar := exec.Command("curl", "-L", jsonschema2pojo_tar_url, "-o", "package.tar.gz")
	err := jsonschema2pojo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonschema2pojo_zip_url := "https://github.com/joelittlejohn/jsonschema2pojo/releases/download/jsonschema2pojo-1.2.1/jsonschema2pojo-1.2.1.zip"
	jsonschema2pojo_cmd_zip := exec.Command("curl", "-L", jsonschema2pojo_zip_url, "-o", "package.zip")
	err = jsonschema2pojo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonschema2pojo_bin_url := "https://github.com/joelittlejohn/jsonschema2pojo/releases/download/jsonschema2pojo-1.2.1/jsonschema2pojo-1.2.1.bin"
	jsonschema2pojo_cmd_bin := exec.Command("curl", "-L", jsonschema2pojo_bin_url, "-o", "binary.bin")
	err = jsonschema2pojo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonschema2pojo_src_url := "https://github.com/joelittlejohn/jsonschema2pojo/releases/download/jsonschema2pojo-1.2.1/jsonschema2pojo-1.2.1.src.tar.gz"
	jsonschema2pojo_cmd_src := exec.Command("curl", "-L", jsonschema2pojo_src_url, "-o", "source.tar.gz")
	err = jsonschema2pojo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonschema2pojo_cmd_direct := exec.Command("./binary")
	err = jsonschema2pojo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
