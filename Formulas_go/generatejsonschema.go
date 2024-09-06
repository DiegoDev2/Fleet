package main

// GenerateJsonSchema - Generate a JSON Schema from Sample JSON
// Homepage: https://github.com/Nijikokun/generate-schema

import (
	"fmt"
	
	"os/exec"
)

func installGenerateJsonSchema() {
	// Método 1: Descargar y extraer .tar.gz
	generatejsonschema_tar_url := "https://registry.npmjs.org/generate-schema/-/generate-schema-2.6.0.tgz"
	generatejsonschema_cmd_tar := exec.Command("curl", "-L", generatejsonschema_tar_url, "-o", "package.tar.gz")
	err := generatejsonschema_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	generatejsonschema_zip_url := "https://registry.npmjs.org/generate-schema/-/generate-schema-2.6.0.tgz"
	generatejsonschema_cmd_zip := exec.Command("curl", "-L", generatejsonschema_zip_url, "-o", "package.zip")
	err = generatejsonschema_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	generatejsonschema_bin_url := "https://registry.npmjs.org/generate-schema/-/generate-schema-2.6.0.tgz"
	generatejsonschema_cmd_bin := exec.Command("curl", "-L", generatejsonschema_bin_url, "-o", "binary.bin")
	err = generatejsonschema_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	generatejsonschema_src_url := "https://registry.npmjs.org/generate-schema/-/generate-schema-2.6.0.tgz"
	generatejsonschema_cmd_src := exec.Command("curl", "-L", generatejsonschema_src_url, "-o", "source.tar.gz")
	err = generatejsonschema_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	generatejsonschema_cmd_direct := exec.Command("./binary")
	err = generatejsonschema_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
