package main

// Libusrsctp - Portable SCTP userland stack
// Homepage: https://github.com/sctplab/usrsctp

import (
	"fmt"
	
	"os/exec"
)

func installLibusrsctp() {
	// Método 1: Descargar y extraer .tar.gz
	libusrsctp_tar_url := "https://github.com/sctplab/usrsctp/archive/refs/tags/0.9.5.0.tar.gz"
	libusrsctp_cmd_tar := exec.Command("curl", "-L", libusrsctp_tar_url, "-o", "package.tar.gz")
	err := libusrsctp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libusrsctp_zip_url := "https://github.com/sctplab/usrsctp/archive/refs/tags/0.9.5.0.zip"
	libusrsctp_cmd_zip := exec.Command("curl", "-L", libusrsctp_zip_url, "-o", "package.zip")
	err = libusrsctp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libusrsctp_bin_url := "https://github.com/sctplab/usrsctp/archive/refs/tags/0.9.5.0.bin"
	libusrsctp_cmd_bin := exec.Command("curl", "-L", libusrsctp_bin_url, "-o", "binary.bin")
	err = libusrsctp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libusrsctp_src_url := "https://github.com/sctplab/usrsctp/archive/refs/tags/0.9.5.0.src.tar.gz"
	libusrsctp_cmd_src := exec.Command("curl", "-L", libusrsctp_src_url, "-o", "source.tar.gz")
	err = libusrsctp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libusrsctp_cmd_direct := exec.Command("./binary")
	err = libusrsctp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
