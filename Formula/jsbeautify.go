package main

// JsBeautify - JavaScript, CSS and HTML unobfuscator and beautifier
// Homepage: https://beautifier.io

import (
	"fmt"
	
	"os/exec"
)

func installJsBeautify() {
	// Método 1: Descargar y extraer .tar.gz
	jsbeautify_tar_url := "https://registry.npmjs.org/js-beautify/-/js-beautify-1.15.1.tgz"
	jsbeautify_cmd_tar := exec.Command("curl", "-L", jsbeautify_tar_url, "-o", "package.tar.gz")
	err := jsbeautify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsbeautify_zip_url := "https://registry.npmjs.org/js-beautify/-/js-beautify-1.15.1.tgz"
	jsbeautify_cmd_zip := exec.Command("curl", "-L", jsbeautify_zip_url, "-o", "package.zip")
	err = jsbeautify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsbeautify_bin_url := "https://registry.npmjs.org/js-beautify/-/js-beautify-1.15.1.tgz"
	jsbeautify_cmd_bin := exec.Command("curl", "-L", jsbeautify_bin_url, "-o", "binary.bin")
	err = jsbeautify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsbeautify_src_url := "https://registry.npmjs.org/js-beautify/-/js-beautify-1.15.1.tgz"
	jsbeautify_cmd_src := exec.Command("curl", "-L", jsbeautify_src_url, "-o", "source.tar.gz")
	err = jsbeautify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsbeautify_cmd_direct := exec.Command("./binary")
	err = jsbeautify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
