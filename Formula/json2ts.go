package main

// Json2ts - Compile JSONSchema to TypeScript type declarations
// Homepage: https://github.com/bcherny/json-schema-to-typescript

import (
	"fmt"
	
	"os/exec"
)

func installJson2ts() {
	// Método 1: Descargar y extraer .tar.gz
	json2ts_tar_url := "https://registry.npmjs.org/json-schema-to-typescript/-/json-schema-to-typescript-15.0.2.tgz"
	json2ts_cmd_tar := exec.Command("curl", "-L", json2ts_tar_url, "-o", "package.tar.gz")
	err := json2ts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	json2ts_zip_url := "https://registry.npmjs.org/json-schema-to-typescript/-/json-schema-to-typescript-15.0.2.tgz"
	json2ts_cmd_zip := exec.Command("curl", "-L", json2ts_zip_url, "-o", "package.zip")
	err = json2ts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	json2ts_bin_url := "https://registry.npmjs.org/json-schema-to-typescript/-/json-schema-to-typescript-15.0.2.tgz"
	json2ts_cmd_bin := exec.Command("curl", "-L", json2ts_bin_url, "-o", "binary.bin")
	err = json2ts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	json2ts_src_url := "https://registry.npmjs.org/json-schema-to-typescript/-/json-schema-to-typescript-15.0.2.tgz"
	json2ts_cmd_src := exec.Command("curl", "-L", json2ts_src_url, "-o", "source.tar.gz")
	err = json2ts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	json2ts_cmd_direct := exec.Command("./binary")
	err = json2ts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
