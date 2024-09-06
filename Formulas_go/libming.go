package main

// Libming - C library for generating Macromedia Flash files
// Homepage: https://github.com/libming/libming

import (
	"fmt"
	
	"os/exec"
)

func installLibming() {
	// Método 1: Descargar y extraer .tar.gz
	libming_tar_url := "https://github.com/libming/libming/archive/refs/tags/ming-0_4_8.tar.gz"
	libming_cmd_tar := exec.Command("curl", "-L", libming_tar_url, "-o", "package.tar.gz")
	err := libming_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libming_zip_url := "https://github.com/libming/libming/archive/refs/tags/ming-0_4_8.zip"
	libming_cmd_zip := exec.Command("curl", "-L", libming_zip_url, "-o", "package.zip")
	err = libming_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libming_bin_url := "https://github.com/libming/libming/archive/refs/tags/ming-0_4_8.bin"
	libming_cmd_bin := exec.Command("curl", "-L", libming_bin_url, "-o", "binary.bin")
	err = libming_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libming_src_url := "https://github.com/libming/libming/archive/refs/tags/ming-0_4_8.src.tar.gz"
	libming_cmd_src := exec.Command("curl", "-L", libming_src_url, "-o", "source.tar.gz")
	err = libming_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libming_cmd_direct := exec.Command("./binary")
	err = libming_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
}
