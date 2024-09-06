package main

// Waffle - C library for selecting an OpenGL API and window system at runtime
// Homepage: https://waffle.freedesktop.org/

import (
	"fmt"
	
	"os/exec"
)

func installWaffle() {
	// Método 1: Descargar y extraer .tar.gz
	waffle_tar_url := "https://waffle.freedesktop.org/files/release/waffle-1.8.1/waffle-1.8.1.tar.xz"
	waffle_cmd_tar := exec.Command("curl", "-L", waffle_tar_url, "-o", "package.tar.gz")
	err := waffle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	waffle_zip_url := "https://waffle.freedesktop.org/files/release/waffle-1.8.1/waffle-1.8.1.tar.xz"
	waffle_cmd_zip := exec.Command("curl", "-L", waffle_zip_url, "-o", "package.zip")
	err = waffle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	waffle_bin_url := "https://waffle.freedesktop.org/files/release/waffle-1.8.1/waffle-1.8.1.tar.xz"
	waffle_cmd_bin := exec.Command("curl", "-L", waffle_bin_url, "-o", "binary.bin")
	err = waffle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	waffle_src_url := "https://waffle.freedesktop.org/files/release/waffle-1.8.1/waffle-1.8.1.tar.xz"
	waffle_cmd_src := exec.Command("curl", "-L", waffle_src_url, "-o", "source.tar.gz")
	err = waffle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	waffle_cmd_direct := exec.Command("./binary")
	err = waffle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
}
