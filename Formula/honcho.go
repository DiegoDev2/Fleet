package main

// Honcho - Python clone of Foreman, for managing Procfile-based applications
// Homepage: https://github.com/nickstenning/honcho

import (
	"fmt"
	
	"os/exec"
)

func installHoncho() {
	// Método 1: Descargar y extraer .tar.gz
	honcho_tar_url := "https://files.pythonhosted.org/packages/0e/7c/c0aa47711b5ada100273cbe190b33cc12297065ce559989699fd6c1ec0cb/honcho-1.1.0.tar.gz"
	honcho_cmd_tar := exec.Command("curl", "-L", honcho_tar_url, "-o", "package.tar.gz")
	err := honcho_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	honcho_zip_url := "https://files.pythonhosted.org/packages/0e/7c/c0aa47711b5ada100273cbe190b33cc12297065ce559989699fd6c1ec0cb/honcho-1.1.0.zip"
	honcho_cmd_zip := exec.Command("curl", "-L", honcho_zip_url, "-o", "package.zip")
	err = honcho_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	honcho_bin_url := "https://files.pythonhosted.org/packages/0e/7c/c0aa47711b5ada100273cbe190b33cc12297065ce559989699fd6c1ec0cb/honcho-1.1.0.bin"
	honcho_cmd_bin := exec.Command("curl", "-L", honcho_bin_url, "-o", "binary.bin")
	err = honcho_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	honcho_src_url := "https://files.pythonhosted.org/packages/0e/7c/c0aa47711b5ada100273cbe190b33cc12297065ce559989699fd6c1ec0cb/honcho-1.1.0.src.tar.gz"
	honcho_cmd_src := exec.Command("curl", "-L", honcho_src_url, "-o", "source.tar.gz")
	err = honcho_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	honcho_cmd_direct := exec.Command("./binary")
	err = honcho_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
