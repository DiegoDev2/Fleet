package main

// Libgosu - 2D game development library
// Homepage: https://libgosu.org

import (
	"fmt"
	
	"os/exec"
)

func installLibgosu() {
	// Método 1: Descargar y extraer .tar.gz
	libgosu_tar_url := "https://github.com/gosu/gosu/archive/refs/tags/v1.4.6.tar.gz"
	libgosu_cmd_tar := exec.Command("curl", "-L", libgosu_tar_url, "-o", "package.tar.gz")
	err := libgosu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgosu_zip_url := "https://github.com/gosu/gosu/archive/refs/tags/v1.4.6.zip"
	libgosu_cmd_zip := exec.Command("curl", "-L", libgosu_zip_url, "-o", "package.zip")
	err = libgosu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgosu_bin_url := "https://github.com/gosu/gosu/archive/refs/tags/v1.4.6.bin"
	libgosu_cmd_bin := exec.Command("curl", "-L", libgosu_bin_url, "-o", "binary.bin")
	err = libgosu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgosu_src_url := "https://github.com/gosu/gosu/archive/refs/tags/v1.4.6.src.tar.gz"
	libgosu_cmd_src := exec.Command("curl", "-L", libgosu_src_url, "-o", "source.tar.gz")
	err = libgosu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgosu_cmd_direct := exec.Command("./binary")
	err = libgosu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
	fmt.Println("Instalando dependencia: openal-soft")
exec.Command("latte", "install", "openal-soft")
}
