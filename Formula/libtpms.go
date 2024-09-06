package main

// Libtpms - Library for software emulation of a Trusted Platform Module
// Homepage: https://github.com/stefanberger/libtpms

import (
	"fmt"
	
	"os/exec"
)

func installLibtpms() {
	// Método 1: Descargar y extraer .tar.gz
	libtpms_tar_url := "https://github.com/stefanberger/libtpms/archive/refs/tags/v0.9.6.tar.gz"
	libtpms_cmd_tar := exec.Command("curl", "-L", libtpms_tar_url, "-o", "package.tar.gz")
	err := libtpms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtpms_zip_url := "https://github.com/stefanberger/libtpms/archive/refs/tags/v0.9.6.zip"
	libtpms_cmd_zip := exec.Command("curl", "-L", libtpms_zip_url, "-o", "package.zip")
	err = libtpms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtpms_bin_url := "https://github.com/stefanberger/libtpms/archive/refs/tags/v0.9.6.bin"
	libtpms_cmd_bin := exec.Command("curl", "-L", libtpms_bin_url, "-o", "binary.bin")
	err = libtpms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtpms_src_url := "https://github.com/stefanberger/libtpms/archive/refs/tags/v0.9.6.src.tar.gz"
	libtpms_cmd_src := exec.Command("curl", "-L", libtpms_src_url, "-o", "source.tar.gz")
	err = libtpms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtpms_cmd_direct := exec.Command("./binary")
	err = libtpms_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
