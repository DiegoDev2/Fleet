package main

// WpebackendFdo - Freedesktop.org backend for WPE WebKit
// Homepage: https://wpewebkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installWpebackendFdo() {
	// Método 1: Descargar y extraer .tar.gz
	wpebackendfdo_tar_url := "https://github.com/Igalia/WPEBackend-fdo/releases/download/1.14.2/wpebackend-fdo-1.14.2.tar.xz"
	wpebackendfdo_cmd_tar := exec.Command("curl", "-L", wpebackendfdo_tar_url, "-o", "package.tar.gz")
	err := wpebackendfdo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wpebackendfdo_zip_url := "https://github.com/Igalia/WPEBackend-fdo/releases/download/1.14.2/wpebackend-fdo-1.14.2.tar.xz"
	wpebackendfdo_cmd_zip := exec.Command("curl", "-L", wpebackendfdo_zip_url, "-o", "package.zip")
	err = wpebackendfdo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wpebackendfdo_bin_url := "https://github.com/Igalia/WPEBackend-fdo/releases/download/1.14.2/wpebackend-fdo-1.14.2.tar.xz"
	wpebackendfdo_cmd_bin := exec.Command("curl", "-L", wpebackendfdo_bin_url, "-o", "binary.bin")
	err = wpebackendfdo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wpebackendfdo_src_url := "https://github.com/Igalia/WPEBackend-fdo/releases/download/1.14.2/wpebackend-fdo-1.14.2.tar.xz"
	wpebackendfdo_cmd_src := exec.Command("curl", "-L", wpebackendfdo_src_url, "-o", "source.tar.gz")
	err = wpebackendfdo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wpebackendfdo_cmd_direct := exec.Command("./binary")
	err = wpebackendfdo_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libepoxy")
	exec.Command("latte", "install", "libepoxy").Run()
	fmt.Println("Instalando dependencia: libwpe")
	exec.Command("latte", "install", "libwpe").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
