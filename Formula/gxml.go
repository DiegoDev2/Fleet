package main

// Gxml - GObject-based XML DOM API
// Homepage: https://wiki.gnome.org/GXml

import (
	"fmt"
	
	"os/exec"
)

func installGxml() {
	// Método 1: Descargar y extraer .tar.gz
	gxml_tar_url := "https://gitlab.gnome.org/GNOME/gxml/-/archive/0.20.4/gxml-0.20.4.tar.bz2"
	gxml_cmd_tar := exec.Command("curl", "-L", gxml_tar_url, "-o", "package.tar.gz")
	err := gxml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gxml_zip_url := "https://gitlab.gnome.org/GNOME/gxml/-/archive/0.20.4/gxml-0.20.4.tar.bz2"
	gxml_cmd_zip := exec.Command("curl", "-L", gxml_zip_url, "-o", "package.zip")
	err = gxml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gxml_bin_url := "https://gitlab.gnome.org/GNOME/gxml/-/archive/0.20.4/gxml-0.20.4.tar.bz2"
	gxml_cmd_bin := exec.Command("curl", "-L", gxml_bin_url, "-o", "binary.bin")
	err = gxml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gxml_src_url := "https://gitlab.gnome.org/GNOME/gxml/-/archive/0.20.4/gxml-0.20.4.tar.bz2"
	gxml_cmd_src := exec.Command("curl", "-L", gxml_src_url, "-o", "source.tar.gz")
	err = gxml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gxml_cmd_direct := exec.Command("./binary")
	err = gxml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libgee")
	exec.Command("latte", "install", "libgee").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
