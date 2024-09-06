package main

// Tinysvm - Support vector machine library for pattern recognition
// Homepage: http://chasen.org/~taku/software/TinySVM/

import (
	"fmt"
	
	"os/exec"
)

func installTinysvm() {
	// Método 1: Descargar y extraer .tar.gz
	tinysvm_tar_url := "http://chasen.org/~taku/software/TinySVM/src/TinySVM-0.09.tar.gz"
	tinysvm_cmd_tar := exec.Command("curl", "-L", tinysvm_tar_url, "-o", "package.tar.gz")
	err := tinysvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tinysvm_zip_url := "http://chasen.org/~taku/software/TinySVM/src/TinySVM-0.09.zip"
	tinysvm_cmd_zip := exec.Command("curl", "-L", tinysvm_zip_url, "-o", "package.zip")
	err = tinysvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tinysvm_bin_url := "http://chasen.org/~taku/software/TinySVM/src/TinySVM-0.09.bin"
	tinysvm_cmd_bin := exec.Command("curl", "-L", tinysvm_bin_url, "-o", "binary.bin")
	err = tinysvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tinysvm_src_url := "http://chasen.org/~taku/software/TinySVM/src/TinySVM-0.09.src.tar.gz"
	tinysvm_cmd_src := exec.Command("curl", "-L", tinysvm_src_url, "-o", "source.tar.gz")
	err = tinysvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tinysvm_cmd_direct := exec.Command("./binary")
	err = tinysvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
