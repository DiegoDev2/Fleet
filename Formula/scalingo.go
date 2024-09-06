package main

// Scalingo - CLI for working with Scalingo's PaaS
// Homepage: https://doc.scalingo.com/cli

import (
	"fmt"
	
	"os/exec"
)

func installScalingo() {
	// Método 1: Descargar y extraer .tar.gz
	scalingo_tar_url := "https://github.com/Scalingo/cli/archive/refs/tags/1.33.0.tar.gz"
	scalingo_cmd_tar := exec.Command("curl", "-L", scalingo_tar_url, "-o", "package.tar.gz")
	err := scalingo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scalingo_zip_url := "https://github.com/Scalingo/cli/archive/refs/tags/1.33.0.zip"
	scalingo_cmd_zip := exec.Command("curl", "-L", scalingo_zip_url, "-o", "package.zip")
	err = scalingo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scalingo_bin_url := "https://github.com/Scalingo/cli/archive/refs/tags/1.33.0.bin"
	scalingo_cmd_bin := exec.Command("curl", "-L", scalingo_bin_url, "-o", "binary.bin")
	err = scalingo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scalingo_src_url := "https://github.com/Scalingo/cli/archive/refs/tags/1.33.0.src.tar.gz"
	scalingo_cmd_src := exec.Command("curl", "-L", scalingo_src_url, "-o", "source.tar.gz")
	err = scalingo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scalingo_cmd_direct := exec.Command("./binary")
	err = scalingo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
