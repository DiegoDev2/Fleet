package main

// Tio - Simple TTY terminal I/O application
// Homepage: https://tio.github.io

import (
	"fmt"
	
	"os/exec"
)

func installTio() {
	// Método 1: Descargar y extraer .tar.gz
	tio_tar_url := "https://github.com/tio/tio/releases/download/v3.7/tio-3.7.tar.xz"
	tio_cmd_tar := exec.Command("curl", "-L", tio_tar_url, "-o", "package.tar.gz")
	err := tio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tio_zip_url := "https://github.com/tio/tio/releases/download/v3.7/tio-3.7.tar.xz"
	tio_cmd_zip := exec.Command("curl", "-L", tio_zip_url, "-o", "package.zip")
	err = tio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tio_bin_url := "https://github.com/tio/tio/releases/download/v3.7/tio-3.7.tar.xz"
	tio_cmd_bin := exec.Command("curl", "-L", tio_bin_url, "-o", "binary.bin")
	err = tio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tio_src_url := "https://github.com/tio/tio/releases/download/v3.7/tio-3.7.tar.xz"
	tio_cmd_src := exec.Command("curl", "-L", tio_src_url, "-o", "source.tar.gz")
	err = tio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tio_cmd_direct := exec.Command("./binary")
	err = tio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
}
