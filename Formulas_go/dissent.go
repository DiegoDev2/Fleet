package main

// Dissent - GTK4 Discord client in Go
// Homepage: https://github.com/diamondburned/dissent

import (
	"fmt"
	
	"os/exec"
)

func installDissent() {
	// Método 1: Descargar y extraer .tar.gz
	dissent_tar_url := "https://github.com/diamondburned/dissent/archive/refs/tags/v0.0.29.tar.gz"
	dissent_cmd_tar := exec.Command("curl", "-L", dissent_tar_url, "-o", "package.tar.gz")
	err := dissent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dissent_zip_url := "https://github.com/diamondburned/dissent/archive/refs/tags/v0.0.29.zip"
	dissent_cmd_zip := exec.Command("curl", "-L", dissent_zip_url, "-o", "package.zip")
	err = dissent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dissent_bin_url := "https://github.com/diamondburned/dissent/archive/refs/tags/v0.0.29.bin"
	dissent_cmd_bin := exec.Command("curl", "-L", dissent_bin_url, "-o", "binary.bin")
	err = dissent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dissent_src_url := "https://github.com/diamondburned/dissent/archive/refs/tags/v0.0.29.src.tar.gz"
	dissent_cmd_src := exec.Command("curl", "-L", dissent_src_url, "-o", "source.tar.gz")
	err = dissent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dissent_cmd_direct := exec.Command("./binary")
	err = dissent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: graphene")
exec.Command("latte", "install", "graphene")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: gtksourceview5")
exec.Command("latte", "install", "gtksourceview5")
	fmt.Println("Instalando dependencia: libadwaita")
exec.Command("latte", "install", "libadwaita")
	fmt.Println("Instalando dependencia: libcanberra")
exec.Command("latte", "install", "libcanberra")
	fmt.Println("Instalando dependencia: libspelling")
exec.Command("latte", "install", "libspelling")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
