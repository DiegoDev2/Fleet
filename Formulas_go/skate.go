package main

// Skate - Personal key value store
// Homepage: https://github.com/charmbracelet/skate

import (
	"fmt"
	
	"os/exec"
)

func installSkate() {
	// Método 1: Descargar y extraer .tar.gz
	skate_tar_url := "https://github.com/charmbracelet/skate/archive/refs/tags/v1.0.0.tar.gz"
	skate_cmd_tar := exec.Command("curl", "-L", skate_tar_url, "-o", "package.tar.gz")
	err := skate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skate_zip_url := "https://github.com/charmbracelet/skate/archive/refs/tags/v1.0.0.zip"
	skate_cmd_zip := exec.Command("curl", "-L", skate_zip_url, "-o", "package.zip")
	err = skate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skate_bin_url := "https://github.com/charmbracelet/skate/archive/refs/tags/v1.0.0.bin"
	skate_cmd_bin := exec.Command("curl", "-L", skate_bin_url, "-o", "binary.bin")
	err = skate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skate_src_url := "https://github.com/charmbracelet/skate/archive/refs/tags/v1.0.0.src.tar.gz"
	skate_cmd_src := exec.Command("curl", "-L", skate_src_url, "-o", "source.tar.gz")
	err = skate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skate_cmd_direct := exec.Command("./binary")
	err = skate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
