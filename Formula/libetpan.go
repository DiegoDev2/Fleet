package main

// Libetpan - Portable mail library handling several protocols
// Homepage: https://www.etpan.org/libetpan.html

import (
	"fmt"
	
	"os/exec"
)

func installLibetpan() {
	// Método 1: Descargar y extraer .tar.gz
	libetpan_tar_url := "https://github.com/dinhvh/libetpan/archive/refs/tags/1.9.4.tar.gz"
	libetpan_cmd_tar := exec.Command("curl", "-L", libetpan_tar_url, "-o", "package.tar.gz")
	err := libetpan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libetpan_zip_url := "https://github.com/dinhvh/libetpan/archive/refs/tags/1.9.4.zip"
	libetpan_cmd_zip := exec.Command("curl", "-L", libetpan_zip_url, "-o", "package.zip")
	err = libetpan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libetpan_bin_url := "https://github.com/dinhvh/libetpan/archive/refs/tags/1.9.4.bin"
	libetpan_cmd_bin := exec.Command("curl", "-L", libetpan_bin_url, "-o", "binary.bin")
	err = libetpan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libetpan_src_url := "https://github.com/dinhvh/libetpan/archive/refs/tags/1.9.4.src.tar.gz"
	libetpan_cmd_src := exec.Command("curl", "-L", libetpan_src_url, "-o", "source.tar.gz")
	err = libetpan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libetpan_cmd_direct := exec.Command("./binary")
	err = libetpan_cmd_direct.Run()
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
