package main

// Libfido2 - Provides library functionality for FIDO U2F & FIDO 2.0, including USB
// Homepage: https://developers.yubico.com/libfido2/

import (
	"fmt"
	
	"os/exec"
)

func installLibfido2() {
	// Método 1: Descargar y extraer .tar.gz
	libfido2_tar_url := "https://github.com/Yubico/libfido2/archive/refs/tags/1.15.0.tar.gz"
	libfido2_cmd_tar := exec.Command("curl", "-L", libfido2_tar_url, "-o", "package.tar.gz")
	err := libfido2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfido2_zip_url := "https://github.com/Yubico/libfido2/archive/refs/tags/1.15.0.zip"
	libfido2_cmd_zip := exec.Command("curl", "-L", libfido2_zip_url, "-o", "package.zip")
	err = libfido2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfido2_bin_url := "https://github.com/Yubico/libfido2/archive/refs/tags/1.15.0.bin"
	libfido2_cmd_bin := exec.Command("curl", "-L", libfido2_bin_url, "-o", "binary.bin")
	err = libfido2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfido2_src_url := "https://github.com/Yubico/libfido2/archive/refs/tags/1.15.0.src.tar.gz"
	libfido2_cmd_src := exec.Command("curl", "-L", libfido2_src_url, "-o", "source.tar.gz")
	err = libfido2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfido2_cmd_direct := exec.Command("./binary")
	err = libfido2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: mandoc")
exec.Command("latte", "install", "mandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libcbor")
exec.Command("latte", "install", "libcbor")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
