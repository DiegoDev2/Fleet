package main

// PurescriptLanguageServer - Language Server Protocol server for PureScript
// Homepage: https://github.com/nwolverson/purescript-language-server

import (
	"fmt"
	
	"os/exec"
)

func installPurescriptLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	purescriptlanguageserver_tar_url := "https://registry.npmjs.org/purescript-language-server/-/purescript-language-server-0.18.0.tgz"
	purescriptlanguageserver_cmd_tar := exec.Command("curl", "-L", purescriptlanguageserver_tar_url, "-o", "package.tar.gz")
	err := purescriptlanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	purescriptlanguageserver_zip_url := "https://registry.npmjs.org/purescript-language-server/-/purescript-language-server-0.18.0.tgz"
	purescriptlanguageserver_cmd_zip := exec.Command("curl", "-L", purescriptlanguageserver_zip_url, "-o", "package.zip")
	err = purescriptlanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	purescriptlanguageserver_bin_url := "https://registry.npmjs.org/purescript-language-server/-/purescript-language-server-0.18.0.tgz"
	purescriptlanguageserver_cmd_bin := exec.Command("curl", "-L", purescriptlanguageserver_bin_url, "-o", "binary.bin")
	err = purescriptlanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	purescriptlanguageserver_src_url := "https://registry.npmjs.org/purescript-language-server/-/purescript-language-server-0.18.0.tgz"
	purescriptlanguageserver_cmd_src := exec.Command("curl", "-L", purescriptlanguageserver_src_url, "-o", "source.tar.gz")
	err = purescriptlanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	purescriptlanguageserver_cmd_direct := exec.Command("./binary")
	err = purescriptlanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: purescript")
	exec.Command("latte", "install", "purescript").Run()
}
