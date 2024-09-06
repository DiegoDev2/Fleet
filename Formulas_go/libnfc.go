package main

// Libnfc - Low level NFC SDK and Programmers API
// Homepage: https://github.com/nfc-tools/libnfc

import (
	"fmt"
	
	"os/exec"
)

func installLibnfc() {
	// Método 1: Descargar y extraer .tar.gz
	libnfc_tar_url := "https://github.com/nfc-tools/libnfc/releases/download/libnfc-1.8.0/libnfc-1.8.0.tar.bz2"
	libnfc_cmd_tar := exec.Command("curl", "-L", libnfc_tar_url, "-o", "package.tar.gz")
	err := libnfc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnfc_zip_url := "https://github.com/nfc-tools/libnfc/releases/download/libnfc-1.8.0/libnfc-1.8.0.tar.bz2"
	libnfc_cmd_zip := exec.Command("curl", "-L", libnfc_zip_url, "-o", "package.zip")
	err = libnfc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnfc_bin_url := "https://github.com/nfc-tools/libnfc/releases/download/libnfc-1.8.0/libnfc-1.8.0.tar.bz2"
	libnfc_cmd_bin := exec.Command("curl", "-L", libnfc_bin_url, "-o", "binary.bin")
	err = libnfc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnfc_src_url := "https://github.com/nfc-tools/libnfc/releases/download/libnfc-1.8.0/libnfc-1.8.0.tar.bz2"
	libnfc_cmd_src := exec.Command("curl", "-L", libnfc_src_url, "-o", "source.tar.gz")
	err = libnfc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnfc_cmd_direct := exec.Command("./binary")
	err = libnfc_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
