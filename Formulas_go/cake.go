package main

// Cake - Cross platform build automation system with a C# DSL
// Homepage: https://cakebuild.net/

import (
	"fmt"
	
	"os/exec"
)

func installCake() {
	// Método 1: Descargar y extraer .tar.gz
	cake_tar_url := "https://github.com/cake-build/cake/archive/refs/tags/v4.0.0.tar.gz"
	cake_cmd_tar := exec.Command("curl", "-L", cake_tar_url, "-o", "package.tar.gz")
	err := cake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cake_zip_url := "https://github.com/cake-build/cake/archive/refs/tags/v4.0.0.zip"
	cake_cmd_zip := exec.Command("curl", "-L", cake_zip_url, "-o", "package.zip")
	err = cake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cake_bin_url := "https://github.com/cake-build/cake/archive/refs/tags/v4.0.0.bin"
	cake_cmd_bin := exec.Command("curl", "-L", cake_bin_url, "-o", "binary.bin")
	err = cake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cake_src_url := "https://github.com/cake-build/cake/archive/refs/tags/v4.0.0.src.tar.gz"
	cake_cmd_src := exec.Command("curl", "-L", cake_src_url, "-o", "source.tar.gz")
	err = cake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cake_cmd_direct := exec.Command("./binary")
	err = cake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet")
exec.Command("latte", "install", "dotnet")
}
