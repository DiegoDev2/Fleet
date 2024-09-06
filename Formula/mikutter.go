package main

// Mikutter - Extensible Twitter client
// Homepage: https://mikutter.hachune.net/

import (
	"fmt"
	
	"os/exec"
)

func installMikutter() {
	// Método 1: Descargar y extraer .tar.gz
	mikutter_tar_url := "https://mikutter.hachune.net/bin/mikutter-5.0.7.tar.gz"
	mikutter_cmd_tar := exec.Command("curl", "-L", mikutter_tar_url, "-o", "package.tar.gz")
	err := mikutter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mikutter_zip_url := "https://mikutter.hachune.net/bin/mikutter-5.0.7.zip"
	mikutter_cmd_zip := exec.Command("curl", "-L", mikutter_zip_url, "-o", "package.zip")
	err = mikutter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mikutter_bin_url := "https://mikutter.hachune.net/bin/mikutter-5.0.7.bin"
	mikutter_cmd_bin := exec.Command("curl", "-L", mikutter_bin_url, "-o", "binary.bin")
	err = mikutter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mikutter_src_url := "https://mikutter.hachune.net/bin/mikutter-5.0.7.src.tar.gz"
	mikutter_cmd_src := exec.Command("curl", "-L", mikutter_src_url, "-o", "source.tar.gz")
	err = mikutter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mikutter_cmd_direct := exec.Command("./binary")
	err = mikutter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: terminal-notifier")
	exec.Command("latte", "install", "terminal-notifier").Run()
}
