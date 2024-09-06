package main

// Libjwt - JSON Web Token C library
// Homepage: https://github.com/benmcollins/libjwt

import (
	"fmt"
	
	"os/exec"
)

func installLibjwt() {
	// Método 1: Descargar y extraer .tar.gz
	libjwt_tar_url := "https://github.com/benmcollins/libjwt/releases/download/v1.17.2/libjwt-1.17.2.tar.bz2"
	libjwt_cmd_tar := exec.Command("curl", "-L", libjwt_tar_url, "-o", "package.tar.gz")
	err := libjwt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libjwt_zip_url := "https://github.com/benmcollins/libjwt/releases/download/v1.17.2/libjwt-1.17.2.tar.bz2"
	libjwt_cmd_zip := exec.Command("curl", "-L", libjwt_zip_url, "-o", "package.zip")
	err = libjwt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libjwt_bin_url := "https://github.com/benmcollins/libjwt/releases/download/v1.17.2/libjwt-1.17.2.tar.bz2"
	libjwt_cmd_bin := exec.Command("curl", "-L", libjwt_bin_url, "-o", "binary.bin")
	err = libjwt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libjwt_src_url := "https://github.com/benmcollins/libjwt/releases/download/v1.17.2/libjwt-1.17.2.tar.bz2"
	libjwt_cmd_src := exec.Command("curl", "-L", libjwt_src_url, "-o", "source.tar.gz")
	err = libjwt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libjwt_cmd_direct := exec.Command("./binary")
	err = libjwt_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
