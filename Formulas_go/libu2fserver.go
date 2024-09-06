package main

// Libu2fServer - Server-side of the Universal 2nd Factor (U2F) protocol
// Homepage: https://developers.yubico.com/libu2f-server/

import (
	"fmt"
	
	"os/exec"
)

func installLibu2fServer() {
	// Método 1: Descargar y extraer .tar.gz
	libu2fserver_tar_url := "https://developers.yubico.com/libu2f-server/Releases/libu2f-server-1.1.0.tar.xz"
	libu2fserver_cmd_tar := exec.Command("curl", "-L", libu2fserver_tar_url, "-o", "package.tar.gz")
	err := libu2fserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libu2fserver_zip_url := "https://developers.yubico.com/libu2f-server/Releases/libu2f-server-1.1.0.tar.xz"
	libu2fserver_cmd_zip := exec.Command("curl", "-L", libu2fserver_zip_url, "-o", "package.zip")
	err = libu2fserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libu2fserver_bin_url := "https://developers.yubico.com/libu2f-server/Releases/libu2f-server-1.1.0.tar.xz"
	libu2fserver_cmd_bin := exec.Command("curl", "-L", libu2fserver_bin_url, "-o", "binary.bin")
	err = libu2fserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libu2fserver_src_url := "https://developers.yubico.com/libu2f-server/Releases/libu2f-server-1.1.0.tar.xz"
	libu2fserver_cmd_src := exec.Command("curl", "-L", libu2fserver_src_url, "-o", "source.tar.gz")
	err = libu2fserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libu2fserver_cmd_direct := exec.Command("./binary")
	err = libu2fserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: check")
exec.Command("latte", "install", "check")
	fmt.Println("Instalando dependencia: gengetopt")
exec.Command("latte", "install", "gengetopt")
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
