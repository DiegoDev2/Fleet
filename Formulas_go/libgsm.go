package main

// Libgsm - Lossy speech compression library
// Homepage: https://www.quut.com/gsm/

import (
	"fmt"
	
	"os/exec"
)

func installLibgsm() {
	// Método 1: Descargar y extraer .tar.gz
	libgsm_tar_url := "https://www.quut.com/gsm/gsm-1.0.22.tar.gz"
	libgsm_cmd_tar := exec.Command("curl", "-L", libgsm_tar_url, "-o", "package.tar.gz")
	err := libgsm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgsm_zip_url := "https://www.quut.com/gsm/gsm-1.0.22.zip"
	libgsm_cmd_zip := exec.Command("curl", "-L", libgsm_zip_url, "-o", "package.zip")
	err = libgsm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgsm_bin_url := "https://www.quut.com/gsm/gsm-1.0.22.bin"
	libgsm_cmd_bin := exec.Command("curl", "-L", libgsm_bin_url, "-o", "binary.bin")
	err = libgsm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgsm_src_url := "https://www.quut.com/gsm/gsm-1.0.22.src.tar.gz"
	libgsm_cmd_src := exec.Command("curl", "-L", libgsm_src_url, "-o", "source.tar.gz")
	err = libgsm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgsm_cmd_direct := exec.Command("./binary")
	err = libgsm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
