package main

// Webarchiver - Allows you to create Safari .webarchive files
// Homepage: https://github.com/newzealandpaul/webarchiver

import (
	"fmt"
	
	"os/exec"
)

func installWebarchiver() {
	// Método 1: Descargar y extraer .tar.gz
	webarchiver_tar_url := "https://github.com/newzealandpaul/webarchiver/archive/refs/tags/0.12.tar.gz"
	webarchiver_cmd_tar := exec.Command("curl", "-L", webarchiver_tar_url, "-o", "package.tar.gz")
	err := webarchiver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webarchiver_zip_url := "https://github.com/newzealandpaul/webarchiver/archive/refs/tags/0.12.zip"
	webarchiver_cmd_zip := exec.Command("curl", "-L", webarchiver_zip_url, "-o", "package.zip")
	err = webarchiver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webarchiver_bin_url := "https://github.com/newzealandpaul/webarchiver/archive/refs/tags/0.12.bin"
	webarchiver_cmd_bin := exec.Command("curl", "-L", webarchiver_bin_url, "-o", "binary.bin")
	err = webarchiver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webarchiver_src_url := "https://github.com/newzealandpaul/webarchiver/archive/refs/tags/0.12.src.tar.gz"
	webarchiver_cmd_src := exec.Command("curl", "-L", webarchiver_src_url, "-o", "source.tar.gz")
	err = webarchiver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webarchiver_cmd_direct := exec.Command("./binary")
	err = webarchiver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
