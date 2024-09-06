package main

// Eatmemory - Simple program to allocate memory from the command-line
// Homepage: https://github.com/julman99/eatmemory

import (
	"fmt"
	
	"os/exec"
)

func installEatmemory() {
	// Método 1: Descargar y extraer .tar.gz
	eatmemory_tar_url := "https://github.com/julman99/eatmemory/archive/refs/tags/v0.1.10.tar.gz"
	eatmemory_cmd_tar := exec.Command("curl", "-L", eatmemory_tar_url, "-o", "package.tar.gz")
	err := eatmemory_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eatmemory_zip_url := "https://github.com/julman99/eatmemory/archive/refs/tags/v0.1.10.zip"
	eatmemory_cmd_zip := exec.Command("curl", "-L", eatmemory_zip_url, "-o", "package.zip")
	err = eatmemory_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eatmemory_bin_url := "https://github.com/julman99/eatmemory/archive/refs/tags/v0.1.10.bin"
	eatmemory_cmd_bin := exec.Command("curl", "-L", eatmemory_bin_url, "-o", "binary.bin")
	err = eatmemory_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eatmemory_src_url := "https://github.com/julman99/eatmemory/archive/refs/tags/v0.1.10.src.tar.gz"
	eatmemory_cmd_src := exec.Command("curl", "-L", eatmemory_src_url, "-o", "source.tar.gz")
	err = eatmemory_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eatmemory_cmd_direct := exec.Command("./binary")
	err = eatmemory_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
