package main

// LaunchSocketServer - Bind to privileged ports without running a server as root
// Homepage: https://github.com/mistydemeo/launch_socket_server

import (
	"fmt"
	
	"os/exec"
)

func installLaunchSocketServer() {
	// Método 1: Descargar y extraer .tar.gz
	launchsocketserver_tar_url := "https://github.com/mistydemeo/launch_socket_server/archive/refs/tags/v2.0.0.tar.gz"
	launchsocketserver_cmd_tar := exec.Command("curl", "-L", launchsocketserver_tar_url, "-o", "package.tar.gz")
	err := launchsocketserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	launchsocketserver_zip_url := "https://github.com/mistydemeo/launch_socket_server/archive/refs/tags/v2.0.0.zip"
	launchsocketserver_cmd_zip := exec.Command("curl", "-L", launchsocketserver_zip_url, "-o", "package.zip")
	err = launchsocketserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	launchsocketserver_bin_url := "https://github.com/mistydemeo/launch_socket_server/archive/refs/tags/v2.0.0.bin"
	launchsocketserver_cmd_bin := exec.Command("curl", "-L", launchsocketserver_bin_url, "-o", "binary.bin")
	err = launchsocketserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	launchsocketserver_src_url := "https://github.com/mistydemeo/launch_socket_server/archive/refs/tags/v2.0.0.src.tar.gz"
	launchsocketserver_cmd_src := exec.Command("curl", "-L", launchsocketserver_src_url, "-o", "source.tar.gz")
	err = launchsocketserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	launchsocketserver_cmd_direct := exec.Command("./binary")
	err = launchsocketserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
