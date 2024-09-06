package main

// Bootloadhid - HID-based USB bootloader for AVR microcontrollers
// Homepage: https://www.obdev.at/products/vusb/bootloadhid.html

import (
	"fmt"
	
	"os/exec"
)

func installBootloadhid() {
	// Método 1: Descargar y extraer .tar.gz
	bootloadhid_tar_url := "https://www.obdev.at/downloads/vusb/bootloadHID.2012-12-08.tar.gz"
	bootloadhid_cmd_tar := exec.Command("curl", "-L", bootloadhid_tar_url, "-o", "package.tar.gz")
	err := bootloadhid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bootloadhid_zip_url := "https://www.obdev.at/downloads/vusb/bootloadHID.2012-12-08.zip"
	bootloadhid_cmd_zip := exec.Command("curl", "-L", bootloadhid_zip_url, "-o", "package.zip")
	err = bootloadhid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bootloadhid_bin_url := "https://www.obdev.at/downloads/vusb/bootloadHID.2012-12-08.bin"
	bootloadhid_cmd_bin := exec.Command("curl", "-L", bootloadhid_bin_url, "-o", "binary.bin")
	err = bootloadhid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bootloadhid_src_url := "https://www.obdev.at/downloads/vusb/bootloadHID.2012-12-08.src.tar.gz"
	bootloadhid_cmd_src := exec.Command("curl", "-L", bootloadhid_src_url, "-o", "source.tar.gz")
	err = bootloadhid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bootloadhid_cmd_direct := exec.Command("./binary")
	err = bootloadhid_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
