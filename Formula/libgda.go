package main

// Libgda - Provides unified data access to the GNOME project
// Homepage: https://www.gnome-db.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibgda() {
	// Método 1: Descargar y extraer .tar.gz
	libgda_tar_url := "https://download.gnome.org/sources/libgda/6.0/libgda-6.0.0.tar.xz"
	libgda_cmd_tar := exec.Command("curl", "-L", libgda_tar_url, "-o", "package.tar.gz")
	err := libgda_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgda_zip_url := "https://download.gnome.org/sources/libgda/6.0/libgda-6.0.0.tar.xz"
	libgda_cmd_zip := exec.Command("curl", "-L", libgda_zip_url, "-o", "package.zip")
	err = libgda_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgda_bin_url := "https://download.gnome.org/sources/libgda/6.0/libgda-6.0.0.tar.xz"
	libgda_cmd_bin := exec.Command("curl", "-L", libgda_bin_url, "-o", "binary.bin")
	err = libgda_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgda_src_url := "https://download.gnome.org/sources/libgda/6.0/libgda-6.0.0.tar.xz"
	libgda_cmd_src := exec.Command("curl", "-L", libgda_src_url, "-o", "source.tar.gz")
	err = libgda_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgda_cmd_direct := exec.Command("./binary")
	err = libgda_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
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
	fmt.Println("Instalando dependencia: iso-codes")
	exec.Command("latte", "install", "iso-codes").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
