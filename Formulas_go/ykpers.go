package main

// Ykpers - YubiKey personalization library and tool
// Homepage: https://developers.yubico.com/yubikey-personalization/

import (
	"fmt"
	
	"os/exec"
)

func installYkpers() {
	// Método 1: Descargar y extraer .tar.gz
	ykpers_tar_url := "https://developers.yubico.com/yubikey-personalization/Releases/ykpers-1.20.0.tar.gz"
	ykpers_cmd_tar := exec.Command("curl", "-L", ykpers_tar_url, "-o", "package.tar.gz")
	err := ykpers_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ykpers_zip_url := "https://developers.yubico.com/yubikey-personalization/Releases/ykpers-1.20.0.zip"
	ykpers_cmd_zip := exec.Command("curl", "-L", ykpers_zip_url, "-o", "package.zip")
	err = ykpers_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ykpers_bin_url := "https://developers.yubico.com/yubikey-personalization/Releases/ykpers-1.20.0.bin"
	ykpers_cmd_bin := exec.Command("curl", "-L", ykpers_bin_url, "-o", "binary.bin")
	err = ykpers_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ykpers_src_url := "https://developers.yubico.com/yubikey-personalization/Releases/ykpers-1.20.0.src.tar.gz"
	ykpers_cmd_src := exec.Command("curl", "-L", ykpers_src_url, "-o", "source.tar.gz")
	err = ykpers_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ykpers_cmd_direct := exec.Command("./binary")
	err = ykpers_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: json-c")
exec.Command("latte", "install", "json-c")
	fmt.Println("Instalando dependencia: libyubikey")
exec.Command("latte", "install", "libyubikey")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
