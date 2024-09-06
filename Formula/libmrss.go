package main

// Libmrss - C library for RSS files or streams
// Homepage: https://github.com/bakulf/libmrss

import (
	"fmt"
	
	"os/exec"
)

func installLibmrss() {
	// Método 1: Descargar y extraer .tar.gz
	libmrss_tar_url := "https://github.com/bakulf/libmrss/archive/refs/tags/0.19.4.tar.gz"
	libmrss_cmd_tar := exec.Command("curl", "-L", libmrss_tar_url, "-o", "package.tar.gz")
	err := libmrss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmrss_zip_url := "https://github.com/bakulf/libmrss/archive/refs/tags/0.19.4.zip"
	libmrss_cmd_zip := exec.Command("curl", "-L", libmrss_zip_url, "-o", "package.zip")
	err = libmrss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmrss_bin_url := "https://github.com/bakulf/libmrss/archive/refs/tags/0.19.4.bin"
	libmrss_cmd_bin := exec.Command("curl", "-L", libmrss_bin_url, "-o", "binary.bin")
	err = libmrss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmrss_src_url := "https://github.com/bakulf/libmrss/archive/refs/tags/0.19.4.src.tar.gz"
	libmrss_cmd_src := exec.Command("curl", "-L", libmrss_src_url, "-o", "source.tar.gz")
	err = libmrss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmrss_cmd_direct := exec.Command("./binary")
	err = libmrss_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libnxml")
	exec.Command("latte", "install", "libnxml").Run()
}
