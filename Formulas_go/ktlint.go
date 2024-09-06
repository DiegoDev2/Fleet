package main

// Ktlint - Anti-bikeshedding Kotlin linter with built-in formatter
// Homepage: https://ktlint.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installKtlint() {
	// Método 1: Descargar y extraer .tar.gz
	ktlint_tar_url := "https://github.com/pinterest/ktlint/releases/download/1.3.1/ktlint-1.3.1.zip"
	ktlint_cmd_tar := exec.Command("curl", "-L", ktlint_tar_url, "-o", "package.tar.gz")
	err := ktlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ktlint_zip_url := "https://github.com/pinterest/ktlint/releases/download/1.3.1/ktlint-1.3.1.zip"
	ktlint_cmd_zip := exec.Command("curl", "-L", ktlint_zip_url, "-o", "package.zip")
	err = ktlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ktlint_bin_url := "https://github.com/pinterest/ktlint/releases/download/1.3.1/ktlint-1.3.1.zip"
	ktlint_cmd_bin := exec.Command("curl", "-L", ktlint_bin_url, "-o", "binary.bin")
	err = ktlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ktlint_src_url := "https://github.com/pinterest/ktlint/releases/download/1.3.1/ktlint-1.3.1.zip"
	ktlint_cmd_src := exec.Command("curl", "-L", ktlint_src_url, "-o", "source.tar.gz")
	err = ktlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ktlint_cmd_direct := exec.Command("./binary")
	err = ktlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
