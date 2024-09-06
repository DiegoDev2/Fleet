package main

// AnycableGo - WebSocket server with action cable protocol
// Homepage: https://github.com/anycable/anycable-go

import (
	"fmt"
	
	"os/exec"
)

func installAnycableGo() {
	// Método 1: Descargar y extraer .tar.gz
	anycablego_tar_url := "https://github.com/anycable/anycable-go/archive/refs/tags/v1.5.3.tar.gz"
	anycablego_cmd_tar := exec.Command("curl", "-L", anycablego_tar_url, "-o", "package.tar.gz")
	err := anycablego_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	anycablego_zip_url := "https://github.com/anycable/anycable-go/archive/refs/tags/v1.5.3.zip"
	anycablego_cmd_zip := exec.Command("curl", "-L", anycablego_zip_url, "-o", "package.zip")
	err = anycablego_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	anycablego_bin_url := "https://github.com/anycable/anycable-go/archive/refs/tags/v1.5.3.bin"
	anycablego_cmd_bin := exec.Command("curl", "-L", anycablego_bin_url, "-o", "binary.bin")
	err = anycablego_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	anycablego_src_url := "https://github.com/anycable/anycable-go/archive/refs/tags/v1.5.3.src.tar.gz"
	anycablego_cmd_src := exec.Command("curl", "-L", anycablego_src_url, "-o", "source.tar.gz")
	err = anycablego_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	anycablego_cmd_direct := exec.Command("./binary")
	err = anycablego_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
