package main

// Nettoe - Tic Tac Toe-like game for the console
// Homepage: https://nettoe.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installNettoe() {
	// Método 1: Descargar y extraer .tar.gz
	nettoe_tar_url := "https://downloads.sourceforge.net/project/nettoe/nettoe/1.5.1/nettoe-1.5.1.tar.gz"
	nettoe_cmd_tar := exec.Command("curl", "-L", nettoe_tar_url, "-o", "package.tar.gz")
	err := nettoe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nettoe_zip_url := "https://downloads.sourceforge.net/project/nettoe/nettoe/1.5.1/nettoe-1.5.1.zip"
	nettoe_cmd_zip := exec.Command("curl", "-L", nettoe_zip_url, "-o", "package.zip")
	err = nettoe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nettoe_bin_url := "https://downloads.sourceforge.net/project/nettoe/nettoe/1.5.1/nettoe-1.5.1.bin"
	nettoe_cmd_bin := exec.Command("curl", "-L", nettoe_bin_url, "-o", "binary.bin")
	err = nettoe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nettoe_src_url := "https://downloads.sourceforge.net/project/nettoe/nettoe/1.5.1/nettoe-1.5.1.src.tar.gz"
	nettoe_cmd_src := exec.Command("curl", "-L", nettoe_src_url, "-o", "source.tar.gz")
	err = nettoe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nettoe_cmd_direct := exec.Command("./binary")
	err = nettoe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
