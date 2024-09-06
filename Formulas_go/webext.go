package main

// WebExt - Command-line tool to help build, run, and test web extensions
// Homepage: https://github.com/mozilla/web-ext

import (
	"fmt"
	
	"os/exec"
)

func installWebExt() {
	// Método 1: Descargar y extraer .tar.gz
	webext_tar_url := "https://registry.npmjs.org/web-ext/-/web-ext-8.2.0.tgz"
	webext_cmd_tar := exec.Command("curl", "-L", webext_tar_url, "-o", "package.tar.gz")
	err := webext_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webext_zip_url := "https://registry.npmjs.org/web-ext/-/web-ext-8.2.0.tgz"
	webext_cmd_zip := exec.Command("curl", "-L", webext_zip_url, "-o", "package.zip")
	err = webext_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webext_bin_url := "https://registry.npmjs.org/web-ext/-/web-ext-8.2.0.tgz"
	webext_cmd_bin := exec.Command("curl", "-L", webext_bin_url, "-o", "binary.bin")
	err = webext_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webext_src_url := "https://registry.npmjs.org/web-ext/-/web-ext-8.2.0.tgz"
	webext_cmd_src := exec.Command("curl", "-L", webext_src_url, "-o", "source.tar.gz")
	err = webext_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webext_cmd_direct := exec.Command("./binary")
	err = webext_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: terminal-notifier")
exec.Command("latte", "install", "terminal-notifier")
}
