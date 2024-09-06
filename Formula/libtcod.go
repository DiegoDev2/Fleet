package main

// Libtcod - API for roguelike developers
// Homepage: https://github.com/libtcod/libtcod

import (
	"fmt"
	
	"os/exec"
)

func installLibtcod() {
	// Método 1: Descargar y extraer .tar.gz
	libtcod_tar_url := "https://github.com/libtcod/libtcod/archive/refs/tags/1.24.0.tar.gz"
	libtcod_cmd_tar := exec.Command("curl", "-L", libtcod_tar_url, "-o", "package.tar.gz")
	err := libtcod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtcod_zip_url := "https://github.com/libtcod/libtcod/archive/refs/tags/1.24.0.zip"
	libtcod_cmd_zip := exec.Command("curl", "-L", libtcod_zip_url, "-o", "package.zip")
	err = libtcod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtcod_bin_url := "https://github.com/libtcod/libtcod/archive/refs/tags/1.24.0.bin"
	libtcod_cmd_bin := exec.Command("curl", "-L", libtcod_bin_url, "-o", "binary.bin")
	err = libtcod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtcod_src_url := "https://github.com/libtcod/libtcod/archive/refs/tags/1.24.0.src.tar.gz"
	libtcod_cmd_src := exec.Command("curl", "-L", libtcod_src_url, "-o", "source.tar.gz")
	err = libtcod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtcod_cmd_direct := exec.Command("./binary")
	err = libtcod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
