package main

// Picocom - Minimal dumb-terminal emulation program
// Homepage: https://github.com/npat-efault/picocom

import (
	"fmt"
	
	"os/exec"
)

func installPicocom() {
	// Método 1: Descargar y extraer .tar.gz
	picocom_tar_url := "https://github.com/npat-efault/picocom/archive/refs/tags/3.1.tar.gz"
	picocom_cmd_tar := exec.Command("curl", "-L", picocom_tar_url, "-o", "package.tar.gz")
	err := picocom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	picocom_zip_url := "https://github.com/npat-efault/picocom/archive/refs/tags/3.1.zip"
	picocom_cmd_zip := exec.Command("curl", "-L", picocom_zip_url, "-o", "package.zip")
	err = picocom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	picocom_bin_url := "https://github.com/npat-efault/picocom/archive/refs/tags/3.1.bin"
	picocom_cmd_bin := exec.Command("curl", "-L", picocom_bin_url, "-o", "binary.bin")
	err = picocom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	picocom_src_url := "https://github.com/npat-efault/picocom/archive/refs/tags/3.1.src.tar.gz"
	picocom_cmd_src := exec.Command("curl", "-L", picocom_src_url, "-o", "source.tar.gz")
	err = picocom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	picocom_cmd_direct := exec.Command("./binary")
	err = picocom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
