package main

// Mogenerator - Generate Objective-C & Swift classes from your Core Data model
// Homepage: https://rentzsch.github.io/mogenerator/

import (
	"fmt"
	
	"os/exec"
)

func installMogenerator() {
	// Método 1: Descargar y extraer .tar.gz
	mogenerator_tar_url := "https://github.com/rentzsch/mogenerator/archive/refs/tags/1.32.tar.gz"
	mogenerator_cmd_tar := exec.Command("curl", "-L", mogenerator_tar_url, "-o", "package.tar.gz")
	err := mogenerator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mogenerator_zip_url := "https://github.com/rentzsch/mogenerator/archive/refs/tags/1.32.zip"
	mogenerator_cmd_zip := exec.Command("curl", "-L", mogenerator_zip_url, "-o", "package.zip")
	err = mogenerator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mogenerator_bin_url := "https://github.com/rentzsch/mogenerator/archive/refs/tags/1.32.bin"
	mogenerator_cmd_bin := exec.Command("curl", "-L", mogenerator_bin_url, "-o", "binary.bin")
	err = mogenerator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mogenerator_src_url := "https://github.com/rentzsch/mogenerator/archive/refs/tags/1.32.src.tar.gz"
	mogenerator_cmd_src := exec.Command("curl", "-L", mogenerator_src_url, "-o", "source.tar.gz")
	err = mogenerator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mogenerator_cmd_direct := exec.Command("./binary")
	err = mogenerator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
