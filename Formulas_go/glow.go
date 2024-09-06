package main

// Glow - Render markdown on the CLI
// Homepage: https://github.com/charmbracelet/glow

import (
	"fmt"
	
	"os/exec"
)

func installGlow() {
	// Método 1: Descargar y extraer .tar.gz
	glow_tar_url := "https://github.com/charmbracelet/glow/archive/refs/tags/v2.0.0.tar.gz"
	glow_cmd_tar := exec.Command("curl", "-L", glow_tar_url, "-o", "package.tar.gz")
	err := glow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glow_zip_url := "https://github.com/charmbracelet/glow/archive/refs/tags/v2.0.0.zip"
	glow_cmd_zip := exec.Command("curl", "-L", glow_zip_url, "-o", "package.zip")
	err = glow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glow_bin_url := "https://github.com/charmbracelet/glow/archive/refs/tags/v2.0.0.bin"
	glow_cmd_bin := exec.Command("curl", "-L", glow_bin_url, "-o", "binary.bin")
	err = glow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glow_src_url := "https://github.com/charmbracelet/glow/archive/refs/tags/v2.0.0.src.tar.gz"
	glow_cmd_src := exec.Command("curl", "-L", glow_src_url, "-o", "source.tar.gz")
	err = glow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glow_cmd_direct := exec.Command("./binary")
	err = glow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
