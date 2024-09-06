package main

// BashLanguageServer - Language Server for Bash
// Homepage: https://github.com/bash-lsp/bash-language-server

import (
	"fmt"
	
	"os/exec"
)

func installBashLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	bashlanguageserver_tar_url := "https://registry.npmjs.org/bash-language-server/-/bash-language-server-5.4.0.tgz"
	bashlanguageserver_cmd_tar := exec.Command("curl", "-L", bashlanguageserver_tar_url, "-o", "package.tar.gz")
	err := bashlanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashlanguageserver_zip_url := "https://registry.npmjs.org/bash-language-server/-/bash-language-server-5.4.0.tgz"
	bashlanguageserver_cmd_zip := exec.Command("curl", "-L", bashlanguageserver_zip_url, "-o", "package.zip")
	err = bashlanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashlanguageserver_bin_url := "https://registry.npmjs.org/bash-language-server/-/bash-language-server-5.4.0.tgz"
	bashlanguageserver_cmd_bin := exec.Command("curl", "-L", bashlanguageserver_bin_url, "-o", "binary.bin")
	err = bashlanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashlanguageserver_src_url := "https://registry.npmjs.org/bash-language-server/-/bash-language-server-5.4.0.tgz"
	bashlanguageserver_cmd_src := exec.Command("curl", "-L", bashlanguageserver_src_url, "-o", "source.tar.gz")
	err = bashlanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashlanguageserver_cmd_direct := exec.Command("./binary")
	err = bashlanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
