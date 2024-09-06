package main

// Libwpd - General purpose library for reading WordPerfect files
// Homepage: https://libwpd.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibwpd() {
	// Método 1: Descargar y extraer .tar.gz
	libwpd_tar_url := "https://dev-www.libreoffice.org/src/libwpd-0.10.3.tar.xz"
	libwpd_cmd_tar := exec.Command("curl", "-L", libwpd_tar_url, "-o", "package.tar.gz")
	err := libwpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwpd_zip_url := "https://dev-www.libreoffice.org/src/libwpd-0.10.3.tar.xz"
	libwpd_cmd_zip := exec.Command("curl", "-L", libwpd_zip_url, "-o", "package.zip")
	err = libwpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwpd_bin_url := "https://dev-www.libreoffice.org/src/libwpd-0.10.3.tar.xz"
	libwpd_cmd_bin := exec.Command("curl", "-L", libwpd_bin_url, "-o", "binary.bin")
	err = libwpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwpd_src_url := "https://dev-www.libreoffice.org/src/libwpd-0.10.3.tar.xz"
	libwpd_cmd_src := exec.Command("curl", "-L", libwpd_src_url, "-o", "source.tar.gz")
	err = libwpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwpd_cmd_direct := exec.Command("./binary")
	err = libwpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libgsf")
exec.Command("latte", "install", "libgsf")
	fmt.Println("Instalando dependencia: librevenge")
exec.Command("latte", "install", "librevenge")
}
