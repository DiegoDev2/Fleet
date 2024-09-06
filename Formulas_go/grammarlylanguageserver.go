package main

// GrammarlyLanguageserver - Language Server for Grammarly
// Homepage: https://github.com/znck/grammarly

import (
	"fmt"
	
	"os/exec"
)

func installGrammarlyLanguageserver() {
	// Método 1: Descargar y extraer .tar.gz
	grammarlylanguageserver_tar_url := "https://registry.npmjs.org/grammarly-languageserver/-/grammarly-languageserver-0.0.4.tgz"
	grammarlylanguageserver_cmd_tar := exec.Command("curl", "-L", grammarlylanguageserver_tar_url, "-o", "package.tar.gz")
	err := grammarlylanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grammarlylanguageserver_zip_url := "https://registry.npmjs.org/grammarly-languageserver/-/grammarly-languageserver-0.0.4.tgz"
	grammarlylanguageserver_cmd_zip := exec.Command("curl", "-L", grammarlylanguageserver_zip_url, "-o", "package.zip")
	err = grammarlylanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grammarlylanguageserver_bin_url := "https://registry.npmjs.org/grammarly-languageserver/-/grammarly-languageserver-0.0.4.tgz"
	grammarlylanguageserver_cmd_bin := exec.Command("curl", "-L", grammarlylanguageserver_bin_url, "-o", "binary.bin")
	err = grammarlylanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grammarlylanguageserver_src_url := "https://registry.npmjs.org/grammarly-languageserver/-/grammarly-languageserver-0.0.4.tgz"
	grammarlylanguageserver_cmd_src := exec.Command("curl", "-L", grammarlylanguageserver_src_url, "-o", "source.tar.gz")
	err = grammarlylanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grammarlylanguageserver_cmd_direct := exec.Command("./binary")
	err = grammarlylanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node@16")
exec.Command("latte", "install", "node@16")
}
