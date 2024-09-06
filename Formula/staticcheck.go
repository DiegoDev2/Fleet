package main

// Staticcheck - State of the art linter for the Go programming language
// Homepage: https://staticcheck.io/

import (
	"fmt"
	
	"os/exec"
)

func installStaticcheck() {
	// Método 1: Descargar y extraer .tar.gz
	staticcheck_tar_url := "https://github.com/dominikh/go-tools/archive/refs/tags/2024.1.1.tar.gz"
	staticcheck_cmd_tar := exec.Command("curl", "-L", staticcheck_tar_url, "-o", "package.tar.gz")
	err := staticcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	staticcheck_zip_url := "https://github.com/dominikh/go-tools/archive/refs/tags/2024.1.1.zip"
	staticcheck_cmd_zip := exec.Command("curl", "-L", staticcheck_zip_url, "-o", "package.zip")
	err = staticcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	staticcheck_bin_url := "https://github.com/dominikh/go-tools/archive/refs/tags/2024.1.1.bin"
	staticcheck_cmd_bin := exec.Command("curl", "-L", staticcheck_bin_url, "-o", "binary.bin")
	err = staticcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	staticcheck_src_url := "https://github.com/dominikh/go-tools/archive/refs/tags/2024.1.1.src.tar.gz"
	staticcheck_cmd_src := exec.Command("curl", "-L", staticcheck_src_url, "-o", "source.tar.gz")
	err = staticcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	staticcheck_cmd_direct := exec.Command("./binary")
	err = staticcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
