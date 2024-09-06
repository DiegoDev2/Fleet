package main

// Chicken - Compiler for the Scheme programming language
// Homepage: https://www.call-cc.org/

import (
	"fmt"
	
	"os/exec"
)

func installChicken() {
	// Método 1: Descargar y extraer .tar.gz
	chicken_tar_url := "https://code.call-cc.org/releases/5.4.0/chicken-5.4.0.tar.gz"
	chicken_cmd_tar := exec.Command("curl", "-L", chicken_tar_url, "-o", "package.tar.gz")
	err := chicken_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chicken_zip_url := "https://code.call-cc.org/releases/5.4.0/chicken-5.4.0.zip"
	chicken_cmd_zip := exec.Command("curl", "-L", chicken_zip_url, "-o", "package.zip")
	err = chicken_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chicken_bin_url := "https://code.call-cc.org/releases/5.4.0/chicken-5.4.0.bin"
	chicken_cmd_bin := exec.Command("curl", "-L", chicken_bin_url, "-o", "binary.bin")
	err = chicken_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chicken_src_url := "https://code.call-cc.org/releases/5.4.0/chicken-5.4.0.src.tar.gz"
	chicken_cmd_src := exec.Command("curl", "-L", chicken_src_url, "-o", "source.tar.gz")
	err = chicken_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chicken_cmd_direct := exec.Command("./binary")
	err = chicken_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
