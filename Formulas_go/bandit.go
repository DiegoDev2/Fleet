package main

// Bandit - Security-oriented static analyser for Python code
// Homepage: https://github.com/PyCQA/bandit

import (
	"fmt"
	
	"os/exec"
)

func installBandit() {
	// Método 1: Descargar y extraer .tar.gz
	bandit_tar_url := "https://files.pythonhosted.org/packages/1c/a4/ee391b0f046a6d8919eef246aed7c39849e299cc2e50d918b54add397de6/bandit-1.7.9.tar.gz"
	bandit_cmd_tar := exec.Command("curl", "-L", bandit_tar_url, "-o", "package.tar.gz")
	err := bandit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bandit_zip_url := "https://files.pythonhosted.org/packages/1c/a4/ee391b0f046a6d8919eef246aed7c39849e299cc2e50d918b54add397de6/bandit-1.7.9.zip"
	bandit_cmd_zip := exec.Command("curl", "-L", bandit_zip_url, "-o", "package.zip")
	err = bandit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bandit_bin_url := "https://files.pythonhosted.org/packages/1c/a4/ee391b0f046a6d8919eef246aed7c39849e299cc2e50d918b54add397de6/bandit-1.7.9.bin"
	bandit_cmd_bin := exec.Command("curl", "-L", bandit_bin_url, "-o", "binary.bin")
	err = bandit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bandit_src_url := "https://files.pythonhosted.org/packages/1c/a4/ee391b0f046a6d8919eef246aed7c39849e299cc2e50d918b54add397de6/bandit-1.7.9.src.tar.gz"
	bandit_cmd_src := exec.Command("curl", "-L", bandit_src_url, "-o", "source.tar.gz")
	err = bandit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bandit_cmd_direct := exec.Command("./binary")
	err = bandit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
