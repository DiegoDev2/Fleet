package main

// UsbIds - Repository of vendor, device, subsystem and device class IDs used in USB devices
// Homepage: http://www.linux-usb.org/usb-ids.html

import (
	"fmt"
	
	"os/exec"
)

func installUsbIds() {
	// Método 1: Descargar y extraer .tar.gz
	usbids_tar_url := "https://deb.debian.org/debian/pool/main/u/usb.ids/usb.ids_2024.07.04.orig.tar.xz"
	usbids_cmd_tar := exec.Command("curl", "-L", usbids_tar_url, "-o", "package.tar.gz")
	err := usbids_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	usbids_zip_url := "https://deb.debian.org/debian/pool/main/u/usb.ids/usb.ids_2024.07.04.orig.tar.xz"
	usbids_cmd_zip := exec.Command("curl", "-L", usbids_zip_url, "-o", "package.zip")
	err = usbids_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	usbids_bin_url := "https://deb.debian.org/debian/pool/main/u/usb.ids/usb.ids_2024.07.04.orig.tar.xz"
	usbids_cmd_bin := exec.Command("curl", "-L", usbids_bin_url, "-o", "binary.bin")
	err = usbids_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	usbids_src_url := "https://deb.debian.org/debian/pool/main/u/usb.ids/usb.ids_2024.07.04.orig.tar.xz"
	usbids_cmd_src := exec.Command("curl", "-L", usbids_src_url, "-o", "source.tar.gz")
	err = usbids_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	usbids_cmd_direct := exec.Command("./binary")
	err = usbids_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
