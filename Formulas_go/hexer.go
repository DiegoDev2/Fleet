package main

// Hexer - Hex editor for the terminal with vi-like interface
// Homepage: https://devel.ringlet.net/editors/hexer/

import (
	"fmt"
	
	"os/exec"
)

func installHexer() {
	// Método 1: Descargar y extraer .tar.gz
	hexer_tar_url := "https://devel.ringlet.net/files/editors/hexer/hexer-1.0.6.tar.gz"
	hexer_cmd_tar := exec.Command("curl", "-L", hexer_tar_url, "-o", "package.tar.gz")
	err := hexer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hexer_zip_url := "https://devel.ringlet.net/files/editors/hexer/hexer-1.0.6.zip"
	hexer_cmd_zip := exec.Command("curl", "-L", hexer_zip_url, "-o", "package.zip")
	err = hexer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hexer_bin_url := "https://devel.ringlet.net/files/editors/hexer/hexer-1.0.6.bin"
	hexer_cmd_bin := exec.Command("curl", "-L", hexer_bin_url, "-o", "binary.bin")
	err = hexer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hexer_src_url := "https://devel.ringlet.net/files/editors/hexer/hexer-1.0.6.src.tar.gz"
	hexer_cmd_src := exec.Command("curl", "-L", hexer_src_url, "-o", "source.tar.gz")
	err = hexer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hexer_cmd_direct := exec.Command("./binary")
	err = hexer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
