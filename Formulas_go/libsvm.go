package main

// Libsvm - Library for support vector machines
// Homepage: https://www.csie.ntu.edu.tw/~cjlin/libsvm/

import (
	"fmt"
	
	"os/exec"
)

func installLibsvm() {
	// Método 1: Descargar y extraer .tar.gz
	libsvm_tar_url := "https://www.csie.ntu.edu.tw/~cjlin/libsvm/libsvm-3.35.tar.gz"
	libsvm_cmd_tar := exec.Command("curl", "-L", libsvm_tar_url, "-o", "package.tar.gz")
	err := libsvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsvm_zip_url := "https://www.csie.ntu.edu.tw/~cjlin/libsvm/libsvm-3.35.zip"
	libsvm_cmd_zip := exec.Command("curl", "-L", libsvm_zip_url, "-o", "package.zip")
	err = libsvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsvm_bin_url := "https://www.csie.ntu.edu.tw/~cjlin/libsvm/libsvm-3.35.bin"
	libsvm_cmd_bin := exec.Command("curl", "-L", libsvm_bin_url, "-o", "binary.bin")
	err = libsvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsvm_src_url := "https://www.csie.ntu.edu.tw/~cjlin/libsvm/libsvm-3.35.src.tar.gz"
	libsvm_cmd_src := exec.Command("curl", "-L", libsvm_src_url, "-o", "source.tar.gz")
	err = libsvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsvm_cmd_direct := exec.Command("./binary")
	err = libsvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
