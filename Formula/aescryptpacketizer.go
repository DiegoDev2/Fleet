package main

// AescryptPacketizer - Encrypt and decrypt using 256-bit AES encryption
// Homepage: https://www.aescrypt.com

import (
	"fmt"
	
	"os/exec"
)

func installAescryptPacketizer() {
	// Método 1: Descargar y extraer .tar.gz
	aescryptpacketizer_tar_url := "https://www.aescrypt.com/download/v3/linux/aescrypt-3.16.tgz"
	aescryptpacketizer_cmd_tar := exec.Command("curl", "-L", aescryptpacketizer_tar_url, "-o", "package.tar.gz")
	err := aescryptpacketizer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aescryptpacketizer_zip_url := "https://www.aescrypt.com/download/v3/linux/aescrypt-3.16.tgz"
	aescryptpacketizer_cmd_zip := exec.Command("curl", "-L", aescryptpacketizer_zip_url, "-o", "package.zip")
	err = aescryptpacketizer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aescryptpacketizer_bin_url := "https://www.aescrypt.com/download/v3/linux/aescrypt-3.16.tgz"
	aescryptpacketizer_cmd_bin := exec.Command("curl", "-L", aescryptpacketizer_bin_url, "-o", "binary.bin")
	err = aescryptpacketizer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aescryptpacketizer_src_url := "https://www.aescrypt.com/download/v3/linux/aescrypt-3.16.tgz"
	aescryptpacketizer_cmd_src := exec.Command("curl", "-L", aescryptpacketizer_src_url, "-o", "source.tar.gz")
	err = aescryptpacketizer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aescryptpacketizer_cmd_direct := exec.Command("./binary")
	err = aescryptpacketizer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
