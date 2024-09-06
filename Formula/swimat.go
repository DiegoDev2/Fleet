package main

// Swimat - Command-line tool to help format Swift code
// Homepage: https://github.com/Jintin/Swimat

import (
	"fmt"
	
	"os/exec"
)

func installSwimat() {
	// Método 1: Descargar y extraer .tar.gz
	swimat_tar_url := "https://github.com/Jintin/Swimat/archive/refs/tags/1.7.0.tar.gz"
	swimat_cmd_tar := exec.Command("curl", "-L", swimat_tar_url, "-o", "package.tar.gz")
	err := swimat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swimat_zip_url := "https://github.com/Jintin/Swimat/archive/refs/tags/1.7.0.zip"
	swimat_cmd_zip := exec.Command("curl", "-L", swimat_zip_url, "-o", "package.zip")
	err = swimat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swimat_bin_url := "https://github.com/Jintin/Swimat/archive/refs/tags/1.7.0.bin"
	swimat_cmd_bin := exec.Command("curl", "-L", swimat_bin_url, "-o", "binary.bin")
	err = swimat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swimat_src_url := "https://github.com/Jintin/Swimat/archive/refs/tags/1.7.0.src.tar.gz"
	swimat_cmd_src := exec.Command("curl", "-L", swimat_src_url, "-o", "source.tar.gz")
	err = swimat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swimat_cmd_direct := exec.Command("./binary")
	err = swimat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
