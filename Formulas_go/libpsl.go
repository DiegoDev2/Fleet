package main

// Libpsl - C library for the Public Suffix List
// Homepage: https://rockdaboot.github.io/libpsl

import (
	"fmt"
	
	"os/exec"
)

func installLibpsl() {
	// Método 1: Descargar y extraer .tar.gz
	libpsl_tar_url := "https://github.com/rockdaboot/libpsl/releases/download/0.21.5/libpsl-0.21.5.tar.gz"
	libpsl_cmd_tar := exec.Command("curl", "-L", libpsl_tar_url, "-o", "package.tar.gz")
	err := libpsl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpsl_zip_url := "https://github.com/rockdaboot/libpsl/releases/download/0.21.5/libpsl-0.21.5.zip"
	libpsl_cmd_zip := exec.Command("curl", "-L", libpsl_zip_url, "-o", "package.zip")
	err = libpsl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpsl_bin_url := "https://github.com/rockdaboot/libpsl/releases/download/0.21.5/libpsl-0.21.5.bin"
	libpsl_cmd_bin := exec.Command("curl", "-L", libpsl_bin_url, "-o", "binary.bin")
	err = libpsl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpsl_src_url := "https://github.com/rockdaboot/libpsl/releases/download/0.21.5/libpsl-0.21.5.src.tar.gz"
	libpsl_cmd_src := exec.Command("curl", "-L", libpsl_src_url, "-o", "source.tar.gz")
	err = libpsl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpsl_cmd_direct := exec.Command("./binary")
	err = libpsl_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
}
