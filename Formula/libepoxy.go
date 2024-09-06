package main

// Libepoxy - Library for handling OpenGL function pointer management
// Homepage: https://github.com/anholt/libepoxy

import (
	"fmt"
	
	"os/exec"
)

func installLibepoxy() {
	// Método 1: Descargar y extraer .tar.gz
	libepoxy_tar_url := "https://download.gnome.org/sources/libepoxy/1.5/libepoxy-1.5.10.tar.xz"
	libepoxy_cmd_tar := exec.Command("curl", "-L", libepoxy_tar_url, "-o", "package.tar.gz")
	err := libepoxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libepoxy_zip_url := "https://download.gnome.org/sources/libepoxy/1.5/libepoxy-1.5.10.tar.xz"
	libepoxy_cmd_zip := exec.Command("curl", "-L", libepoxy_zip_url, "-o", "package.zip")
	err = libepoxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libepoxy_bin_url := "https://download.gnome.org/sources/libepoxy/1.5/libepoxy-1.5.10.tar.xz"
	libepoxy_cmd_bin := exec.Command("curl", "-L", libepoxy_bin_url, "-o", "binary.bin")
	err = libepoxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libepoxy_src_url := "https://download.gnome.org/sources/libepoxy/1.5/libepoxy-1.5.10.tar.xz"
	libepoxy_cmd_src := exec.Command("curl", "-L", libepoxy_src_url, "-o", "source.tar.gz")
	err = libepoxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libepoxy_cmd_direct := exec.Command("./binary")
	err = libepoxy_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: freeglut")
	exec.Command("latte", "install", "freeglut").Run()
}
