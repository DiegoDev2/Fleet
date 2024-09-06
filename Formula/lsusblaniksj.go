package main

// LsusbLaniksj - List USB devices, just like the Linux lsusb command
// Homepage: https://github.com/LanikSJ/lsusb

import (
	"fmt"
	
	"os/exec"
)

func installLsusbLaniksj() {
	// Método 1: Descargar y extraer .tar.gz
	lsusblaniksj_tar_url := "https://github.com/LanikSJ/lsusb/archive/refs/tags/1.1.4.tar.gz"
	lsusblaniksj_cmd_tar := exec.Command("curl", "-L", lsusblaniksj_tar_url, "-o", "package.tar.gz")
	err := lsusblaniksj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lsusblaniksj_zip_url := "https://github.com/LanikSJ/lsusb/archive/refs/tags/1.1.4.zip"
	lsusblaniksj_cmd_zip := exec.Command("curl", "-L", lsusblaniksj_zip_url, "-o", "package.zip")
	err = lsusblaniksj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lsusblaniksj_bin_url := "https://github.com/LanikSJ/lsusb/archive/refs/tags/1.1.4.bin"
	lsusblaniksj_cmd_bin := exec.Command("curl", "-L", lsusblaniksj_bin_url, "-o", "binary.bin")
	err = lsusblaniksj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lsusblaniksj_src_url := "https://github.com/LanikSJ/lsusb/archive/refs/tags/1.1.4.src.tar.gz"
	lsusblaniksj_cmd_src := exec.Command("curl", "-L", lsusblaniksj_src_url, "-o", "source.tar.gz")
	err = lsusblaniksj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lsusblaniksj_cmd_direct := exec.Command("./binary")
	err = lsusblaniksj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
