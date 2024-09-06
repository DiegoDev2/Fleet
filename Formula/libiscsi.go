package main

// Libiscsi - Client library and utilities for iscsi
// Homepage: https://github.com/sahlberg/libiscsi

import (
	"fmt"
	
	"os/exec"
)

func installLibiscsi() {
	// Método 1: Descargar y extraer .tar.gz
	libiscsi_tar_url := "https://github.com/sahlberg/libiscsi/archive/refs/tags/1.20.0.tar.gz"
	libiscsi_cmd_tar := exec.Command("curl", "-L", libiscsi_tar_url, "-o", "package.tar.gz")
	err := libiscsi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libiscsi_zip_url := "https://github.com/sahlberg/libiscsi/archive/refs/tags/1.20.0.zip"
	libiscsi_cmd_zip := exec.Command("curl", "-L", libiscsi_zip_url, "-o", "package.zip")
	err = libiscsi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libiscsi_bin_url := "https://github.com/sahlberg/libiscsi/archive/refs/tags/1.20.0.bin"
	libiscsi_cmd_bin := exec.Command("curl", "-L", libiscsi_bin_url, "-o", "binary.bin")
	err = libiscsi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libiscsi_src_url := "https://github.com/sahlberg/libiscsi/archive/refs/tags/1.20.0.src.tar.gz"
	libiscsi_cmd_src := exec.Command("curl", "-L", libiscsi_src_url, "-o", "source.tar.gz")
	err = libiscsi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libiscsi_cmd_direct := exec.Command("./binary")
	err = libiscsi_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cunit")
	exec.Command("latte", "install", "cunit").Run()
}
