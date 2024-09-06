package main

// Libproxy - Library that provides automatic proxy configuration management
// Homepage: https://libproxy.github.io/libproxy/

import (
	"fmt"
	
	"os/exec"
)

func installLibproxy() {
	// Método 1: Descargar y extraer .tar.gz
	libproxy_tar_url := "https://github.com/libproxy/libproxy/archive/refs/tags/0.5.8.tar.gz"
	libproxy_cmd_tar := exec.Command("curl", "-L", libproxy_tar_url, "-o", "package.tar.gz")
	err := libproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libproxy_zip_url := "https://github.com/libproxy/libproxy/archive/refs/tags/0.5.8.zip"
	libproxy_cmd_zip := exec.Command("curl", "-L", libproxy_zip_url, "-o", "package.zip")
	err = libproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libproxy_bin_url := "https://github.com/libproxy/libproxy/archive/refs/tags/0.5.8.bin"
	libproxy_cmd_bin := exec.Command("curl", "-L", libproxy_bin_url, "-o", "binary.bin")
	err = libproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libproxy_src_url := "https://github.com/libproxy/libproxy/archive/refs/tags/0.5.8.src.tar.gz"
	libproxy_cmd_src := exec.Command("curl", "-L", libproxy_src_url, "-o", "source.tar.gz")
	err = libproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libproxy_cmd_direct := exec.Command("./binary")
	err = libproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gsettings-desktop-schemas")
exec.Command("latte", "install", "gsettings-desktop-schemas")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: duktape")
exec.Command("latte", "install", "duktape")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
}
