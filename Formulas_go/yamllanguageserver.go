package main

// YamlLanguageServer - Language Server for Yaml Files
// Homepage: https://github.com/redhat-developer/yaml-language-server

import (
	"fmt"
	
	"os/exec"
)

func installYamlLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	yamllanguageserver_tar_url := "https://registry.npmjs.org/yaml-language-server/-/yaml-language-server-1.15.0.tgz"
	yamllanguageserver_cmd_tar := exec.Command("curl", "-L", yamllanguageserver_tar_url, "-o", "package.tar.gz")
	err := yamllanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yamllanguageserver_zip_url := "https://registry.npmjs.org/yaml-language-server/-/yaml-language-server-1.15.0.tgz"
	yamllanguageserver_cmd_zip := exec.Command("curl", "-L", yamllanguageserver_zip_url, "-o", "package.zip")
	err = yamllanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yamllanguageserver_bin_url := "https://registry.npmjs.org/yaml-language-server/-/yaml-language-server-1.15.0.tgz"
	yamllanguageserver_cmd_bin := exec.Command("curl", "-L", yamllanguageserver_bin_url, "-o", "binary.bin")
	err = yamllanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yamllanguageserver_src_url := "https://registry.npmjs.org/yaml-language-server/-/yaml-language-server-1.15.0.tgz"
	yamllanguageserver_cmd_src := exec.Command("curl", "-L", yamllanguageserver_src_url, "-o", "source.tar.gz")
	err = yamllanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yamllanguageserver_cmd_direct := exec.Command("./binary")
	err = yamllanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
