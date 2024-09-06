package main

// Libslirp - General purpose TCP-IP emulator
// Homepage: https://gitlab.freedesktop.org/slirp/libslirp

import (
	"fmt"
	
	"os/exec"
)

func installLibslirp() {
	// Método 1: Descargar y extraer .tar.gz
	libslirp_tar_url := "https://gitlab.freedesktop.org/slirp/libslirp/-/archive/v4.8.0/libslirp-v4.8.0.tar.gz"
	libslirp_cmd_tar := exec.Command("curl", "-L", libslirp_tar_url, "-o", "package.tar.gz")
	err := libslirp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libslirp_zip_url := "https://gitlab.freedesktop.org/slirp/libslirp/-/archive/v4.8.0/libslirp-v4.8.0.zip"
	libslirp_cmd_zip := exec.Command("curl", "-L", libslirp_zip_url, "-o", "package.zip")
	err = libslirp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libslirp_bin_url := "https://gitlab.freedesktop.org/slirp/libslirp/-/archive/v4.8.0/libslirp-v4.8.0.bin"
	libslirp_cmd_bin := exec.Command("curl", "-L", libslirp_bin_url, "-o", "binary.bin")
	err = libslirp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libslirp_src_url := "https://gitlab.freedesktop.org/slirp/libslirp/-/archive/v4.8.0/libslirp-v4.8.0.src.tar.gz"
	libslirp_cmd_src := exec.Command("curl", "-L", libslirp_src_url, "-o", "source.tar.gz")
	err = libslirp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libslirp_cmd_direct := exec.Command("./binary")
	err = libslirp_cmd_direct.Run()
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
}
