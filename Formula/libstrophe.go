package main

// Libstrophe - XMPP library for C
// Homepage: https://strophe.im/libstrophe/

import (
	"fmt"
	
	"os/exec"
)

func installLibstrophe() {
	// Método 1: Descargar y extraer .tar.gz
	libstrophe_tar_url := "https://github.com/strophe/libstrophe/archive/refs/tags/0.13.1.tar.gz"
	libstrophe_cmd_tar := exec.Command("curl", "-L", libstrophe_tar_url, "-o", "package.tar.gz")
	err := libstrophe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libstrophe_zip_url := "https://github.com/strophe/libstrophe/archive/refs/tags/0.13.1.zip"
	libstrophe_cmd_zip := exec.Command("curl", "-L", libstrophe_zip_url, "-o", "package.zip")
	err = libstrophe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libstrophe_bin_url := "https://github.com/strophe/libstrophe/archive/refs/tags/0.13.1.bin"
	libstrophe_cmd_bin := exec.Command("curl", "-L", libstrophe_bin_url, "-o", "binary.bin")
	err = libstrophe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libstrophe_src_url := "https://github.com/strophe/libstrophe/archive/refs/tags/0.13.1.src.tar.gz"
	libstrophe_cmd_src := exec.Command("curl", "-L", libstrophe_src_url, "-o", "source.tar.gz")
	err = libstrophe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libstrophe_cmd_direct := exec.Command("./binary")
	err = libstrophe_cmd_direct.Run()
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
