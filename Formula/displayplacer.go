package main

// Displayplacer - Utility to configure multi-display resolutions and arrangements
// Homepage: https://github.com/jakehilborn/displayplacer

import (
	"fmt"
	
	"os/exec"
)

func installDisplayplacer() {
	// Método 1: Descargar y extraer .tar.gz
	displayplacer_tar_url := "https://github.com/jakehilborn/displayplacer/archive/refs/tags/v1.4.0.tar.gz"
	displayplacer_cmd_tar := exec.Command("curl", "-L", displayplacer_tar_url, "-o", "package.tar.gz")
	err := displayplacer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	displayplacer_zip_url := "https://github.com/jakehilborn/displayplacer/archive/refs/tags/v1.4.0.zip"
	displayplacer_cmd_zip := exec.Command("curl", "-L", displayplacer_zip_url, "-o", "package.zip")
	err = displayplacer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	displayplacer_bin_url := "https://github.com/jakehilborn/displayplacer/archive/refs/tags/v1.4.0.bin"
	displayplacer_cmd_bin := exec.Command("curl", "-L", displayplacer_bin_url, "-o", "binary.bin")
	err = displayplacer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	displayplacer_src_url := "https://github.com/jakehilborn/displayplacer/archive/refs/tags/v1.4.0.src.tar.gz"
	displayplacer_cmd_src := exec.Command("curl", "-L", displayplacer_src_url, "-o", "source.tar.gz")
	err = displayplacer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	displayplacer_cmd_direct := exec.Command("./binary")
	err = displayplacer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
