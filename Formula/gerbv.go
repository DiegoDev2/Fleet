package main

// Gerbv - Gerber (RS-274X) viewer
// Homepage: https://gerbv.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installGerbv() {
	// Método 1: Descargar y extraer .tar.gz
	gerbv_tar_url := "https://github.com/gerbv/gerbv/archive/refs/tags/v2.10.0.tar.gz"
	gerbv_cmd_tar := exec.Command("curl", "-L", gerbv_tar_url, "-o", "package.tar.gz")
	err := gerbv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gerbv_zip_url := "https://github.com/gerbv/gerbv/archive/refs/tags/v2.10.0.zip"
	gerbv_cmd_zip := exec.Command("curl", "-L", gerbv_zip_url, "-o", "package.zip")
	err = gerbv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gerbv_bin_url := "https://github.com/gerbv/gerbv/archive/refs/tags/v2.10.0.bin"
	gerbv_cmd_bin := exec.Command("curl", "-L", gerbv_bin_url, "-o", "binary.bin")
	err = gerbv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gerbv_src_url := "https://github.com/gerbv/gerbv/archive/refs/tags/v2.10.0.src.tar.gz"
	gerbv_cmd_src := exec.Command("curl", "-L", gerbv_src_url, "-o", "source.tar.gz")
	err = gerbv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gerbv_cmd_direct := exec.Command("./binary")
	err = gerbv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
