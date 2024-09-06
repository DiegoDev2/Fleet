package main

// DockerfileLanguageServer - Language server for Dockerfiles powered by Node, TypeScript, and VSCode
// Homepage: https://github.com/rcjsuen/dockerfile-language-server

import (
	"fmt"
	
	"os/exec"
)

func installDockerfileLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	dockerfilelanguageserver_tar_url := "https://registry.npmjs.org/dockerfile-language-server-nodejs/-/dockerfile-language-server-nodejs-0.13.0.tgz"
	dockerfilelanguageserver_cmd_tar := exec.Command("curl", "-L", dockerfilelanguageserver_tar_url, "-o", "package.tar.gz")
	err := dockerfilelanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockerfilelanguageserver_zip_url := "https://registry.npmjs.org/dockerfile-language-server-nodejs/-/dockerfile-language-server-nodejs-0.13.0.tgz"
	dockerfilelanguageserver_cmd_zip := exec.Command("curl", "-L", dockerfilelanguageserver_zip_url, "-o", "package.zip")
	err = dockerfilelanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockerfilelanguageserver_bin_url := "https://registry.npmjs.org/dockerfile-language-server-nodejs/-/dockerfile-language-server-nodejs-0.13.0.tgz"
	dockerfilelanguageserver_cmd_bin := exec.Command("curl", "-L", dockerfilelanguageserver_bin_url, "-o", "binary.bin")
	err = dockerfilelanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockerfilelanguageserver_src_url := "https://registry.npmjs.org/dockerfile-language-server-nodejs/-/dockerfile-language-server-nodejs-0.13.0.tgz"
	dockerfilelanguageserver_cmd_src := exec.Command("curl", "-L", dockerfilelanguageserver_src_url, "-o", "source.tar.gz")
	err = dockerfilelanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockerfilelanguageserver_cmd_direct := exec.Command("./binary")
	err = dockerfilelanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
