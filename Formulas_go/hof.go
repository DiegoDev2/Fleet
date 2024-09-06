package main

// Hof - Flexible data modeling & code generation system
// Homepage: https://hofstadter.io/

import (
	"fmt"
	
	"os/exec"
)

func installHof() {
	// Método 1: Descargar y extraer .tar.gz
	hof_tar_url := "https://github.com/hofstadter-io/hof/archive/refs/tags/v0.6.9.tar.gz"
	hof_cmd_tar := exec.Command("curl", "-L", hof_tar_url, "-o", "package.tar.gz")
	err := hof_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hof_zip_url := "https://github.com/hofstadter-io/hof/archive/refs/tags/v0.6.9.zip"
	hof_cmd_zip := exec.Command("curl", "-L", hof_zip_url, "-o", "package.zip")
	err = hof_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hof_bin_url := "https://github.com/hofstadter-io/hof/archive/refs/tags/v0.6.9.bin"
	hof_cmd_bin := exec.Command("curl", "-L", hof_bin_url, "-o", "binary.bin")
	err = hof_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hof_src_url := "https://github.com/hofstadter-io/hof/archive/refs/tags/v0.6.9.src.tar.gz"
	hof_cmd_src := exec.Command("curl", "-L", hof_src_url, "-o", "source.tar.gz")
	err = hof_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hof_cmd_direct := exec.Command("./binary")
	err = hof_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
exec.Command("latte", "install", "go@1.22")
}
