package main

// Libdap - Framework for scientific data networking
// Homepage: https://www.opendap.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibdap() {
	// Método 1: Descargar y extraer .tar.gz
	libdap_tar_url := "https://www.opendap.org/pub/source/libdap-3.21.0-27.tar.gz"
	libdap_cmd_tar := exec.Command("curl", "-L", libdap_tar_url, "-o", "package.tar.gz")
	err := libdap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdap_zip_url := "https://www.opendap.org/pub/source/libdap-3.21.0-27.zip"
	libdap_cmd_zip := exec.Command("curl", "-L", libdap_zip_url, "-o", "package.zip")
	err = libdap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdap_bin_url := "https://www.opendap.org/pub/source/libdap-3.21.0-27.bin"
	libdap_cmd_bin := exec.Command("curl", "-L", libdap_bin_url, "-o", "binary.bin")
	err = libdap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdap_src_url := "https://www.opendap.org/pub/source/libdap-3.21.0-27.src.tar.gz"
	libdap_cmd_src := exec.Command("curl", "-L", libdap_src_url, "-o", "source.tar.gz")
	err = libdap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdap_cmd_direct := exec.Command("./binary")
	err = libdap_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
