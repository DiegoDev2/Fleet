package main

// Zbar - Suite of barcodes-reading tools
// Homepage: https://github.com/mchehab/zbar

import (
	"fmt"
	
	"os/exec"
)

func installZbar() {
	// Método 1: Descargar y extraer .tar.gz
	zbar_tar_url := "https://linuxtv.org/downloads/zbar/zbar-0.23.93.tar.bz2"
	zbar_cmd_tar := exec.Command("curl", "-L", zbar_tar_url, "-o", "package.tar.gz")
	err := zbar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zbar_zip_url := "https://linuxtv.org/downloads/zbar/zbar-0.23.93.tar.bz2"
	zbar_cmd_zip := exec.Command("curl", "-L", zbar_zip_url, "-o", "package.zip")
	err = zbar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zbar_bin_url := "https://linuxtv.org/downloads/zbar/zbar-0.23.93.tar.bz2"
	zbar_cmd_bin := exec.Command("curl", "-L", zbar_bin_url, "-o", "binary.bin")
	err = zbar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zbar_src_url := "https://linuxtv.org/downloads/zbar/zbar-0.23.93.tar.bz2"
	zbar_cmd_src := exec.Command("curl", "-L", zbar_src_url, "-o", "source.tar.gz")
	err = zbar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zbar_cmd_direct := exec.Command("./binary")
	err = zbar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: liblqr")
exec.Command("latte", "install", "liblqr")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
}
