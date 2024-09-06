package main

// Pick - Utility to choose one option from a set of choices
// Homepage: https://github.com/mptre/pick

import (
	"fmt"
	
	"os/exec"
)

func installPick() {
	// Método 1: Descargar y extraer .tar.gz
	pick_tar_url := "https://github.com/mptre/pick/releases/download/v4.0.0/pick-4.0.0.tar.gz"
	pick_cmd_tar := exec.Command("curl", "-L", pick_tar_url, "-o", "package.tar.gz")
	err := pick_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pick_zip_url := "https://github.com/mptre/pick/releases/download/v4.0.0/pick-4.0.0.zip"
	pick_cmd_zip := exec.Command("curl", "-L", pick_zip_url, "-o", "package.zip")
	err = pick_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pick_bin_url := "https://github.com/mptre/pick/releases/download/v4.0.0/pick-4.0.0.bin"
	pick_cmd_bin := exec.Command("curl", "-L", pick_bin_url, "-o", "binary.bin")
	err = pick_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pick_src_url := "https://github.com/mptre/pick/releases/download/v4.0.0/pick-4.0.0.src.tar.gz"
	pick_cmd_src := exec.Command("curl", "-L", pick_src_url, "-o", "source.tar.gz")
	err = pick_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pick_cmd_direct := exec.Command("./binary")
	err = pick_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
