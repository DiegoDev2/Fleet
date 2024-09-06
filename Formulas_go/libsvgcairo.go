package main

// LibsvgCairo - SVG rendering library using Cairo
// Homepage: https://cairographics.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibsvgCairo() {
	// Método 1: Descargar y extraer .tar.gz
	libsvgcairo_tar_url := "https://cairographics.org/snapshots/libsvg-cairo-0.1.6.tar.gz"
	libsvgcairo_cmd_tar := exec.Command("curl", "-L", libsvgcairo_tar_url, "-o", "package.tar.gz")
	err := libsvgcairo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsvgcairo_zip_url := "https://cairographics.org/snapshots/libsvg-cairo-0.1.6.zip"
	libsvgcairo_cmd_zip := exec.Command("curl", "-L", libsvgcairo_zip_url, "-o", "package.zip")
	err = libsvgcairo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsvgcairo_bin_url := "https://cairographics.org/snapshots/libsvg-cairo-0.1.6.bin"
	libsvgcairo_cmd_bin := exec.Command("curl", "-L", libsvgcairo_bin_url, "-o", "binary.bin")
	err = libsvgcairo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsvgcairo_src_url := "https://cairographics.org/snapshots/libsvg-cairo-0.1.6.src.tar.gz"
	libsvgcairo_cmd_src := exec.Command("curl", "-L", libsvgcairo_src_url, "-o", "source.tar.gz")
	err = libsvgcairo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsvgcairo_cmd_direct := exec.Command("./binary")
	err = libsvgcairo_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsvg")
exec.Command("latte", "install", "libsvg")
}
