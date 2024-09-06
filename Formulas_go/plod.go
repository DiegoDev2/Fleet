package main

// Plod - Keep an online journal of what you're working on
// Homepage: https://deer-run.com/users/hal/

import (
	"fmt"
	
	"os/exec"
)

func installPlod() {
	// Método 1: Descargar y extraer .tar.gz
	plod_tar_url := "https://deer-run.com/~hal/plod/plod.shar"
	plod_cmd_tar := exec.Command("curl", "-L", plod_tar_url, "-o", "package.tar.gz")
	err := plod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plod_zip_url := "https://deer-run.com/~hal/plod/plod.shar"
	plod_cmd_zip := exec.Command("curl", "-L", plod_zip_url, "-o", "package.zip")
	err = plod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plod_bin_url := "https://deer-run.com/~hal/plod/plod.shar"
	plod_cmd_bin := exec.Command("curl", "-L", plod_bin_url, "-o", "binary.bin")
	err = plod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plod_src_url := "https://deer-run.com/~hal/plod/plod.shar"
	plod_cmd_src := exec.Command("curl", "-L", plod_src_url, "-o", "source.tar.gz")
	err = plod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plod_cmd_direct := exec.Command("./binary")
	err = plod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
