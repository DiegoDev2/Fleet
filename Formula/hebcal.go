package main

// Hebcal - Perpetual Jewish calendar for the command-line
// Homepage: https://github.com/hebcal/hebcal

import (
	"fmt"
	
	"os/exec"
)

func installHebcal() {
	// Método 1: Descargar y extraer .tar.gz
	hebcal_tar_url := "https://github.com/hebcal/hebcal/archive/refs/tags/v5.8.7.tar.gz"
	hebcal_cmd_tar := exec.Command("curl", "-L", hebcal_tar_url, "-o", "package.tar.gz")
	err := hebcal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hebcal_zip_url := "https://github.com/hebcal/hebcal/archive/refs/tags/v5.8.7.zip"
	hebcal_cmd_zip := exec.Command("curl", "-L", hebcal_zip_url, "-o", "package.zip")
	err = hebcal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hebcal_bin_url := "https://github.com/hebcal/hebcal/archive/refs/tags/v5.8.7.bin"
	hebcal_cmd_bin := exec.Command("curl", "-L", hebcal_bin_url, "-o", "binary.bin")
	err = hebcal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hebcal_src_url := "https://github.com/hebcal/hebcal/archive/refs/tags/v5.8.7.src.tar.gz"
	hebcal_cmd_src := exec.Command("curl", "-L", hebcal_src_url, "-o", "source.tar.gz")
	err = hebcal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hebcal_cmd_direct := exec.Command("./binary")
	err = hebcal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
