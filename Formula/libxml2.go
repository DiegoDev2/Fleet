package main

// Libxml2 - GNOME XML library
// Homepage: http://xmlsoft.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxml2() {
	// Método 1: Descargar y extraer .tar.gz
	libxml2_tar_url := "https://download.gnome.org/sources/libxml2/2.12/libxml2-2.12.8.tar.xz"
	libxml2_cmd_tar := exec.Command("curl", "-L", libxml2_tar_url, "-o", "package.tar.gz")
	err := libxml2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxml2_zip_url := "https://download.gnome.org/sources/libxml2/2.12/libxml2-2.12.8.tar.xz"
	libxml2_cmd_zip := exec.Command("curl", "-L", libxml2_zip_url, "-o", "package.zip")
	err = libxml2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxml2_bin_url := "https://download.gnome.org/sources/libxml2/2.12/libxml2-2.12.8.tar.xz"
	libxml2_cmd_bin := exec.Command("curl", "-L", libxml2_bin_url, "-o", "binary.bin")
	err = libxml2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxml2_src_url := "https://download.gnome.org/sources/libxml2/2.12/libxml2-2.12.8.tar.xz"
	libxml2_cmd_src := exec.Command("curl", "-L", libxml2_src_url, "-o", "source.tar.gz")
	err = libxml2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxml2_cmd_direct := exec.Command("./binary")
	err = libxml2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
