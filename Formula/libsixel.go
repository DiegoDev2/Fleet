package main

// Libsixel - SIXEL encoder/decoder implementation
// Homepage: https://github.com/saitoha/sixel

import (
	"fmt"
	
	"os/exec"
)

func installLibsixel() {
	// Método 1: Descargar y extraer .tar.gz
	libsixel_tar_url := "https://github.com/libsixel/libsixel/archive/refs/tags/v1.10.3.tar.gz"
	libsixel_cmd_tar := exec.Command("curl", "-L", libsixel_tar_url, "-o", "package.tar.gz")
	err := libsixel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsixel_zip_url := "https://github.com/libsixel/libsixel/archive/refs/tags/v1.10.3.zip"
	libsixel_cmd_zip := exec.Command("curl", "-L", libsixel_zip_url, "-o", "package.zip")
	err = libsixel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsixel_bin_url := "https://github.com/libsixel/libsixel/archive/refs/tags/v1.10.3.bin"
	libsixel_cmd_bin := exec.Command("curl", "-L", libsixel_bin_url, "-o", "binary.bin")
	err = libsixel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsixel_src_url := "https://github.com/libsixel/libsixel/archive/refs/tags/v1.10.3.src.tar.gz"
	libsixel_cmd_src := exec.Command("curl", "-L", libsixel_src_url, "-o", "source.tar.gz")
	err = libsixel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsixel_cmd_direct := exec.Command("./binary")
	err = libsixel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
