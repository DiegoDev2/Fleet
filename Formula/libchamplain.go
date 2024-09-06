package main

// Libchamplain - ClutterActor for displaying maps
// Homepage: https://wiki.gnome.org/Projects/libchamplain

import (
	"fmt"
	
	"os/exec"
)

func installLibchamplain() {
	// Método 1: Descargar y extraer .tar.gz
	libchamplain_tar_url := "https://download.gnome.org/sources/libchamplain/0.12/libchamplain-0.12.21.tar.xz"
	libchamplain_cmd_tar := exec.Command("curl", "-L", libchamplain_tar_url, "-o", "package.tar.gz")
	err := libchamplain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libchamplain_zip_url := "https://download.gnome.org/sources/libchamplain/0.12/libchamplain-0.12.21.tar.xz"
	libchamplain_cmd_zip := exec.Command("curl", "-L", libchamplain_zip_url, "-o", "package.zip")
	err = libchamplain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libchamplain_bin_url := "https://download.gnome.org/sources/libchamplain/0.12/libchamplain-0.12.21.tar.xz"
	libchamplain_cmd_bin := exec.Command("curl", "-L", libchamplain_bin_url, "-o", "binary.bin")
	err = libchamplain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libchamplain_src_url := "https://download.gnome.org/sources/libchamplain/0.12/libchamplain-0.12.21.tar.xz"
	libchamplain_cmd_src := exec.Command("curl", "-L", libchamplain_src_url, "-o", "source.tar.gz")
	err = libchamplain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libchamplain_cmd_direct := exec.Command("./binary")
	err = libchamplain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnome-common")
	exec.Command("latte", "install", "gnome-common").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: clutter")
	exec.Command("latte", "install", "clutter").Run()
	fmt.Println("Instalando dependencia: clutter-gtk")
	exec.Command("latte", "install", "clutter-gtk").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: libsoup")
	exec.Command("latte", "install", "libsoup").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
}
