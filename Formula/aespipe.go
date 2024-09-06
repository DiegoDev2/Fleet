package main

// Aespipe - AES encryption or decryption for pipes
// Homepage: https://loop-aes.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAespipe() {
	// Método 1: Descargar y extraer .tar.gz
	aespipe_tar_url := "https://loop-aes.sourceforge.net/aespipe/aespipe-v2.4h.tar.bz2"
	aespipe_cmd_tar := exec.Command("curl", "-L", aespipe_tar_url, "-o", "package.tar.gz")
	err := aespipe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aespipe_zip_url := "https://loop-aes.sourceforge.net/aespipe/aespipe-v2.4h.tar.bz2"
	aespipe_cmd_zip := exec.Command("curl", "-L", aespipe_zip_url, "-o", "package.zip")
	err = aespipe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aespipe_bin_url := "https://loop-aes.sourceforge.net/aespipe/aespipe-v2.4h.tar.bz2"
	aespipe_cmd_bin := exec.Command("curl", "-L", aespipe_bin_url, "-o", "binary.bin")
	err = aespipe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aespipe_src_url := "https://loop-aes.sourceforge.net/aespipe/aespipe-v2.4h.tar.bz2"
	aespipe_cmd_src := exec.Command("curl", "-L", aespipe_src_url, "-o", "source.tar.gz")
	err = aespipe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aespipe_cmd_direct := exec.Command("./binary")
	err = aespipe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
