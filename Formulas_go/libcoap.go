package main

// Libcoap - Lightweight application-protocol for resource-constrained devices
// Homepage: https://github.com/obgm/libcoap

import (
	"fmt"
	
	"os/exec"
)

func installLibcoap() {
	// Método 1: Descargar y extraer .tar.gz
	libcoap_tar_url := "https://github.com/obgm/libcoap/archive/refs/tags/v4.3.4a.tar.gz"
	libcoap_cmd_tar := exec.Command("curl", "-L", libcoap_tar_url, "-o", "package.tar.gz")
	err := libcoap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcoap_zip_url := "https://github.com/obgm/libcoap/archive/refs/tags/v4.3.4a.zip"
	libcoap_cmd_zip := exec.Command("curl", "-L", libcoap_zip_url, "-o", "package.zip")
	err = libcoap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcoap_bin_url := "https://github.com/obgm/libcoap/archive/refs/tags/v4.3.4a.bin"
	libcoap_cmd_bin := exec.Command("curl", "-L", libcoap_bin_url, "-o", "binary.bin")
	err = libcoap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcoap_src_url := "https://github.com/obgm/libcoap/archive/refs/tags/v4.3.4a.src.tar.gz"
	libcoap_cmd_src := exec.Command("curl", "-L", libcoap_src_url, "-o", "source.tar.gz")
	err = libcoap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcoap_cmd_direct := exec.Command("./binary")
	err = libcoap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
