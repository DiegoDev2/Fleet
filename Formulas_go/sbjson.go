package main

// Sbjson - JSON CLI parser & reformatter based on SBJson v5
// Homepage: https://github.com/SBJson/SBJson

import (
	"fmt"
	
	"os/exec"
)

func installSbjson() {
	// Método 1: Descargar y extraer .tar.gz
	sbjson_tar_url := "https://github.com/SBJson/SBJson/archive/refs/tags/v5.0.3.tar.gz"
	sbjson_cmd_tar := exec.Command("curl", "-L", sbjson_tar_url, "-o", "package.tar.gz")
	err := sbjson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sbjson_zip_url := "https://github.com/SBJson/SBJson/archive/refs/tags/v5.0.3.zip"
	sbjson_cmd_zip := exec.Command("curl", "-L", sbjson_zip_url, "-o", "package.zip")
	err = sbjson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sbjson_bin_url := "https://github.com/SBJson/SBJson/archive/refs/tags/v5.0.3.bin"
	sbjson_cmd_bin := exec.Command("curl", "-L", sbjson_bin_url, "-o", "binary.bin")
	err = sbjson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sbjson_src_url := "https://github.com/SBJson/SBJson/archive/refs/tags/v5.0.3.src.tar.gz"
	sbjson_cmd_src := exec.Command("curl", "-L", sbjson_src_url, "-o", "source.tar.gz")
	err = sbjson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sbjson_cmd_direct := exec.Command("./binary")
	err = sbjson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
