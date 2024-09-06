package main

// Libwpg - Library for reading and parsing Word Perfect Graphics format
// Homepage: https://libwpg.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibwpg() {
	// Método 1: Descargar y extraer .tar.gz
	libwpg_tar_url := "https://dev-www.libreoffice.org/src/libwpg-0.3.4.tar.xz"
	libwpg_cmd_tar := exec.Command("curl", "-L", libwpg_tar_url, "-o", "package.tar.gz")
	err := libwpg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwpg_zip_url := "https://dev-www.libreoffice.org/src/libwpg-0.3.4.tar.xz"
	libwpg_cmd_zip := exec.Command("curl", "-L", libwpg_zip_url, "-o", "package.zip")
	err = libwpg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwpg_bin_url := "https://dev-www.libreoffice.org/src/libwpg-0.3.4.tar.xz"
	libwpg_cmd_bin := exec.Command("curl", "-L", libwpg_bin_url, "-o", "binary.bin")
	err = libwpg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwpg_src_url := "https://dev-www.libreoffice.org/src/libwpg-0.3.4.tar.xz"
	libwpg_cmd_src := exec.Command("curl", "-L", libwpg_src_url, "-o", "source.tar.gz")
	err = libwpg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwpg_cmd_direct := exec.Command("./binary")
	err = libwpg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: librevenge")
exec.Command("latte", "install", "librevenge")
	fmt.Println("Instalando dependencia: libwpd")
exec.Command("latte", "install", "libwpd")
}
