package main

// Ykclient - Library to validate YubiKey OTPs against YubiCloud
// Homepage: https://developers.yubico.com/yubico-c-client/

import (
	"fmt"
	
	"os/exec"
)

func installYkclient() {
	// Método 1: Descargar y extraer .tar.gz
	ykclient_tar_url := "https://developers.yubico.com/yubico-c-client/Releases/ykclient-2.15.tar.gz"
	ykclient_cmd_tar := exec.Command("curl", "-L", ykclient_tar_url, "-o", "package.tar.gz")
	err := ykclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ykclient_zip_url := "https://developers.yubico.com/yubico-c-client/Releases/ykclient-2.15.zip"
	ykclient_cmd_zip := exec.Command("curl", "-L", ykclient_zip_url, "-o", "package.zip")
	err = ykclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ykclient_bin_url := "https://developers.yubico.com/yubico-c-client/Releases/ykclient-2.15.bin"
	ykclient_cmd_bin := exec.Command("curl", "-L", ykclient_bin_url, "-o", "binary.bin")
	err = ykclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ykclient_src_url := "https://developers.yubico.com/yubico-c-client/Releases/ykclient-2.15.src.tar.gz"
	ykclient_cmd_src := exec.Command("curl", "-L", ykclient_src_url, "-o", "source.tar.gz")
	err = ykclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ykclient_cmd_direct := exec.Command("./binary")
	err = ykclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
