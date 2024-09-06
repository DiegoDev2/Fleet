package main

// Libxslt - C XSLT library for GNOME
// Homepage: http://xmlsoft.org/XSLT/

import (
	"fmt"
	
	"os/exec"
)

func installLibxslt() {
	// Método 1: Descargar y extraer .tar.gz
	libxslt_tar_url := "https://download.gnome.org/sources/libxslt/1.1/libxslt-1.1.42.tar.xz"
	libxslt_cmd_tar := exec.Command("curl", "-L", libxslt_tar_url, "-o", "package.tar.gz")
	err := libxslt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxslt_zip_url := "https://download.gnome.org/sources/libxslt/1.1/libxslt-1.1.42.tar.xz"
	libxslt_cmd_zip := exec.Command("curl", "-L", libxslt_zip_url, "-o", "package.zip")
	err = libxslt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxslt_bin_url := "https://download.gnome.org/sources/libxslt/1.1/libxslt-1.1.42.tar.xz"
	libxslt_cmd_bin := exec.Command("curl", "-L", libxslt_bin_url, "-o", "binary.bin")
	err = libxslt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxslt_src_url := "https://download.gnome.org/sources/libxslt/1.1/libxslt-1.1.42.tar.xz"
	libxslt_cmd_src := exec.Command("curl", "-L", libxslt_src_url, "-o", "source.tar.gz")
	err = libxslt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxslt_cmd_direct := exec.Command("./binary")
	err = libxslt_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
