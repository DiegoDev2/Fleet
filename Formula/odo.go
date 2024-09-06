package main

// Odo - Atomic odometer for the command-line
// Homepage: https://github.com/atomicobject/odo

import (
	"fmt"
	
	"os/exec"
)

func installOdo() {
	// Método 1: Descargar y extraer .tar.gz
	odo_tar_url := "https://github.com/atomicobject/odo/archive/refs/tags/v0.2.2.tar.gz"
	odo_cmd_tar := exec.Command("curl", "-L", odo_tar_url, "-o", "package.tar.gz")
	err := odo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	odo_zip_url := "https://github.com/atomicobject/odo/archive/refs/tags/v0.2.2.zip"
	odo_cmd_zip := exec.Command("curl", "-L", odo_zip_url, "-o", "package.zip")
	err = odo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	odo_bin_url := "https://github.com/atomicobject/odo/archive/refs/tags/v0.2.2.bin"
	odo_cmd_bin := exec.Command("curl", "-L", odo_bin_url, "-o", "binary.bin")
	err = odo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	odo_src_url := "https://github.com/atomicobject/odo/archive/refs/tags/v0.2.2.src.tar.gz"
	odo_cmd_src := exec.Command("curl", "-L", odo_src_url, "-o", "source.tar.gz")
	err = odo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	odo_cmd_direct := exec.Command("./binary")
	err = odo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
