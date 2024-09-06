package main

// Libftdi - Library to talk to FTDI chips
// Homepage: https://www.intra2net.com/en/developer/libftdi

import (
	"fmt"
	
	"os/exec"
)

func installLibftdi() {
	// Método 1: Descargar y extraer .tar.gz
	libftdi_tar_url := "https://www.intra2net.com/en/developer/libftdi/download/libftdi1-1.5.tar.bz2"
	libftdi_cmd_tar := exec.Command("curl", "-L", libftdi_tar_url, "-o", "package.tar.gz")
	err := libftdi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libftdi_zip_url := "https://www.intra2net.com/en/developer/libftdi/download/libftdi1-1.5.tar.bz2"
	libftdi_cmd_zip := exec.Command("curl", "-L", libftdi_zip_url, "-o", "package.zip")
	err = libftdi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libftdi_bin_url := "https://www.intra2net.com/en/developer/libftdi/download/libftdi1-1.5.tar.bz2"
	libftdi_cmd_bin := exec.Command("curl", "-L", libftdi_bin_url, "-o", "binary.bin")
	err = libftdi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libftdi_src_url := "https://www.intra2net.com/en/developer/libftdi/download/libftdi1-1.5.tar.bz2"
	libftdi_cmd_src := exec.Command("curl", "-L", libftdi_src_url, "-o", "source.tar.gz")
	err = libftdi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libftdi_cmd_direct := exec.Command("./binary")
	err = libftdi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: confuse")
exec.Command("latte", "install", "confuse")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
