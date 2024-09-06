package main

// Lwtools - Cross-development tools for Motorola 6809 and Hitachi 6309
// Homepage: http://www.lwtools.ca/

import (
	"fmt"
	
	"os/exec"
)

func installLwtools() {
	// Método 1: Descargar y extraer .tar.gz
	lwtools_tar_url := "http://www.lwtools.ca/releases/lwtools/lwtools-4.23.tar.gz"
	lwtools_cmd_tar := exec.Command("curl", "-L", lwtools_tar_url, "-o", "package.tar.gz")
	err := lwtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lwtools_zip_url := "http://www.lwtools.ca/releases/lwtools/lwtools-4.23.zip"
	lwtools_cmd_zip := exec.Command("curl", "-L", lwtools_zip_url, "-o", "package.zip")
	err = lwtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lwtools_bin_url := "http://www.lwtools.ca/releases/lwtools/lwtools-4.23.bin"
	lwtools_cmd_bin := exec.Command("curl", "-L", lwtools_bin_url, "-o", "binary.bin")
	err = lwtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lwtools_src_url := "http://www.lwtools.ca/releases/lwtools/lwtools-4.23.src.tar.gz"
	lwtools_cmd_src := exec.Command("curl", "-L", lwtools_src_url, "-o", "source.tar.gz")
	err = lwtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lwtools_cmd_direct := exec.Command("./binary")
	err = lwtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
