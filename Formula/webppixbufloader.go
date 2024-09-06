package main

// WebpPixbufLoader - WebP Image format GdkPixbuf loader
// Homepage: https://github.com/aruiz/webp-pixbuf-loader

import (
	"fmt"
	
	"os/exec"
)

func installWebpPixbufLoader() {
	// Método 1: Descargar y extraer .tar.gz
	webppixbufloader_tar_url := "https://github.com/aruiz/webp-pixbuf-loader/archive/refs/tags/0.2.7.tar.gz"
	webppixbufloader_cmd_tar := exec.Command("curl", "-L", webppixbufloader_tar_url, "-o", "package.tar.gz")
	err := webppixbufloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webppixbufloader_zip_url := "https://github.com/aruiz/webp-pixbuf-loader/archive/refs/tags/0.2.7.zip"
	webppixbufloader_cmd_zip := exec.Command("curl", "-L", webppixbufloader_zip_url, "-o", "package.zip")
	err = webppixbufloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webppixbufloader_bin_url := "https://github.com/aruiz/webp-pixbuf-loader/archive/refs/tags/0.2.7.bin"
	webppixbufloader_cmd_bin := exec.Command("curl", "-L", webppixbufloader_bin_url, "-o", "binary.bin")
	err = webppixbufloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webppixbufloader_src_url := "https://github.com/aruiz/webp-pixbuf-loader/archive/refs/tags/0.2.7.src.tar.gz"
	webppixbufloader_cmd_src := exec.Command("curl", "-L", webppixbufloader_src_url, "-o", "source.tar.gz")
	err = webppixbufloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webppixbufloader_cmd_direct := exec.Command("./binary")
	err = webppixbufloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
}
