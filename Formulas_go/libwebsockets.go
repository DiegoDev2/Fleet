package main

// Libwebsockets - C websockets server library
// Homepage: https://github.com/warmcat/libwebsockets

import (
	"fmt"
	
	"os/exec"
)

func installLibwebsockets() {
	// Método 1: Descargar y extraer .tar.gz
	libwebsockets_tar_url := "https://github.com/warmcat/libwebsockets/archive/refs/tags/v4.3.3.tar.gz"
	libwebsockets_cmd_tar := exec.Command("curl", "-L", libwebsockets_tar_url, "-o", "package.tar.gz")
	err := libwebsockets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwebsockets_zip_url := "https://github.com/warmcat/libwebsockets/archive/refs/tags/v4.3.3.zip"
	libwebsockets_cmd_zip := exec.Command("curl", "-L", libwebsockets_zip_url, "-o", "package.zip")
	err = libwebsockets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwebsockets_bin_url := "https://github.com/warmcat/libwebsockets/archive/refs/tags/v4.3.3.bin"
	libwebsockets_cmd_bin := exec.Command("curl", "-L", libwebsockets_bin_url, "-o", "binary.bin")
	err = libwebsockets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwebsockets_src_url := "https://github.com/warmcat/libwebsockets/archive/refs/tags/v4.3.3.src.tar.gz"
	libwebsockets_cmd_src := exec.Command("curl", "-L", libwebsockets_src_url, "-o", "source.tar.gz")
	err = libwebsockets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwebsockets_cmd_direct := exec.Command("./binary")
	err = libwebsockets_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
