package main

// Mfterm - Terminal for working with Mifare Classic 1-4k Tags
// Homepage: https://github.com/4ZM/mfterm

import (
	"fmt"
	
	"os/exec"
)

func installMfterm() {
	// Método 1: Descargar y extraer .tar.gz
	mfterm_tar_url := "https://github.com/4ZM/mfterm/releases/download/v1.0.7/mfterm-1.0.7.tar.gz"
	mfterm_cmd_tar := exec.Command("curl", "-L", mfterm_tar_url, "-o", "package.tar.gz")
	err := mfterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mfterm_zip_url := "https://github.com/4ZM/mfterm/releases/download/v1.0.7/mfterm-1.0.7.zip"
	mfterm_cmd_zip := exec.Command("curl", "-L", mfterm_zip_url, "-o", "package.zip")
	err = mfterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mfterm_bin_url := "https://github.com/4ZM/mfterm/releases/download/v1.0.7/mfterm-1.0.7.bin"
	mfterm_cmd_bin := exec.Command("curl", "-L", mfterm_bin_url, "-o", "binary.bin")
	err = mfterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mfterm_src_url := "https://github.com/4ZM/mfterm/releases/download/v1.0.7/mfterm-1.0.7.src.tar.gz"
	mfterm_cmd_src := exec.Command("curl", "-L", mfterm_src_url, "-o", "source.tar.gz")
	err = mfterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mfterm_cmd_direct := exec.Command("./binary")
	err = mfterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libnfc")
exec.Command("latte", "install", "libnfc")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
