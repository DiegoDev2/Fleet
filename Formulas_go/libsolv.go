package main

// Libsolv - Library for solving packages and reading repositories
// Homepage: https://github.com/openSUSE/libsolv

import (
	"fmt"
	
	"os/exec"
)

func installLibsolv() {
	// Método 1: Descargar y extraer .tar.gz
	libsolv_tar_url := "https://github.com/openSUSE/libsolv/archive/refs/tags/0.7.30.tar.gz"
	libsolv_cmd_tar := exec.Command("curl", "-L", libsolv_tar_url, "-o", "package.tar.gz")
	err := libsolv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsolv_zip_url := "https://github.com/openSUSE/libsolv/archive/refs/tags/0.7.30.zip"
	libsolv_cmd_zip := exec.Command("curl", "-L", libsolv_zip_url, "-o", "package.zip")
	err = libsolv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsolv_bin_url := "https://github.com/openSUSE/libsolv/archive/refs/tags/0.7.30.bin"
	libsolv_cmd_bin := exec.Command("curl", "-L", libsolv_bin_url, "-o", "binary.bin")
	err = libsolv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsolv_src_url := "https://github.com/openSUSE/libsolv/archive/refs/tags/0.7.30.src.tar.gz"
	libsolv_cmd_src := exec.Command("curl", "-L", libsolv_src_url, "-o", "source.tar.gz")
	err = libsolv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsolv_cmd_direct := exec.Command("./binary")
	err = libsolv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
