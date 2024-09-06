package main

// Ttyd - Command-line tool for sharing terminal over the web
// Homepage: https://tsl0922.github.io/ttyd/

import (
	"fmt"
	
	"os/exec"
)

func installTtyd() {
	// Método 1: Descargar y extraer .tar.gz
	ttyd_tar_url := "https://github.com/tsl0922/ttyd/archive/refs/tags/1.7.7.tar.gz"
	ttyd_cmd_tar := exec.Command("curl", "-L", ttyd_tar_url, "-o", "package.tar.gz")
	err := ttyd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttyd_zip_url := "https://github.com/tsl0922/ttyd/archive/refs/tags/1.7.7.zip"
	ttyd_cmd_zip := exec.Command("curl", "-L", ttyd_zip_url, "-o", "package.zip")
	err = ttyd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttyd_bin_url := "https://github.com/tsl0922/ttyd/archive/refs/tags/1.7.7.bin"
	ttyd_cmd_bin := exec.Command("curl", "-L", ttyd_bin_url, "-o", "binary.bin")
	err = ttyd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttyd_src_url := "https://github.com/tsl0922/ttyd/archive/refs/tags/1.7.7.src.tar.gz"
	ttyd_cmd_src := exec.Command("curl", "-L", ttyd_src_url, "-o", "source.tar.gz")
	err = ttyd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttyd_cmd_direct := exec.Command("./binary")
	err = ttyd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: libwebsockets")
exec.Command("latte", "install", "libwebsockets")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
