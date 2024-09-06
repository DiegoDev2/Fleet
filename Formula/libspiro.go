package main

// Libspiro - Library to simplify the drawing of curves
// Homepage: https://github.com/fontforge/libspiro

import (
	"fmt"
	
	"os/exec"
)

func installLibspiro() {
	// Método 1: Descargar y extraer .tar.gz
	libspiro_tar_url := "https://github.com/fontforge/libspiro/releases/download/20240903/libspiro-dist-20240903.tar.gz"
	libspiro_cmd_tar := exec.Command("curl", "-L", libspiro_tar_url, "-o", "package.tar.gz")
	err := libspiro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspiro_zip_url := "https://github.com/fontforge/libspiro/releases/download/20240903/libspiro-dist-20240903.zip"
	libspiro_cmd_zip := exec.Command("curl", "-L", libspiro_zip_url, "-o", "package.zip")
	err = libspiro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspiro_bin_url := "https://github.com/fontforge/libspiro/releases/download/20240903/libspiro-dist-20240903.bin"
	libspiro_cmd_bin := exec.Command("curl", "-L", libspiro_bin_url, "-o", "binary.bin")
	err = libspiro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspiro_src_url := "https://github.com/fontforge/libspiro/releases/download/20240903/libspiro-dist-20240903.src.tar.gz"
	libspiro_cmd_src := exec.Command("curl", "-L", libspiro_src_url, "-o", "source.tar.gz")
	err = libspiro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspiro_cmd_direct := exec.Command("./binary")
	err = libspiro_cmd_direct.Run()
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
