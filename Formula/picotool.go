package main

// Picotool - Tool for interacting with RP2040/RP2350 devices and binaries
// Homepage: https://github.com/raspberrypi/picotool

import (
	"fmt"
	
	"os/exec"
)

func installPicotool() {
	// Método 1: Descargar y extraer .tar.gz
	picotool_tar_url := "https://github.com/raspberrypi/picotool/archive/refs/tags/2.0.0.tar.gz"
	picotool_cmd_tar := exec.Command("curl", "-L", picotool_tar_url, "-o", "package.tar.gz")
	err := picotool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	picotool_zip_url := "https://github.com/raspberrypi/picotool/archive/refs/tags/2.0.0.zip"
	picotool_cmd_zip := exec.Command("curl", "-L", picotool_zip_url, "-o", "package.zip")
	err = picotool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	picotool_bin_url := "https://github.com/raspberrypi/picotool/archive/refs/tags/2.0.0.bin"
	picotool_cmd_bin := exec.Command("curl", "-L", picotool_bin_url, "-o", "binary.bin")
	err = picotool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	picotool_src_url := "https://github.com/raspberrypi/picotool/archive/refs/tags/2.0.0.src.tar.gz"
	picotool_cmd_src := exec.Command("curl", "-L", picotool_src_url, "-o", "source.tar.gz")
	err = picotool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	picotool_cmd_direct := exec.Command("./binary")
	err = picotool_cmd_direct.Run()
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
