package main

// Sourcery - Meta-programming for Swift, stop writing boilerplate code
// Homepage: https://github.com/krzysztofzablocki/Sourcery

import (
	"fmt"
	
	"os/exec"
)

func installSourcery() {
	// Método 1: Descargar y extraer .tar.gz
	sourcery_tar_url := "https://github.com/krzysztofzablocki/Sourcery/archive/refs/tags/2.2.5.tar.gz"
	sourcery_cmd_tar := exec.Command("curl", "-L", sourcery_tar_url, "-o", "package.tar.gz")
	err := sourcery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sourcery_zip_url := "https://github.com/krzysztofzablocki/Sourcery/archive/refs/tags/2.2.5.zip"
	sourcery_cmd_zip := exec.Command("curl", "-L", sourcery_zip_url, "-o", "package.zip")
	err = sourcery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sourcery_bin_url := "https://github.com/krzysztofzablocki/Sourcery/archive/refs/tags/2.2.5.bin"
	sourcery_cmd_bin := exec.Command("curl", "-L", sourcery_bin_url, "-o", "binary.bin")
	err = sourcery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sourcery_src_url := "https://github.com/krzysztofzablocki/Sourcery/archive/refs/tags/2.2.5.src.tar.gz"
	sourcery_cmd_src := exec.Command("curl", "-L", sourcery_src_url, "-o", "source.tar.gz")
	err = sourcery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sourcery_cmd_direct := exec.Command("./binary")
	err = sourcery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
