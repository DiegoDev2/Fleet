package main

// Libcroco - CSS parsing and manipulation toolkit for GNOME
// Homepage: https://gitlab.gnome.org/Archive/libcroco

import (
	"fmt"
	
	"os/exec"
)

func installLibcroco() {
	// Método 1: Descargar y extraer .tar.gz
	libcroco_tar_url := "https://download.gnome.org/sources/libcroco/0.6/libcroco-0.6.13.tar.xz"
	libcroco_cmd_tar := exec.Command("curl", "-L", libcroco_tar_url, "-o", "package.tar.gz")
	err := libcroco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcroco_zip_url := "https://download.gnome.org/sources/libcroco/0.6/libcroco-0.6.13.tar.xz"
	libcroco_cmd_zip := exec.Command("curl", "-L", libcroco_zip_url, "-o", "package.zip")
	err = libcroco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcroco_bin_url := "https://download.gnome.org/sources/libcroco/0.6/libcroco-0.6.13.tar.xz"
	libcroco_cmd_bin := exec.Command("curl", "-L", libcroco_bin_url, "-o", "binary.bin")
	err = libcroco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcroco_src_url := "https://download.gnome.org/sources/libcroco/0.6/libcroco-0.6.13.tar.xz"
	libcroco_cmd_src := exec.Command("curl", "-L", libcroco_src_url, "-o", "source.tar.gz")
	err = libcroco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcroco_cmd_direct := exec.Command("./binary")
	err = libcroco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
