package main

// Libwpe - General-purpose library for WPE WebKit
// Homepage: https://wpewebkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibwpe() {
	// Método 1: Descargar y extraer .tar.gz
	libwpe_tar_url := "https://github.com/WebPlatformForEmbedded/libwpe/releases/download/1.16.0/libwpe-1.16.0.tar.xz"
	libwpe_cmd_tar := exec.Command("curl", "-L", libwpe_tar_url, "-o", "package.tar.gz")
	err := libwpe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libwpe_zip_url := "https://github.com/WebPlatformForEmbedded/libwpe/releases/download/1.16.0/libwpe-1.16.0.tar.xz"
	libwpe_cmd_zip := exec.Command("curl", "-L", libwpe_zip_url, "-o", "package.zip")
	err = libwpe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libwpe_bin_url := "https://github.com/WebPlatformForEmbedded/libwpe/releases/download/1.16.0/libwpe-1.16.0.tar.xz"
	libwpe_cmd_bin := exec.Command("curl", "-L", libwpe_bin_url, "-o", "binary.bin")
	err = libwpe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libwpe_src_url := "https://github.com/WebPlatformForEmbedded/libwpe/releases/download/1.16.0/libwpe-1.16.0.tar.xz"
	libwpe_cmd_src := exec.Command("curl", "-L", libwpe_src_url, "-o", "source.tar.gz")
	err = libwpe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libwpe_cmd_direct := exec.Command("./binary")
	err = libwpe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxkbcommon")
exec.Command("latte", "install", "libxkbcommon")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
