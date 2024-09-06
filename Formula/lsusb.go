package main

// Lsusb - List USB devices, just like the Linux lsusb command
// Homepage: https://github.com/jlhonora/lsusb

import (
	"fmt"
	
	"os/exec"
)

func installLsusb() {
	// Método 1: Descargar y extraer .tar.gz
	lsusb_tar_url := "https://github.com/jlhonora/lsusb/releases/download/1.0/lsusb-1.0.tar.gz"
	lsusb_cmd_tar := exec.Command("curl", "-L", lsusb_tar_url, "-o", "package.tar.gz")
	err := lsusb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lsusb_zip_url := "https://github.com/jlhonora/lsusb/releases/download/1.0/lsusb-1.0.zip"
	lsusb_cmd_zip := exec.Command("curl", "-L", lsusb_zip_url, "-o", "package.zip")
	err = lsusb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lsusb_bin_url := "https://github.com/jlhonora/lsusb/releases/download/1.0/lsusb-1.0.bin"
	lsusb_cmd_bin := exec.Command("curl", "-L", lsusb_bin_url, "-o", "binary.bin")
	err = lsusb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lsusb_src_url := "https://github.com/jlhonora/lsusb/releases/download/1.0/lsusb-1.0.src.tar.gz"
	lsusb_cmd_src := exec.Command("curl", "-L", lsusb_src_url, "-o", "source.tar.gz")
	err = lsusb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lsusb_cmd_direct := exec.Command("./binary")
	err = lsusb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
