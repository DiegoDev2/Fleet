package main

// Gexiv2 - GObject wrapper around the Exiv2 photo metadata library
// Homepage: https://wiki.gnome.org/Projects/gexiv2

import (
	"fmt"
	
	"os/exec"
)

func installGexiv2() {
	// Método 1: Descargar y extraer .tar.gz
	gexiv2_tar_url := "https://download.gnome.org/sources/gexiv2/0.14/gexiv2-0.14.3.tar.xz"
	gexiv2_cmd_tar := exec.Command("curl", "-L", gexiv2_tar_url, "-o", "package.tar.gz")
	err := gexiv2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gexiv2_zip_url := "https://download.gnome.org/sources/gexiv2/0.14/gexiv2-0.14.3.tar.xz"
	gexiv2_cmd_zip := exec.Command("curl", "-L", gexiv2_zip_url, "-o", "package.zip")
	err = gexiv2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gexiv2_bin_url := "https://download.gnome.org/sources/gexiv2/0.14/gexiv2-0.14.3.tar.xz"
	gexiv2_cmd_bin := exec.Command("curl", "-L", gexiv2_bin_url, "-o", "binary.bin")
	err = gexiv2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gexiv2_src_url := "https://download.gnome.org/sources/gexiv2/0.14/gexiv2-0.14.3.tar.xz"
	gexiv2_cmd_src := exec.Command("curl", "-L", gexiv2_src_url, "-o", "source.tar.gz")
	err = gexiv2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gexiv2_cmd_direct := exec.Command("./binary")
	err = gexiv2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pygobject3")
exec.Command("latte", "install", "pygobject3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: exiv2")
exec.Command("latte", "install", "exiv2")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
