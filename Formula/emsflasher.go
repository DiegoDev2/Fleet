package main

// EmsFlasher - Software for flashing the EMS Gameboy USB cart
// Homepage: https://lacklustre.net/projects/ems-flasher/

import (
	"fmt"
	
	"os/exec"
)

func installEmsFlasher() {
	// Método 1: Descargar y extraer .tar.gz
	emsflasher_tar_url := "https://lacklustre.net/projects/ems-flasher/ems-flasher-0.03.tgz"
	emsflasher_cmd_tar := exec.Command("curl", "-L", emsflasher_tar_url, "-o", "package.tar.gz")
	err := emsflasher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emsflasher_zip_url := "https://lacklustre.net/projects/ems-flasher/ems-flasher-0.03.tgz"
	emsflasher_cmd_zip := exec.Command("curl", "-L", emsflasher_zip_url, "-o", "package.zip")
	err = emsflasher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emsflasher_bin_url := "https://lacklustre.net/projects/ems-flasher/ems-flasher-0.03.tgz"
	emsflasher_cmd_bin := exec.Command("curl", "-L", emsflasher_bin_url, "-o", "binary.bin")
	err = emsflasher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emsflasher_src_url := "https://lacklustre.net/projects/ems-flasher/ems-flasher-0.03.tgz"
	emsflasher_cmd_src := exec.Command("curl", "-L", emsflasher_src_url, "-o", "source.tar.gz")
	err = emsflasher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emsflasher_cmd_direct := exec.Command("./binary")
	err = emsflasher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
