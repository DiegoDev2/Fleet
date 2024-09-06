package main

// Bas55 - Minimal BASIC programming language interpreter as defined by ECMA-55
// Homepage: https://jorgicor.niobe.org/bas55/

import (
	"fmt"
	
	"os/exec"
)

func installBas55() {
	// Método 1: Descargar y extraer .tar.gz
	bas55_tar_url := "https://jorgicor.niobe.org/bas55/bas55-2.0.tar.gz"
	bas55_cmd_tar := exec.Command("curl", "-L", bas55_tar_url, "-o", "package.tar.gz")
	err := bas55_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bas55_zip_url := "https://jorgicor.niobe.org/bas55/bas55-2.0.zip"
	bas55_cmd_zip := exec.Command("curl", "-L", bas55_zip_url, "-o", "package.zip")
	err = bas55_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bas55_bin_url := "https://jorgicor.niobe.org/bas55/bas55-2.0.bin"
	bas55_cmd_bin := exec.Command("curl", "-L", bas55_bin_url, "-o", "binary.bin")
	err = bas55_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bas55_src_url := "https://jorgicor.niobe.org/bas55/bas55-2.0.src.tar.gz"
	bas55_cmd_src := exec.Command("curl", "-L", bas55_src_url, "-o", "source.tar.gz")
	err = bas55_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bas55_cmd_direct := exec.Command("./binary")
	err = bas55_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
