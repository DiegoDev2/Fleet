package main

// Alexjs - Catch insensitive, inconsiderate writing
// Homepage: https://alexjs.com

import (
	"fmt"
	
	"os/exec"
)

func installAlexjs() {
	// Método 1: Descargar y extraer .tar.gz
	alexjs_tar_url := "https://github.com/get-alex/alex/archive/refs/tags/11.0.1.tar.gz"
	alexjs_cmd_tar := exec.Command("curl", "-L", alexjs_tar_url, "-o", "package.tar.gz")
	err := alexjs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alexjs_zip_url := "https://github.com/get-alex/alex/archive/refs/tags/11.0.1.zip"
	alexjs_cmd_zip := exec.Command("curl", "-L", alexjs_zip_url, "-o", "package.zip")
	err = alexjs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alexjs_bin_url := "https://github.com/get-alex/alex/archive/refs/tags/11.0.1.bin"
	alexjs_cmd_bin := exec.Command("curl", "-L", alexjs_bin_url, "-o", "binary.bin")
	err = alexjs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alexjs_src_url := "https://github.com/get-alex/alex/archive/refs/tags/11.0.1.src.tar.gz"
	alexjs_cmd_src := exec.Command("curl", "-L", alexjs_src_url, "-o", "source.tar.gz")
	err = alexjs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alexjs_cmd_direct := exec.Command("./binary")
	err = alexjs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
