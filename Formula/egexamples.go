package main

// EgExamples - Useful examples at the command-line
// Homepage: https://github.com/srsudar/eg

import (
	"fmt"
	
	"os/exec"
)

func installEgExamples() {
	// Método 1: Descargar y extraer .tar.gz
	egexamples_tar_url := "https://files.pythonhosted.org/packages/5f/3f/f55eef404adae2d5429728722d6a81ad6ac50a80e9b47be046cfbe97bc44/eg-1.2.2.tar.gz"
	egexamples_cmd_tar := exec.Command("curl", "-L", egexamples_tar_url, "-o", "package.tar.gz")
	err := egexamples_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	egexamples_zip_url := "https://files.pythonhosted.org/packages/5f/3f/f55eef404adae2d5429728722d6a81ad6ac50a80e9b47be046cfbe97bc44/eg-1.2.2.zip"
	egexamples_cmd_zip := exec.Command("curl", "-L", egexamples_zip_url, "-o", "package.zip")
	err = egexamples_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	egexamples_bin_url := "https://files.pythonhosted.org/packages/5f/3f/f55eef404adae2d5429728722d6a81ad6ac50a80e9b47be046cfbe97bc44/eg-1.2.2.bin"
	egexamples_cmd_bin := exec.Command("curl", "-L", egexamples_bin_url, "-o", "binary.bin")
	err = egexamples_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	egexamples_src_url := "https://files.pythonhosted.org/packages/5f/3f/f55eef404adae2d5429728722d6a81ad6ac50a80e9b47be046cfbe97bc44/eg-1.2.2.src.tar.gz"
	egexamples_cmd_src := exec.Command("curl", "-L", egexamples_src_url, "-o", "source.tar.gz")
	err = egexamples_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	egexamples_cmd_direct := exec.Command("./binary")
	err = egexamples_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
