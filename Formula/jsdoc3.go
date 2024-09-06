package main

// Jsdoc3 - API documentation generator for JavaScript
// Homepage: https://jsdoc.app/

import (
	"fmt"
	
	"os/exec"
)

func installJsdoc3() {
	// Método 1: Descargar y extraer .tar.gz
	jsdoc3_tar_url := "https://registry.npmjs.org/jsdoc/-/jsdoc-4.0.3.tgz"
	jsdoc3_cmd_tar := exec.Command("curl", "-L", jsdoc3_tar_url, "-o", "package.tar.gz")
	err := jsdoc3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsdoc3_zip_url := "https://registry.npmjs.org/jsdoc/-/jsdoc-4.0.3.tgz"
	jsdoc3_cmd_zip := exec.Command("curl", "-L", jsdoc3_zip_url, "-o", "package.zip")
	err = jsdoc3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsdoc3_bin_url := "https://registry.npmjs.org/jsdoc/-/jsdoc-4.0.3.tgz"
	jsdoc3_cmd_bin := exec.Command("curl", "-L", jsdoc3_bin_url, "-o", "binary.bin")
	err = jsdoc3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsdoc3_src_url := "https://registry.npmjs.org/jsdoc/-/jsdoc-4.0.3.tgz"
	jsdoc3_cmd_src := exec.Command("curl", "-L", jsdoc3_src_url, "-o", "source.tar.gz")
	err = jsdoc3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsdoc3_cmd_direct := exec.Command("./binary")
	err = jsdoc3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
