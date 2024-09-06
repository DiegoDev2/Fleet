package main

// Libxcb - X.Org: Interface to the X Window System protocol
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxcb() {
	// Método 1: Descargar y extraer .tar.gz
	libxcb_tar_url := "https://xorg.freedesktop.org/archive/individual/lib/libxcb-1.17.0.tar.xz"
	libxcb_cmd_tar := exec.Command("curl", "-L", libxcb_tar_url, "-o", "package.tar.gz")
	err := libxcb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxcb_zip_url := "https://xorg.freedesktop.org/archive/individual/lib/libxcb-1.17.0.tar.xz"
	libxcb_cmd_zip := exec.Command("curl", "-L", libxcb_zip_url, "-o", "package.zip")
	err = libxcb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxcb_bin_url := "https://xorg.freedesktop.org/archive/individual/lib/libxcb-1.17.0.tar.xz"
	libxcb_cmd_bin := exec.Command("curl", "-L", libxcb_bin_url, "-o", "binary.bin")
	err = libxcb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxcb_src_url := "https://xorg.freedesktop.org/archive/individual/lib/libxcb-1.17.0.tar.xz"
	libxcb_cmd_src := exec.Command("curl", "-L", libxcb_src_url, "-o", "source.tar.gz")
	err = libxcb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxcb_cmd_direct := exec.Command("./binary")
	err = libxcb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: xcb-proto")
exec.Command("latte", "install", "xcb-proto")
	fmt.Println("Instalando dependencia: libxau")
exec.Command("latte", "install", "libxau")
	fmt.Println("Instalando dependencia: libxdmcp")
exec.Command("latte", "install", "libxdmcp")
}
