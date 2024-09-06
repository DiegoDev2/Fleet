package main

// Wy60 - Wyse 60 compatible terminal emulator
// Homepage: https://code.google.com/archive/p/wy60/

import (
	"fmt"
	
	"os/exec"
)

func installWy60() {
	// Método 1: Descargar y extraer .tar.gz
	wy60_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/wy60/wy60-2.0.9.tar.gz"
	wy60_cmd_tar := exec.Command("curl", "-L", wy60_tar_url, "-o", "package.tar.gz")
	err := wy60_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wy60_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/wy60/wy60-2.0.9.zip"
	wy60_cmd_zip := exec.Command("curl", "-L", wy60_zip_url, "-o", "package.zip")
	err = wy60_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wy60_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/wy60/wy60-2.0.9.bin"
	wy60_cmd_bin := exec.Command("curl", "-L", wy60_bin_url, "-o", "binary.bin")
	err = wy60_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wy60_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/wy60/wy60-2.0.9.src.tar.gz"
	wy60_cmd_src := exec.Command("curl", "-L", wy60_src_url, "-o", "source.tar.gz")
	err = wy60_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wy60_cmd_direct := exec.Command("./binary")
	err = wy60_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
