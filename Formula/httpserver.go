package main

// HttpServer - Simple zero-configuration command-line HTTP server
// Homepage: https://github.com/http-party/http-server

import (
	"fmt"
	
	"os/exec"
)

func installHttpServer() {
	// Método 1: Descargar y extraer .tar.gz
	httpserver_tar_url := "https://registry.npmjs.org/http-server/-/http-server-14.1.1.tgz"
	httpserver_cmd_tar := exec.Command("curl", "-L", httpserver_tar_url, "-o", "package.tar.gz")
	err := httpserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpserver_zip_url := "https://registry.npmjs.org/http-server/-/http-server-14.1.1.tgz"
	httpserver_cmd_zip := exec.Command("curl", "-L", httpserver_zip_url, "-o", "package.zip")
	err = httpserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpserver_bin_url := "https://registry.npmjs.org/http-server/-/http-server-14.1.1.tgz"
	httpserver_cmd_bin := exec.Command("curl", "-L", httpserver_bin_url, "-o", "binary.bin")
	err = httpserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpserver_src_url := "https://registry.npmjs.org/http-server/-/http-server-14.1.1.tgz"
	httpserver_cmd_src := exec.Command("curl", "-L", httpserver_src_url, "-o", "source.tar.gz")
	err = httpserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpserver_cmd_direct := exec.Command("./binary")
	err = httpserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
