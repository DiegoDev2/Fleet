package main

// Librealsense - Intel RealSense D400 series and SR300 capture
// Homepage: https://github.com/IntelRealSense/librealsense

import (
	"fmt"
	
	"os/exec"
)

func installLibrealsense() {
	// Método 1: Descargar y extraer .tar.gz
	librealsense_tar_url := "https://github.com/IntelRealSense/librealsense/archive/refs/tags/v2.55.1.tar.gz"
	librealsense_cmd_tar := exec.Command("curl", "-L", librealsense_tar_url, "-o", "package.tar.gz")
	err := librealsense_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librealsense_zip_url := "https://github.com/IntelRealSense/librealsense/archive/refs/tags/v2.55.1.zip"
	librealsense_cmd_zip := exec.Command("curl", "-L", librealsense_zip_url, "-o", "package.zip")
	err = librealsense_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librealsense_bin_url := "https://github.com/IntelRealSense/librealsense/archive/refs/tags/v2.55.1.bin"
	librealsense_cmd_bin := exec.Command("curl", "-L", librealsense_bin_url, "-o", "binary.bin")
	err = librealsense_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librealsense_src_url := "https://github.com/IntelRealSense/librealsense/archive/refs/tags/v2.55.1.src.tar.gz"
	librealsense_cmd_src := exec.Command("curl", "-L", librealsense_src_url, "-o", "source.tar.gz")
	err = librealsense_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librealsense_cmd_direct := exec.Command("./binary")
	err = librealsense_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glfw")
exec.Command("latte", "install", "glfw")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
