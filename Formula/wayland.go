package main

// Wayland - Protocol for a compositor to talk to its clients
// Homepage: https://wayland.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installWayland() {
	// Método 1: Descargar y extraer .tar.gz
	wayland_tar_url := "https://gitlab.freedesktop.org/wayland/wayland/-/releases/1.23.0/downloads/wayland-1.23.0.tar.xz"
	wayland_cmd_tar := exec.Command("curl", "-L", wayland_tar_url, "-o", "package.tar.gz")
	err := wayland_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wayland_zip_url := "https://gitlab.freedesktop.org/wayland/wayland/-/releases/1.23.0/downloads/wayland-1.23.0.tar.xz"
	wayland_cmd_zip := exec.Command("curl", "-L", wayland_zip_url, "-o", "package.zip")
	err = wayland_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wayland_bin_url := "https://gitlab.freedesktop.org/wayland/wayland/-/releases/1.23.0/downloads/wayland-1.23.0.tar.xz"
	wayland_cmd_bin := exec.Command("curl", "-L", wayland_bin_url, "-o", "binary.bin")
	err = wayland_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wayland_src_url := "https://gitlab.freedesktop.org/wayland/wayland/-/releases/1.23.0/downloads/wayland-1.23.0.tar.xz"
	wayland_cmd_src := exec.Command("curl", "-L", wayland_src_url, "-o", "source.tar.gz")
	err = wayland_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wayland_cmd_direct := exec.Command("./binary")
	err = wayland_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: expat")
	exec.Command("latte", "install", "expat").Run()
	fmt.Println("Instalando dependencia: libffi")
	exec.Command("latte", "install", "libffi").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
}
