package main

// Mockserver - Mock HTTP server and proxy
// Homepage: https://www.mock-server.com/

import (
	"fmt"
	
	"os/exec"
)

func installMockserver() {
	// Método 1: Descargar y extraer .tar.gz
	mockserver_tar_url := "https://search.maven.org/remotecontent?filepath=org/mock-server/mockserver-netty/5.15.0/mockserver-netty-5.15.0-brew-tar.tar"
	mockserver_cmd_tar := exec.Command("curl", "-L", mockserver_tar_url, "-o", "package.tar.gz")
	err := mockserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mockserver_zip_url := "https://search.maven.org/remotecontent?filepath=org/mock-server/mockserver-netty/5.15.0/mockserver-netty-5.15.0-brew-tar.tar"
	mockserver_cmd_zip := exec.Command("curl", "-L", mockserver_zip_url, "-o", "package.zip")
	err = mockserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mockserver_bin_url := "https://search.maven.org/remotecontent?filepath=org/mock-server/mockserver-netty/5.15.0/mockserver-netty-5.15.0-brew-tar.tar"
	mockserver_cmd_bin := exec.Command("curl", "-L", mockserver_bin_url, "-o", "binary.bin")
	err = mockserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mockserver_src_url := "https://search.maven.org/remotecontent?filepath=org/mock-server/mockserver-netty/5.15.0/mockserver-netty-5.15.0-brew-tar.tar"
	mockserver_cmd_src := exec.Command("curl", "-L", mockserver_src_url, "-o", "source.tar.gz")
	err = mockserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mockserver_cmd_direct := exec.Command("./binary")
	err = mockserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
