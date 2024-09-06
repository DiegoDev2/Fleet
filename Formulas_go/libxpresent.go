package main

// Libxpresent - Xlib-based library for the X Present Extension
// Homepage: https://gitlab.freedesktop.org/xorg/lib/libxpresent

import (
	"fmt"
	
	"os/exec"
)

func installLibxpresent() {
	// Método 1: Descargar y extraer .tar.gz
	libxpresent_tar_url := "https://www.x.org/archive/individual/lib/libXpresent-1.0.1.tar.xz"
	libxpresent_cmd_tar := exec.Command("curl", "-L", libxpresent_tar_url, "-o", "package.tar.gz")
	err := libxpresent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxpresent_zip_url := "https://www.x.org/archive/individual/lib/libXpresent-1.0.1.tar.xz"
	libxpresent_cmd_zip := exec.Command("curl", "-L", libxpresent_zip_url, "-o", "package.zip")
	err = libxpresent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxpresent_bin_url := "https://www.x.org/archive/individual/lib/libXpresent-1.0.1.tar.xz"
	libxpresent_cmd_bin := exec.Command("curl", "-L", libxpresent_bin_url, "-o", "binary.bin")
	err = libxpresent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxpresent_src_url := "https://www.x.org/archive/individual/lib/libXpresent-1.0.1.tar.xz"
	libxpresent_cmd_src := exec.Command("curl", "-L", libxpresent_src_url, "-o", "source.tar.gz")
	err = libxpresent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxpresent_cmd_direct := exec.Command("./binary")
	err = libxpresent_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
}
