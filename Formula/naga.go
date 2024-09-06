package main

// Naga - Terminal implementation of the Snake game
// Homepage: https://github.com/anayjoshi/naga/

import (
	"fmt"
	
	"os/exec"
)

func installNaga() {
	// Método 1: Descargar y extraer .tar.gz
	naga_tar_url := "https://github.com/anayjoshi/naga/archive/refs/tags/naga-v1.0.tar.gz"
	naga_cmd_tar := exec.Command("curl", "-L", naga_tar_url, "-o", "package.tar.gz")
	err := naga_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	naga_zip_url := "https://github.com/anayjoshi/naga/archive/refs/tags/naga-v1.0.zip"
	naga_cmd_zip := exec.Command("curl", "-L", naga_zip_url, "-o", "package.zip")
	err = naga_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	naga_bin_url := "https://github.com/anayjoshi/naga/archive/refs/tags/naga-v1.0.bin"
	naga_cmd_bin := exec.Command("curl", "-L", naga_bin_url, "-o", "binary.bin")
	err = naga_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	naga_src_url := "https://github.com/anayjoshi/naga/archive/refs/tags/naga-v1.0.src.tar.gz"
	naga_cmd_src := exec.Command("curl", "-L", naga_src_url, "-o", "source.tar.gz")
	err = naga_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	naga_cmd_direct := exec.Command("./binary")
	err = naga_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
