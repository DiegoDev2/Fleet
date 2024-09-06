package main

// Libb2 - Secure hashing function
// Homepage: https://blake2.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibb2() {
	// Método 1: Descargar y extraer .tar.gz
	libb2_tar_url := "https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.tar.gz"
	libb2_cmd_tar := exec.Command("curl", "-L", libb2_tar_url, "-o", "package.tar.gz")
	err := libb2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libb2_zip_url := "https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.zip"
	libb2_cmd_zip := exec.Command("curl", "-L", libb2_zip_url, "-o", "package.zip")
	err = libb2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libb2_bin_url := "https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.bin"
	libb2_cmd_bin := exec.Command("curl", "-L", libb2_bin_url, "-o", "binary.bin")
	err = libb2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libb2_src_url := "https://github.com/BLAKE2/libb2/releases/download/v0.98.1/libb2-0.98.1.src.tar.gz"
	libb2_cmd_src := exec.Command("curl", "-L", libb2_src_url, "-o", "source.tar.gz")
	err = libb2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libb2_cmd_direct := exec.Command("./binary")
	err = libb2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
