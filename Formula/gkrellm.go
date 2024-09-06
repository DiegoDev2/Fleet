package main

// Gkrellm - Extensible GTK system monitoring application
// Homepage: https://billw2.github.io/gkrellm/gkrellm.html

import (
	"fmt"
	
	"os/exec"
)

func installGkrellm() {
	// Método 1: Descargar y extraer .tar.gz
	gkrellm_tar_url := "http://gkrellm.srcbox.net/releases/gkrellm-2.3.11.tar.bz2"
	gkrellm_cmd_tar := exec.Command("curl", "-L", gkrellm_tar_url, "-o", "package.tar.gz")
	err := gkrellm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gkrellm_zip_url := "http://gkrellm.srcbox.net/releases/gkrellm-2.3.11.tar.bz2"
	gkrellm_cmd_zip := exec.Command("curl", "-L", gkrellm_zip_url, "-o", "package.zip")
	err = gkrellm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gkrellm_bin_url := "http://gkrellm.srcbox.net/releases/gkrellm-2.3.11.tar.bz2"
	gkrellm_cmd_bin := exec.Command("curl", "-L", gkrellm_bin_url, "-o", "binary.bin")
	err = gkrellm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gkrellm_src_url := "http://gkrellm.srcbox.net/releases/gkrellm-2.3.11.tar.bz2"
	gkrellm_cmd_src := exec.Command("curl", "-L", gkrellm_src_url, "-o", "source.tar.gz")
	err = gkrellm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gkrellm_cmd_direct := exec.Command("./binary")
	err = gkrellm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
