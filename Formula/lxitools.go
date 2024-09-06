package main

// LxiTools - Open source tools for managing network attached LXI compatible instruments
// Homepage: https://github.com/lxi-tools/lxi-tools

import (
	"fmt"
	
	"os/exec"
)

func installLxiTools() {
	// Método 1: Descargar y extraer .tar.gz
	lxitools_tar_url := "https://github.com/lxi-tools/lxi-tools/archive/refs/tags/v2.7.tar.gz"
	lxitools_cmd_tar := exec.Command("curl", "-L", lxitools_tar_url, "-o", "package.tar.gz")
	err := lxitools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lxitools_zip_url := "https://github.com/lxi-tools/lxi-tools/archive/refs/tags/v2.7.zip"
	lxitools_cmd_zip := exec.Command("curl", "-L", lxitools_zip_url, "-o", "package.zip")
	err = lxitools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lxitools_bin_url := "https://github.com/lxi-tools/lxi-tools/archive/refs/tags/v2.7.bin"
	lxitools_cmd_bin := exec.Command("curl", "-L", lxitools_bin_url, "-o", "binary.bin")
	err = lxitools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lxitools_src_url := "https://github.com/lxi-tools/lxi-tools/archive/refs/tags/v2.7.src.tar.gz"
	lxitools_cmd_src := exec.Command("curl", "-L", lxitools_src_url, "-o", "source.tar.gz")
	err = lxitools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lxitools_cmd_direct := exec.Command("./binary")
	err = lxitools_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: desktop-file-utils")
	exec.Command("latte", "install", "desktop-file-utils").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: gtksourceview5")
	exec.Command("latte", "install", "gtksourceview5").Run()
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
	exec.Command("latte", "install", "hicolor-icon-theme").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libadwaita")
	exec.Command("latte", "install", "libadwaita").Run()
	fmt.Println("Instalando dependencia: liblxi")
	exec.Command("latte", "install", "liblxi").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: graphene")
	exec.Command("latte", "install", "graphene").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
