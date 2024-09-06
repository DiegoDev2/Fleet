package main

// GoCritic - Opinionated Go source code linter
// Homepage: https://go-critic.com

import (
	"fmt"
	
	"os/exec"
)

func installGoCritic() {
	// Método 1: Descargar y extraer .tar.gz
	gocritic_tar_url := "https://github.com/go-critic/go-critic/archive/refs/tags/v0.11.4.tar.gz"
	gocritic_cmd_tar := exec.Command("curl", "-L", gocritic_tar_url, "-o", "package.tar.gz")
	err := gocritic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gocritic_zip_url := "https://github.com/go-critic/go-critic/archive/refs/tags/v0.11.4.zip"
	gocritic_cmd_zip := exec.Command("curl", "-L", gocritic_zip_url, "-o", "package.zip")
	err = gocritic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gocritic_bin_url := "https://github.com/go-critic/go-critic/archive/refs/tags/v0.11.4.bin"
	gocritic_cmd_bin := exec.Command("curl", "-L", gocritic_bin_url, "-o", "binary.bin")
	err = gocritic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gocritic_src_url := "https://github.com/go-critic/go-critic/archive/refs/tags/v0.11.4.src.tar.gz"
	gocritic_cmd_src := exec.Command("curl", "-L", gocritic_src_url, "-o", "source.tar.gz")
	err = gocritic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gocritic_cmd_direct := exec.Command("./binary")
	err = gocritic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
