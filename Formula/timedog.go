package main

// Timedog - Lists files that were saved by a backup of the macOS Time Machine
// Homepage: https://github.com/nlfiedler/timedog

import (
	"fmt"
	
	"os/exec"
)

func installTimedog() {
	// Método 1: Descargar y extraer .tar.gz
	timedog_tar_url := "https://github.com/nlfiedler/timedog/archive/refs/tags/v1.4.tar.gz"
	timedog_cmd_tar := exec.Command("curl", "-L", timedog_tar_url, "-o", "package.tar.gz")
	err := timedog_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	timedog_zip_url := "https://github.com/nlfiedler/timedog/archive/refs/tags/v1.4.zip"
	timedog_cmd_zip := exec.Command("curl", "-L", timedog_zip_url, "-o", "package.zip")
	err = timedog_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	timedog_bin_url := "https://github.com/nlfiedler/timedog/archive/refs/tags/v1.4.bin"
	timedog_cmd_bin := exec.Command("curl", "-L", timedog_bin_url, "-o", "binary.bin")
	err = timedog_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	timedog_src_url := "https://github.com/nlfiedler/timedog/archive/refs/tags/v1.4.src.tar.gz"
	timedog_cmd_src := exec.Command("curl", "-L", timedog_src_url, "-o", "source.tar.gz")
	err = timedog_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	timedog_cmd_direct := exec.Command("./binary")
	err = timedog_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
