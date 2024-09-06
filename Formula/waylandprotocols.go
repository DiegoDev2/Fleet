package main

// WaylandProtocols - Additional Wayland protocols
// Homepage: https://wayland.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installWaylandProtocols() {
	// Método 1: Descargar y extraer .tar.gz
	waylandprotocols_tar_url := "https://gitlab.freedesktop.org/wayland/wayland-protocols/-/releases/1.36/downloads/wayland-protocols-1.36.tar.xz"
	waylandprotocols_cmd_tar := exec.Command("curl", "-L", waylandprotocols_tar_url, "-o", "package.tar.gz")
	err := waylandprotocols_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	waylandprotocols_zip_url := "https://gitlab.freedesktop.org/wayland/wayland-protocols/-/releases/1.36/downloads/wayland-protocols-1.36.tar.xz"
	waylandprotocols_cmd_zip := exec.Command("curl", "-L", waylandprotocols_zip_url, "-o", "package.zip")
	err = waylandprotocols_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	waylandprotocols_bin_url := "https://gitlab.freedesktop.org/wayland/wayland-protocols/-/releases/1.36/downloads/wayland-protocols-1.36.tar.xz"
	waylandprotocols_cmd_bin := exec.Command("curl", "-L", waylandprotocols_bin_url, "-o", "binary.bin")
	err = waylandprotocols_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	waylandprotocols_src_url := "https://gitlab.freedesktop.org/wayland/wayland-protocols/-/releases/1.36/downloads/wayland-protocols-1.36.tar.xz"
	waylandprotocols_cmd_src := exec.Command("curl", "-L", waylandprotocols_src_url, "-o", "source.tar.gz")
	err = waylandprotocols_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	waylandprotocols_cmd_direct := exec.Command("./binary")
	err = waylandprotocols_cmd_direct.Run()
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
}
