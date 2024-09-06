package main

// Pict - Pairwise Independent Combinatorial Tool
// Homepage: https://github.com/Microsoft/pict/

import (
	"fmt"
	
	"os/exec"
)

func installPict() {
	// Método 1: Descargar y extraer .tar.gz
	pict_tar_url := "https://github.com/Microsoft/pict/archive/refs/tags/v3.7.4.tar.gz"
	pict_cmd_tar := exec.Command("curl", "-L", pict_tar_url, "-o", "package.tar.gz")
	err := pict_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pict_zip_url := "https://github.com/Microsoft/pict/archive/refs/tags/v3.7.4.zip"
	pict_cmd_zip := exec.Command("curl", "-L", pict_zip_url, "-o", "package.zip")
	err = pict_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pict_bin_url := "https://github.com/Microsoft/pict/archive/refs/tags/v3.7.4.bin"
	pict_cmd_bin := exec.Command("curl", "-L", pict_bin_url, "-o", "binary.bin")
	err = pict_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pict_src_url := "https://github.com/Microsoft/pict/archive/refs/tags/v3.7.4.src.tar.gz"
	pict_cmd_src := exec.Command("curl", "-L", pict_src_url, "-o", "source.tar.gz")
	err = pict_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pict_cmd_direct := exec.Command("./binary")
	err = pict_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
