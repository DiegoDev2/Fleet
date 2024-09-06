package main

// Libfreefare - API for MIFARE card manipulations
// Homepage: https://github.com/nfc-tools/libfreefare

import (
	"fmt"
	
	"os/exec"
)

func installLibfreefare() {
	// Método 1: Descargar y extraer .tar.gz
	libfreefare_tar_url := "https://github.com/nfc-tools/libfreefare/releases/download/libfreefare-0.4.0/libfreefare-0.4.0.tar.bz2"
	libfreefare_cmd_tar := exec.Command("curl", "-L", libfreefare_tar_url, "-o", "package.tar.gz")
	err := libfreefare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfreefare_zip_url := "https://github.com/nfc-tools/libfreefare/releases/download/libfreefare-0.4.0/libfreefare-0.4.0.tar.bz2"
	libfreefare_cmd_zip := exec.Command("curl", "-L", libfreefare_zip_url, "-o", "package.zip")
	err = libfreefare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfreefare_bin_url := "https://github.com/nfc-tools/libfreefare/releases/download/libfreefare-0.4.0/libfreefare-0.4.0.tar.bz2"
	libfreefare_cmd_bin := exec.Command("curl", "-L", libfreefare_bin_url, "-o", "binary.bin")
	err = libfreefare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfreefare_src_url := "https://github.com/nfc-tools/libfreefare/releases/download/libfreefare-0.4.0/libfreefare-0.4.0.tar.bz2"
	libfreefare_cmd_src := exec.Command("curl", "-L", libfreefare_src_url, "-o", "source.tar.gz")
	err = libfreefare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfreefare_cmd_direct := exec.Command("./binary")
	err = libfreefare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libnfc")
exec.Command("latte", "install", "libnfc")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
