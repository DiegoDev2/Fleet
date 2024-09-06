package main

// Gcab - Windows installer (.MSI) tool
// Homepage: https://wiki.gnome.org/msitools

import (
	"fmt"
	
	"os/exec"
)

func installGcab() {
	// Método 1: Descargar y extraer .tar.gz
	gcab_tar_url := "https://download.gnome.org/sources/gcab/1.6/gcab-1.6.tar.xz"
	gcab_cmd_tar := exec.Command("curl", "-L", gcab_tar_url, "-o", "package.tar.gz")
	err := gcab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gcab_zip_url := "https://download.gnome.org/sources/gcab/1.6/gcab-1.6.tar.xz"
	gcab_cmd_zip := exec.Command("curl", "-L", gcab_zip_url, "-o", "package.zip")
	err = gcab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gcab_bin_url := "https://download.gnome.org/sources/gcab/1.6/gcab-1.6.tar.xz"
	gcab_cmd_bin := exec.Command("curl", "-L", gcab_bin_url, "-o", "binary.bin")
	err = gcab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gcab_src_url := "https://download.gnome.org/sources/gcab/1.6/gcab-1.6.tar.xz"
	gcab_cmd_src := exec.Command("curl", "-L", gcab_src_url, "-o", "source.tar.gz")
	err = gcab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gcab_cmd_direct := exec.Command("./binary")
	err = gcab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
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
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
