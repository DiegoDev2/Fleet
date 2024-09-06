package main

// Wego - Weather app for the terminal
// Homepage: https://github.com/schachmat/wego

import (
	"fmt"
	
	"os/exec"
)

func installWego() {
	// Método 1: Descargar y extraer .tar.gz
	wego_tar_url := "https://github.com/schachmat/wego/archive/refs/tags/2.3.tar.gz"
	wego_cmd_tar := exec.Command("curl", "-L", wego_tar_url, "-o", "package.tar.gz")
	err := wego_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wego_zip_url := "https://github.com/schachmat/wego/archive/refs/tags/2.3.zip"
	wego_cmd_zip := exec.Command("curl", "-L", wego_zip_url, "-o", "package.zip")
	err = wego_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wego_bin_url := "https://github.com/schachmat/wego/archive/refs/tags/2.3.bin"
	wego_cmd_bin := exec.Command("curl", "-L", wego_bin_url, "-o", "binary.bin")
	err = wego_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wego_src_url := "https://github.com/schachmat/wego/archive/refs/tags/2.3.src.tar.gz"
	wego_cmd_src := exec.Command("curl", "-L", wego_src_url, "-o", "source.tar.gz")
	err = wego_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wego_cmd_direct := exec.Command("./binary")
	err = wego_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
