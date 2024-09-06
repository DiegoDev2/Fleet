package main

// Libvatek - User library to control VATek chips
// Homepage: https://github.com/VisionAdvanceTechnologyInc/vatek_sdk_2

import (
	"fmt"
	
	"os/exec"
)

func installLibvatek() {
	// Método 1: Descargar y extraer .tar.gz
	libvatek_tar_url := "https://github.com/VisionAdvanceTechnologyInc/vatek_sdk_2/archive/refs/tags/v3.10.tar.gz"
	libvatek_cmd_tar := exec.Command("curl", "-L", libvatek_tar_url, "-o", "package.tar.gz")
	err := libvatek_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvatek_zip_url := "https://github.com/VisionAdvanceTechnologyInc/vatek_sdk_2/archive/refs/tags/v3.10.zip"
	libvatek_cmd_zip := exec.Command("curl", "-L", libvatek_zip_url, "-o", "package.zip")
	err = libvatek_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvatek_bin_url := "https://github.com/VisionAdvanceTechnologyInc/vatek_sdk_2/archive/refs/tags/v3.10.bin"
	libvatek_cmd_bin := exec.Command("curl", "-L", libvatek_bin_url, "-o", "binary.bin")
	err = libvatek_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvatek_src_url := "https://github.com/VisionAdvanceTechnologyInc/vatek_sdk_2/archive/refs/tags/v3.10.src.tar.gz"
	libvatek_cmd_src := exec.Command("curl", "-L", libvatek_src_url, "-o", "source.tar.gz")
	err = libvatek_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvatek_cmd_direct := exec.Command("./binary")
	err = libvatek_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
