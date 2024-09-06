package main

// Liblinear - Library for large linear classification
// Homepage: https://www.csie.ntu.edu.tw/~cjlin/liblinear/

import (
	"fmt"
	
	"os/exec"
)

func installLiblinear() {
	// Método 1: Descargar y extraer .tar.gz
	liblinear_tar_url := "https://www.csie.ntu.edu.tw/~cjlin/liblinear/oldfiles/liblinear-2.47.tar.gz"
	liblinear_cmd_tar := exec.Command("curl", "-L", liblinear_tar_url, "-o", "package.tar.gz")
	err := liblinear_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblinear_zip_url := "https://www.csie.ntu.edu.tw/~cjlin/liblinear/oldfiles/liblinear-2.47.zip"
	liblinear_cmd_zip := exec.Command("curl", "-L", liblinear_zip_url, "-o", "package.zip")
	err = liblinear_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblinear_bin_url := "https://www.csie.ntu.edu.tw/~cjlin/liblinear/oldfiles/liblinear-2.47.bin"
	liblinear_cmd_bin := exec.Command("curl", "-L", liblinear_bin_url, "-o", "binary.bin")
	err = liblinear_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblinear_src_url := "https://www.csie.ntu.edu.tw/~cjlin/liblinear/oldfiles/liblinear-2.47.src.tar.gz"
	liblinear_cmd_src := exec.Command("curl", "-L", liblinear_src_url, "-o", "source.tar.gz")
	err = liblinear_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblinear_cmd_direct := exec.Command("./binary")
	err = liblinear_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
