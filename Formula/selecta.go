package main

// Selecta - Fuzzy text selector for files and anything else you need to select
// Homepage: https://github.com/garybernhardt/selecta

import (
	"fmt"
	
	"os/exec"
)

func installSelecta() {
	// Método 1: Descargar y extraer .tar.gz
	selecta_tar_url := "https://github.com/garybernhardt/selecta/archive/refs/tags/v0.0.7.tar.gz"
	selecta_cmd_tar := exec.Command("curl", "-L", selecta_tar_url, "-o", "package.tar.gz")
	err := selecta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	selecta_zip_url := "https://github.com/garybernhardt/selecta/archive/refs/tags/v0.0.7.zip"
	selecta_cmd_zip := exec.Command("curl", "-L", selecta_zip_url, "-o", "package.zip")
	err = selecta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	selecta_bin_url := "https://github.com/garybernhardt/selecta/archive/refs/tags/v0.0.7.bin"
	selecta_cmd_bin := exec.Command("curl", "-L", selecta_bin_url, "-o", "binary.bin")
	err = selecta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	selecta_src_url := "https://github.com/garybernhardt/selecta/archive/refs/tags/v0.0.7.src.tar.gz"
	selecta_cmd_src := exec.Command("curl", "-L", selecta_src_url, "-o", "source.tar.gz")
	err = selecta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	selecta_cmd_direct := exec.Command("./binary")
	err = selecta_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
