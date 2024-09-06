package main

// Gum - Tool for glamorous shell scripts
// Homepage: https://github.com/charmbracelet/gum

import (
	"fmt"
	
	"os/exec"
)

func installGum() {
	// Método 1: Descargar y extraer .tar.gz
	gum_tar_url := "https://github.com/charmbracelet/gum/archive/refs/tags/v0.14.4.tar.gz"
	gum_cmd_tar := exec.Command("curl", "-L", gum_tar_url, "-o", "package.tar.gz")
	err := gum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gum_zip_url := "https://github.com/charmbracelet/gum/archive/refs/tags/v0.14.4.zip"
	gum_cmd_zip := exec.Command("curl", "-L", gum_zip_url, "-o", "package.zip")
	err = gum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gum_bin_url := "https://github.com/charmbracelet/gum/archive/refs/tags/v0.14.4.bin"
	gum_cmd_bin := exec.Command("curl", "-L", gum_bin_url, "-o", "binary.bin")
	err = gum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gum_src_url := "https://github.com/charmbracelet/gum/archive/refs/tags/v0.14.4.src.tar.gz"
	gum_cmd_src := exec.Command("curl", "-L", gum_src_url, "-o", "source.tar.gz")
	err = gum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gum_cmd_direct := exec.Command("./binary")
	err = gum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
