package main

// Libplacebo - Reusable library for GPU-accelerated image/video processing primitives
// Homepage: https://code.videolan.org/videolan/libplacebo

import (
	"fmt"
	
	"os/exec"
)

func installLibplacebo() {
	// Método 1: Descargar y extraer .tar.gz
	libplacebo_tar_url := "https://code.videolan.org/videolan/libplacebo/-/archive/v7.349.0/libplacebo-v7.349.0.tar.bz2"
	libplacebo_cmd_tar := exec.Command("curl", "-L", libplacebo_tar_url, "-o", "package.tar.gz")
	err := libplacebo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libplacebo_zip_url := "https://code.videolan.org/videolan/libplacebo/-/archive/v7.349.0/libplacebo-v7.349.0.tar.bz2"
	libplacebo_cmd_zip := exec.Command("curl", "-L", libplacebo_zip_url, "-o", "package.zip")
	err = libplacebo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libplacebo_bin_url := "https://code.videolan.org/videolan/libplacebo/-/archive/v7.349.0/libplacebo-v7.349.0.tar.bz2"
	libplacebo_cmd_bin := exec.Command("curl", "-L", libplacebo_bin_url, "-o", "binary.bin")
	err = libplacebo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libplacebo_src_url := "https://code.videolan.org/videolan/libplacebo/-/archive/v7.349.0/libplacebo-v7.349.0.tar.bz2"
	libplacebo_cmd_src := exec.Command("curl", "-L", libplacebo_src_url, "-o", "source.tar.gz")
	err = libplacebo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libplacebo_cmd_direct := exec.Command("./binary")
	err = libplacebo_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vulkan-headers")
exec.Command("latte", "install", "vulkan-headers")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: shaderc")
exec.Command("latte", "install", "shaderc")
	fmt.Println("Instalando dependencia: vulkan-loader")
exec.Command("latte", "install", "vulkan-loader")
}
