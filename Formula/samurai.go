package main

// Samurai - Ninja-compatible build tool written in C
// Homepage: https://github.com/michaelforney/samurai

import (
	"fmt"
	
	"os/exec"
)

func installSamurai() {
	// Método 1: Descargar y extraer .tar.gz
	samurai_tar_url := "https://github.com/michaelforney/samurai/releases/download/1.2/samurai-1.2.tar.gz"
	samurai_cmd_tar := exec.Command("curl", "-L", samurai_tar_url, "-o", "package.tar.gz")
	err := samurai_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	samurai_zip_url := "https://github.com/michaelforney/samurai/releases/download/1.2/samurai-1.2.zip"
	samurai_cmd_zip := exec.Command("curl", "-L", samurai_zip_url, "-o", "package.zip")
	err = samurai_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	samurai_bin_url := "https://github.com/michaelforney/samurai/releases/download/1.2/samurai-1.2.bin"
	samurai_cmd_bin := exec.Command("curl", "-L", samurai_bin_url, "-o", "binary.bin")
	err = samurai_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	samurai_src_url := "https://github.com/michaelforney/samurai/releases/download/1.2/samurai-1.2.src.tar.gz"
	samurai_cmd_src := exec.Command("curl", "-L", samurai_src_url, "-o", "source.tar.gz")
	err = samurai_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	samurai_cmd_direct := exec.Command("./binary")
	err = samurai_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
