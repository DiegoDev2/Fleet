package main

// Jsonschema - Implementation of JSON Schema for Python
// Homepage: https://github.com/python-jsonschema/jsonschema

import (
	"fmt"
	
	"os/exec"
)

func installJsonschema() {
	// Método 1: Descargar y extraer .tar.gz
	jsonschema_tar_url := "https://files.pythonhosted.org/packages/36/3d/ca032d5ac064dff543aa13c984737795ac81abc9fb130cd2fcff17cfabc7/jsonschema-4.17.3.tar.gz"
	jsonschema_cmd_tar := exec.Command("curl", "-L", jsonschema_tar_url, "-o", "package.tar.gz")
	err := jsonschema_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonschema_zip_url := "https://files.pythonhosted.org/packages/36/3d/ca032d5ac064dff543aa13c984737795ac81abc9fb130cd2fcff17cfabc7/jsonschema-4.17.3.zip"
	jsonschema_cmd_zip := exec.Command("curl", "-L", jsonschema_zip_url, "-o", "package.zip")
	err = jsonschema_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonschema_bin_url := "https://files.pythonhosted.org/packages/36/3d/ca032d5ac064dff543aa13c984737795ac81abc9fb130cd2fcff17cfabc7/jsonschema-4.17.3.bin"
	jsonschema_cmd_bin := exec.Command("curl", "-L", jsonschema_bin_url, "-o", "binary.bin")
	err = jsonschema_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonschema_src_url := "https://files.pythonhosted.org/packages/36/3d/ca032d5ac064dff543aa13c984737795ac81abc9fb130cd2fcff17cfabc7/jsonschema-4.17.3.src.tar.gz"
	jsonschema_cmd_src := exec.Command("curl", "-L", jsonschema_src_url, "-o", "source.tar.gz")
	err = jsonschema_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonschema_cmd_direct := exec.Command("./binary")
	err = jsonschema_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
}
