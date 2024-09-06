package main

// Pygobject3 - GNOME Python bindings (based on GObject Introspection)
// Homepage: https://pygobject.gnome.org

import (
	"fmt"
	
	"os/exec"
)

func installPygobject3() {
	// Método 1: Descargar y extraer .tar.gz
	pygobject3_tar_url := "https://download.gnome.org/sources/pygobject/3.48/pygobject-3.48.2.tar.xz"
	pygobject3_cmd_tar := exec.Command("curl", "-L", pygobject3_tar_url, "-o", "package.tar.gz")
	err := pygobject3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pygobject3_zip_url := "https://download.gnome.org/sources/pygobject/3.48/pygobject-3.48.2.tar.xz"
	pygobject3_cmd_zip := exec.Command("curl", "-L", pygobject3_zip_url, "-o", "package.zip")
	err = pygobject3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pygobject3_bin_url := "https://download.gnome.org/sources/pygobject/3.48/pygobject-3.48.2.tar.xz"
	pygobject3_cmd_bin := exec.Command("curl", "-L", pygobject3_bin_url, "-o", "binary.bin")
	err = pygobject3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pygobject3_src_url := "https://download.gnome.org/sources/pygobject/3.48/pygobject-3.48.2.tar.xz"
	pygobject3_cmd_src := exec.Command("curl", "-L", pygobject3_src_url, "-o", "source.tar.gz")
	err = pygobject3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pygobject3_cmd_direct := exec.Command("./binary")
	err = pygobject3_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: py3cairo")
	exec.Command("latte", "install", "py3cairo").Run()
}
