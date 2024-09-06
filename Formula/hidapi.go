package main

// Hidapi - Library for communicating with USB and Bluetooth HID devices
// Homepage: https://github.com/libusb/hidapi

import (
	"fmt"
	
	"os/exec"
)

func installHidapi() {
	// Método 1: Descargar y extraer .tar.gz
	hidapi_tar_url := "https://github.com/libusb/hidapi/archive/refs/tags/hidapi-0.14.0.tar.gz"
	hidapi_cmd_tar := exec.Command("curl", "-L", hidapi_tar_url, "-o", "package.tar.gz")
	err := hidapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hidapi_zip_url := "https://github.com/libusb/hidapi/archive/refs/tags/hidapi-0.14.0.zip"
	hidapi_cmd_zip := exec.Command("curl", "-L", hidapi_zip_url, "-o", "package.zip")
	err = hidapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hidapi_bin_url := "https://github.com/libusb/hidapi/archive/refs/tags/hidapi-0.14.0.bin"
	hidapi_cmd_bin := exec.Command("curl", "-L", hidapi_bin_url, "-o", "binary.bin")
	err = hidapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hidapi_src_url := "https://github.com/libusb/hidapi/archive/refs/tags/hidapi-0.14.0.src.tar.gz"
	hidapi_cmd_src := exec.Command("curl", "-L", hidapi_src_url, "-o", "source.tar.gz")
	err = hidapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hidapi_cmd_direct := exec.Command("./binary")
	err = hidapi_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
