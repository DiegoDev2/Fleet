package main

// Libspng - C library for reading and writing PNG format files
// Homepage: https://libspng.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibspng() {
	// Método 1: Descargar y extraer .tar.gz
	libspng_tar_url := "https://github.com/randy408/libspng/archive/refs/tags/v0.7.4.tar.gz"
	libspng_cmd_tar := exec.Command("curl", "-L", libspng_tar_url, "-o", "package.tar.gz")
	err := libspng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspng_zip_url := "https://github.com/randy408/libspng/archive/refs/tags/v0.7.4.zip"
	libspng_cmd_zip := exec.Command("curl", "-L", libspng_zip_url, "-o", "package.zip")
	err = libspng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspng_bin_url := "https://github.com/randy408/libspng/archive/refs/tags/v0.7.4.bin"
	libspng_cmd_bin := exec.Command("curl", "-L", libspng_bin_url, "-o", "binary.bin")
	err = libspng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspng_src_url := "https://github.com/randy408/libspng/archive/refs/tags/v0.7.4.src.tar.gz"
	libspng_cmd_src := exec.Command("curl", "-L", libspng_src_url, "-o", "source.tar.gz")
	err = libspng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspng_cmd_direct := exec.Command("./binary")
	err = libspng_cmd_direct.Run()
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
