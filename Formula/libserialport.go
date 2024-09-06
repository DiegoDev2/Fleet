package main

// Libserialport - Cross-platform serial port C library
// Homepage: https://sigrok.org/wiki/Libserialport

import (
	"fmt"
	
	"os/exec"
)

func installLibserialport() {
	// Método 1: Descargar y extraer .tar.gz
	libserialport_tar_url := "https://sigrok.org/download/source/libserialport/libserialport-0.1.1.tar.gz"
	libserialport_cmd_tar := exec.Command("curl", "-L", libserialport_tar_url, "-o", "package.tar.gz")
	err := libserialport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libserialport_zip_url := "https://sigrok.org/download/source/libserialport/libserialport-0.1.1.zip"
	libserialport_cmd_zip := exec.Command("curl", "-L", libserialport_zip_url, "-o", "package.zip")
	err = libserialport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libserialport_bin_url := "https://sigrok.org/download/source/libserialport/libserialport-0.1.1.bin"
	libserialport_cmd_bin := exec.Command("curl", "-L", libserialport_bin_url, "-o", "binary.bin")
	err = libserialport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libserialport_src_url := "https://sigrok.org/download/source/libserialport/libserialport-0.1.1.src.tar.gz"
	libserialport_cmd_src := exec.Command("curl", "-L", libserialport_src_url, "-o", "source.tar.gz")
	err = libserialport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libserialport_cmd_direct := exec.Command("./binary")
	err = libserialport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
