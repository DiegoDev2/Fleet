package main

// Ctlptl - Making local Kubernetes clusters fun and easy to set up
// Homepage: https://github.com/tilt-dev/ctlptl

import (
	"fmt"
	
	"os/exec"
)

func installCtlptl() {
	// Método 1: Descargar y extraer .tar.gz
	ctlptl_tar_url := "https://github.com/tilt-dev/ctlptl/archive/refs/tags/v0.8.33.tar.gz"
	ctlptl_cmd_tar := exec.Command("curl", "-L", ctlptl_tar_url, "-o", "package.tar.gz")
	err := ctlptl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ctlptl_zip_url := "https://github.com/tilt-dev/ctlptl/archive/refs/tags/v0.8.33.zip"
	ctlptl_cmd_zip := exec.Command("curl", "-L", ctlptl_zip_url, "-o", "package.zip")
	err = ctlptl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ctlptl_bin_url := "https://github.com/tilt-dev/ctlptl/archive/refs/tags/v0.8.33.bin"
	ctlptl_cmd_bin := exec.Command("curl", "-L", ctlptl_bin_url, "-o", "binary.bin")
	err = ctlptl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ctlptl_src_url := "https://github.com/tilt-dev/ctlptl/archive/refs/tags/v0.8.33.src.tar.gz"
	ctlptl_cmd_src := exec.Command("curl", "-L", ctlptl_src_url, "-o", "source.tar.gz")
	err = ctlptl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ctlptl_cmd_direct := exec.Command("./binary")
	err = ctlptl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
