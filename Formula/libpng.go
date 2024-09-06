package main

// Libpng - Library for manipulating PNG images
// Homepage: http://www.libpng.org/pub/png/libpng.html

import (
	"fmt"
	
	"os/exec"
)

func installLibpng() {
	// Método 1: Descargar y extraer .tar.gz
	libpng_tar_url := "https://downloads.sourceforge.net/project/libpng/libpng16/1.6.43/libpng-1.6.43.tar.xz"
	libpng_cmd_tar := exec.Command("curl", "-L", libpng_tar_url, "-o", "package.tar.gz")
	err := libpng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpng_zip_url := "https://downloads.sourceforge.net/project/libpng/libpng16/1.6.43/libpng-1.6.43.tar.xz"
	libpng_cmd_zip := exec.Command("curl", "-L", libpng_zip_url, "-o", "package.zip")
	err = libpng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpng_bin_url := "https://downloads.sourceforge.net/project/libpng/libpng16/1.6.43/libpng-1.6.43.tar.xz"
	libpng_cmd_bin := exec.Command("curl", "-L", libpng_bin_url, "-o", "binary.bin")
	err = libpng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpng_src_url := "https://downloads.sourceforge.net/project/libpng/libpng16/1.6.43/libpng-1.6.43.tar.xz"
	libpng_cmd_src := exec.Command("curl", "-L", libpng_src_url, "-o", "source.tar.gz")
	err = libpng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpng_cmd_direct := exec.Command("./binary")
	err = libpng_cmd_direct.Run()
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
}
