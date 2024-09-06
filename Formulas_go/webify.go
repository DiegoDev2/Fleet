package main

// Webify - Wrapper for shell commands as web services
// Homepage: https://github.com/beefsack/webify

import (
	"fmt"
	
	"os/exec"
)

func installWebify() {
	// Método 1: Descargar y extraer .tar.gz
	webify_tar_url := "https://github.com/beefsack/webify/archive/refs/tags/v1.5.0.tar.gz"
	webify_cmd_tar := exec.Command("curl", "-L", webify_tar_url, "-o", "package.tar.gz")
	err := webify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webify_zip_url := "https://github.com/beefsack/webify/archive/refs/tags/v1.5.0.zip"
	webify_cmd_zip := exec.Command("curl", "-L", webify_zip_url, "-o", "package.zip")
	err = webify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webify_bin_url := "https://github.com/beefsack/webify/archive/refs/tags/v1.5.0.bin"
	webify_cmd_bin := exec.Command("curl", "-L", webify_bin_url, "-o", "binary.bin")
	err = webify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webify_src_url := "https://github.com/beefsack/webify/archive/refs/tags/v1.5.0.src.tar.gz"
	webify_cmd_src := exec.Command("curl", "-L", webify_src_url, "-o", "source.tar.gz")
	err = webify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webify_cmd_direct := exec.Command("./binary")
	err = webify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
