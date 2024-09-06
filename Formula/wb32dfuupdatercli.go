package main

// Wb32DfuUpdaterCli - USB programmer for downloading and uploading firmware to/from USB devices
// Homepage: https://github.com/WestberryTech/wb32-dfu-updater

import (
	"fmt"
	
	"os/exec"
)

func installWb32DfuUpdaterCli() {
	// Método 1: Descargar y extraer .tar.gz
	wb32dfuupdatercli_tar_url := "https://github.com/WestberryTech/wb32-dfu-updater/archive/refs/tags/1.0.0.tar.gz"
	wb32dfuupdatercli_cmd_tar := exec.Command("curl", "-L", wb32dfuupdatercli_tar_url, "-o", "package.tar.gz")
	err := wb32dfuupdatercli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wb32dfuupdatercli_zip_url := "https://github.com/WestberryTech/wb32-dfu-updater/archive/refs/tags/1.0.0.zip"
	wb32dfuupdatercli_cmd_zip := exec.Command("curl", "-L", wb32dfuupdatercli_zip_url, "-o", "package.zip")
	err = wb32dfuupdatercli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wb32dfuupdatercli_bin_url := "https://github.com/WestberryTech/wb32-dfu-updater/archive/refs/tags/1.0.0.bin"
	wb32dfuupdatercli_cmd_bin := exec.Command("curl", "-L", wb32dfuupdatercli_bin_url, "-o", "binary.bin")
	err = wb32dfuupdatercli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wb32dfuupdatercli_src_url := "https://github.com/WestberryTech/wb32-dfu-updater/archive/refs/tags/1.0.0.src.tar.gz"
	wb32dfuupdatercli_cmd_src := exec.Command("curl", "-L", wb32dfuupdatercli_src_url, "-o", "source.tar.gz")
	err = wb32dfuupdatercli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wb32dfuupdatercli_cmd_direct := exec.Command("./binary")
	err = wb32dfuupdatercli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
