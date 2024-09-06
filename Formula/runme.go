package main

// Runme - Execute commands inside your runbooks, docs, and READMEs
// Homepage: https://runme.dev/

import (
	"fmt"
	
	"os/exec"
)

func installRunme() {
	// Método 1: Descargar y extraer .tar.gz
	runme_tar_url := "https://github.com/stateful/runme/archive/refs/tags/v3.7.0.tar.gz"
	runme_cmd_tar := exec.Command("curl", "-L", runme_tar_url, "-o", "package.tar.gz")
	err := runme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	runme_zip_url := "https://github.com/stateful/runme/archive/refs/tags/v3.7.0.zip"
	runme_cmd_zip := exec.Command("curl", "-L", runme_zip_url, "-o", "package.zip")
	err = runme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	runme_bin_url := "https://github.com/stateful/runme/archive/refs/tags/v3.7.0.bin"
	runme_cmd_bin := exec.Command("curl", "-L", runme_bin_url, "-o", "binary.bin")
	err = runme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	runme_src_url := "https://github.com/stateful/runme/archive/refs/tags/v3.7.0.src.tar.gz"
	runme_cmd_src := exec.Command("curl", "-L", runme_src_url, "-o", "source.tar.gz")
	err = runme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	runme_cmd_direct := exec.Command("./binary")
	err = runme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
