package main

// Libuvc - Cross-platform library for USB video devices
// Homepage: https://github.com/libuvc/libuvc

import (
	"fmt"
	
	"os/exec"
)

func installLibuvc() {
	// Método 1: Descargar y extraer .tar.gz
	libuvc_tar_url := "https://github.com/libuvc/libuvc/archive/refs/tags/v0.0.7.tar.gz"
	libuvc_cmd_tar := exec.Command("curl", "-L", libuvc_tar_url, "-o", "package.tar.gz")
	err := libuvc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libuvc_zip_url := "https://github.com/libuvc/libuvc/archive/refs/tags/v0.0.7.zip"
	libuvc_cmd_zip := exec.Command("curl", "-L", libuvc_zip_url, "-o", "package.zip")
	err = libuvc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libuvc_bin_url := "https://github.com/libuvc/libuvc/archive/refs/tags/v0.0.7.bin"
	libuvc_cmd_bin := exec.Command("curl", "-L", libuvc_bin_url, "-o", "binary.bin")
	err = libuvc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libuvc_src_url := "https://github.com/libuvc/libuvc/archive/refs/tags/v0.0.7.src.tar.gz"
	libuvc_cmd_src := exec.Command("curl", "-L", libuvc_src_url, "-o", "source.tar.gz")
	err = libuvc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libuvc_cmd_direct := exec.Command("./binary")
	err = libuvc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
