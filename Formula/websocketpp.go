package main

// Websocketpp - WebSocket++ is a cross platform header only C++ library
// Homepage: https://www.zaphoyd.com/websocketpp

import (
	"fmt"
	
	"os/exec"
)

func installWebsocketpp() {
	// Método 1: Descargar y extraer .tar.gz
	websocketpp_tar_url := "https://github.com/zaphoyd/websocketpp/archive/refs/tags/0.8.2.tar.gz"
	websocketpp_cmd_tar := exec.Command("curl", "-L", websocketpp_tar_url, "-o", "package.tar.gz")
	err := websocketpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	websocketpp_zip_url := "https://github.com/zaphoyd/websocketpp/archive/refs/tags/0.8.2.zip"
	websocketpp_cmd_zip := exec.Command("curl", "-L", websocketpp_zip_url, "-o", "package.zip")
	err = websocketpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	websocketpp_bin_url := "https://github.com/zaphoyd/websocketpp/archive/refs/tags/0.8.2.bin"
	websocketpp_cmd_bin := exec.Command("curl", "-L", websocketpp_bin_url, "-o", "binary.bin")
	err = websocketpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	websocketpp_src_url := "https://github.com/zaphoyd/websocketpp/archive/refs/tags/0.8.2.src.tar.gz"
	websocketpp_cmd_src := exec.Command("curl", "-L", websocketpp_src_url, "-o", "source.tar.gz")
	err = websocketpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	websocketpp_cmd_direct := exec.Command("./binary")
	err = websocketpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
