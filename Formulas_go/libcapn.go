package main

// Libcapn - C library to send push notifications to Apple devices
// Homepage: https://web.archive.org/web/20181220090839/libcapn.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibcapn() {
	// Método 1: Descargar y extraer .tar.gz
	libcapn_tar_url := "https://github.com/adobkin/libcapn/archive/refs/tags/v2.0.0.tar.gz"
	libcapn_cmd_tar := exec.Command("curl", "-L", libcapn_tar_url, "-o", "package.tar.gz")
	err := libcapn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcapn_zip_url := "https://github.com/adobkin/libcapn/archive/refs/tags/v2.0.0.zip"
	libcapn_cmd_zip := exec.Command("curl", "-L", libcapn_zip_url, "-o", "package.zip")
	err = libcapn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcapn_bin_url := "https://github.com/adobkin/libcapn/archive/refs/tags/v2.0.0.bin"
	libcapn_cmd_bin := exec.Command("curl", "-L", libcapn_bin_url, "-o", "binary.bin")
	err = libcapn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcapn_src_url := "https://github.com/adobkin/libcapn/archive/refs/tags/v2.0.0.src.tar.gz"
	libcapn_cmd_src := exec.Command("curl", "-L", libcapn_src_url, "-o", "source.tar.gz")
	err = libcapn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcapn_cmd_direct := exec.Command("./binary")
	err = libcapn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
