package main

// LibgeditGfls - Gedit Technology - File loading and saving
// Homepage: https://gitlab.gnome.org/World/gedit/libgedit-gfls

import (
	"fmt"
	
	"os/exec"
)

func installLibgeditGfls() {
	// Método 1: Descargar y extraer .tar.gz
	libgeditgfls_tar_url := "https://gitlab.gnome.org/World/gedit/libgedit-gfls/-/archive/0.2.0/libgedit-gfls-0.2.0.tar.bz2"
	libgeditgfls_cmd_tar := exec.Command("curl", "-L", libgeditgfls_tar_url, "-o", "package.tar.gz")
	err := libgeditgfls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgeditgfls_zip_url := "https://gitlab.gnome.org/World/gedit/libgedit-gfls/-/archive/0.2.0/libgedit-gfls-0.2.0.tar.bz2"
	libgeditgfls_cmd_zip := exec.Command("curl", "-L", libgeditgfls_zip_url, "-o", "package.zip")
	err = libgeditgfls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgeditgfls_bin_url := "https://gitlab.gnome.org/World/gedit/libgedit-gfls/-/archive/0.2.0/libgedit-gfls-0.2.0.tar.bz2"
	libgeditgfls_cmd_bin := exec.Command("curl", "-L", libgeditgfls_bin_url, "-o", "binary.bin")
	err = libgeditgfls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgeditgfls_src_url := "https://gitlab.gnome.org/World/gedit/libgedit-gfls/-/archive/0.2.0/libgedit-gfls-0.2.0.tar.bz2"
	libgeditgfls_cmd_src := exec.Command("curl", "-L", libgeditgfls_src_url, "-o", "source.tar.gz")
	err = libgeditgfls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgeditgfls_cmd_direct := exec.Command("./binary")
	err = libgeditgfls_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
}
