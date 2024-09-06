package main

// ActionsBatch - Time-sharing supercomputer built on GitHub Actions
// Homepage: https://github.com/alexellis/actions-batch

import (
	"fmt"
	
	"os/exec"
)

func installActionsBatch() {
	// Método 1: Descargar y extraer .tar.gz
	actionsbatch_tar_url := "https://github.com/alexellis/actions-batch/archive/refs/tags/v0.0.3.tar.gz"
	actionsbatch_cmd_tar := exec.Command("curl", "-L", actionsbatch_tar_url, "-o", "package.tar.gz")
	err := actionsbatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	actionsbatch_zip_url := "https://github.com/alexellis/actions-batch/archive/refs/tags/v0.0.3.zip"
	actionsbatch_cmd_zip := exec.Command("curl", "-L", actionsbatch_zip_url, "-o", "package.zip")
	err = actionsbatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	actionsbatch_bin_url := "https://github.com/alexellis/actions-batch/archive/refs/tags/v0.0.3.bin"
	actionsbatch_cmd_bin := exec.Command("curl", "-L", actionsbatch_bin_url, "-o", "binary.bin")
	err = actionsbatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	actionsbatch_src_url := "https://github.com/alexellis/actions-batch/archive/refs/tags/v0.0.3.src.tar.gz"
	actionsbatch_cmd_src := exec.Command("curl", "-L", actionsbatch_src_url, "-o", "source.tar.gz")
	err = actionsbatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	actionsbatch_cmd_direct := exec.Command("./binary")
	err = actionsbatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
