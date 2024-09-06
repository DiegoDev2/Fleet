package main

// Legit - Command-line interface for Git, optimized for workflow simplicity
// Homepage: https://frostming.github.io/legit/

import (
	"fmt"
	
	"os/exec"
)

func installLegit() {
	// Método 1: Descargar y extraer .tar.gz
	legit_tar_url := "https://files.pythonhosted.org/packages/cb/e4/8cc5904c486241bf2edc4dd84f357fa96686dc85f48eedb835af65f821bf/legit-1.2.0.post0.tar.gz"
	legit_cmd_tar := exec.Command("curl", "-L", legit_tar_url, "-o", "package.tar.gz")
	err := legit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	legit_zip_url := "https://files.pythonhosted.org/packages/cb/e4/8cc5904c486241bf2edc4dd84f357fa96686dc85f48eedb835af65f821bf/legit-1.2.0.post0.zip"
	legit_cmd_zip := exec.Command("curl", "-L", legit_zip_url, "-o", "package.zip")
	err = legit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	legit_bin_url := "https://files.pythonhosted.org/packages/cb/e4/8cc5904c486241bf2edc4dd84f357fa96686dc85f48eedb835af65f821bf/legit-1.2.0.post0.bin"
	legit_cmd_bin := exec.Command("curl", "-L", legit_bin_url, "-o", "binary.bin")
	err = legit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	legit_src_url := "https://files.pythonhosted.org/packages/cb/e4/8cc5904c486241bf2edc4dd84f357fa96686dc85f48eedb835af65f821bf/legit-1.2.0.post0.src.tar.gz"
	legit_cmd_src := exec.Command("curl", "-L", legit_src_url, "-o", "source.tar.gz")
	err = legit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	legit_cmd_direct := exec.Command("./binary")
	err = legit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
