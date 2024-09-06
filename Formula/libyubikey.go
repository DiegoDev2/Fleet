package main

// Libyubikey - C library for manipulating Yubico one-time passwords
// Homepage: https://yubico.github.io/yubico-c/

import (
	"fmt"
	
	"os/exec"
)

func installLibyubikey() {
	// Método 1: Descargar y extraer .tar.gz
	libyubikey_tar_url := "https://developers.yubico.com/yubico-c/Releases/libyubikey-1.13.tar.gz"
	libyubikey_cmd_tar := exec.Command("curl", "-L", libyubikey_tar_url, "-o", "package.tar.gz")
	err := libyubikey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libyubikey_zip_url := "https://developers.yubico.com/yubico-c/Releases/libyubikey-1.13.zip"
	libyubikey_cmd_zip := exec.Command("curl", "-L", libyubikey_zip_url, "-o", "package.zip")
	err = libyubikey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libyubikey_bin_url := "https://developers.yubico.com/yubico-c/Releases/libyubikey-1.13.bin"
	libyubikey_cmd_bin := exec.Command("curl", "-L", libyubikey_bin_url, "-o", "binary.bin")
	err = libyubikey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libyubikey_src_url := "https://developers.yubico.com/yubico-c/Releases/libyubikey-1.13.src.tar.gz"
	libyubikey_cmd_src := exec.Command("curl", "-L", libyubikey_src_url, "-o", "source.tar.gz")
	err = libyubikey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libyubikey_cmd_direct := exec.Command("./binary")
	err = libyubikey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
