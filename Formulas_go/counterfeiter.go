package main

// Counterfeiter - Tool for generating self-contained, type-safe test doubles in go
// Homepage: https://github.com/maxbrunsfeld/counterfeiter

import (
	"fmt"
	
	"os/exec"
)

func installCounterfeiter() {
	// Método 1: Descargar y extraer .tar.gz
	counterfeiter_tar_url := "https://github.com/maxbrunsfeld/counterfeiter/archive/refs/tags/v6.8.1.tar.gz"
	counterfeiter_cmd_tar := exec.Command("curl", "-L", counterfeiter_tar_url, "-o", "package.tar.gz")
	err := counterfeiter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	counterfeiter_zip_url := "https://github.com/maxbrunsfeld/counterfeiter/archive/refs/tags/v6.8.1.zip"
	counterfeiter_cmd_zip := exec.Command("curl", "-L", counterfeiter_zip_url, "-o", "package.zip")
	err = counterfeiter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	counterfeiter_bin_url := "https://github.com/maxbrunsfeld/counterfeiter/archive/refs/tags/v6.8.1.bin"
	counterfeiter_cmd_bin := exec.Command("curl", "-L", counterfeiter_bin_url, "-o", "binary.bin")
	err = counterfeiter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	counterfeiter_src_url := "https://github.com/maxbrunsfeld/counterfeiter/archive/refs/tags/v6.8.1.src.tar.gz"
	counterfeiter_cmd_src := exec.Command("curl", "-L", counterfeiter_src_url, "-o", "source.tar.gz")
	err = counterfeiter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	counterfeiter_cmd_direct := exec.Command("./binary")
	err = counterfeiter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
