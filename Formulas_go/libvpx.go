package main

// Libvpx - VP8/VP9 video codec
// Homepage: https://www.webmproject.org/code/

import (
	"fmt"
	
	"os/exec"
)

func installLibvpx() {
	// Método 1: Descargar y extraer .tar.gz
	libvpx_tar_url := "https://github.com/webmproject/libvpx/archive/refs/tags/v1.13.1.tar.gz"
	libvpx_cmd_tar := exec.Command("curl", "-L", libvpx_tar_url, "-o", "package.tar.gz")
	err := libvpx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvpx_zip_url := "https://github.com/webmproject/libvpx/archive/refs/tags/v1.13.1.zip"
	libvpx_cmd_zip := exec.Command("curl", "-L", libvpx_zip_url, "-o", "package.zip")
	err = libvpx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvpx_bin_url := "https://github.com/webmproject/libvpx/archive/refs/tags/v1.13.1.bin"
	libvpx_cmd_bin := exec.Command("curl", "-L", libvpx_bin_url, "-o", "binary.bin")
	err = libvpx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvpx_src_url := "https://github.com/webmproject/libvpx/archive/refs/tags/v1.13.1.src.tar.gz"
	libvpx_cmd_src := exec.Command("curl", "-L", libvpx_src_url, "-o", "source.tar.gz")
	err = libvpx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvpx_cmd_direct := exec.Command("./binary")
	err = libvpx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: yasm")
exec.Command("latte", "install", "yasm")
}
