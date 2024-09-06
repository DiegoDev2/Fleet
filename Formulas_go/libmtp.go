package main

// Libmtp - Implementation of Microsoft's Media Transfer Protocol (MTP)
// Homepage: https://libmtp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibmtp() {
	// Método 1: Descargar y extraer .tar.gz
	libmtp_tar_url := "https://downloads.sourceforge.net/project/libmtp/libmtp/1.1.21/libmtp-1.1.21.tar.gz"
	libmtp_cmd_tar := exec.Command("curl", "-L", libmtp_tar_url, "-o", "package.tar.gz")
	err := libmtp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmtp_zip_url := "https://downloads.sourceforge.net/project/libmtp/libmtp/1.1.21/libmtp-1.1.21.zip"
	libmtp_cmd_zip := exec.Command("curl", "-L", libmtp_zip_url, "-o", "package.zip")
	err = libmtp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmtp_bin_url := "https://downloads.sourceforge.net/project/libmtp/libmtp/1.1.21/libmtp-1.1.21.bin"
	libmtp_cmd_bin := exec.Command("curl", "-L", libmtp_bin_url, "-o", "binary.bin")
	err = libmtp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmtp_src_url := "https://downloads.sourceforge.net/project/libmtp/libmtp/1.1.21/libmtp-1.1.21.src.tar.gz"
	libmtp_cmd_src := exec.Command("curl", "-L", libmtp_src_url, "-o", "source.tar.gz")
	err = libmtp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmtp_cmd_direct := exec.Command("./binary")
	err = libmtp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
