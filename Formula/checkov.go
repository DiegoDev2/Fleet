package main

// Checkov - Prevent cloud misconfigurations during build-time for IaC tools
// Homepage: https://www.checkov.io/

import (
	"fmt"
	
	"os/exec"
)

func installCheckov() {
	// Método 1: Descargar y extraer .tar.gz
	checkov_tar_url := "https://files.pythonhosted.org/packages/73/44/e2f0c503d1b37d1a4f0660909e0745625dd0ec6fd01cb4dac2cb5c3082f5/checkov-3.2.240.tar.gz"
	checkov_cmd_tar := exec.Command("curl", "-L", checkov_tar_url, "-o", "package.tar.gz")
	err := checkov_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	checkov_zip_url := "https://files.pythonhosted.org/packages/73/44/e2f0c503d1b37d1a4f0660909e0745625dd0ec6fd01cb4dac2cb5c3082f5/checkov-3.2.240.zip"
	checkov_cmd_zip := exec.Command("curl", "-L", checkov_zip_url, "-o", "package.zip")
	err = checkov_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	checkov_bin_url := "https://files.pythonhosted.org/packages/73/44/e2f0c503d1b37d1a4f0660909e0745625dd0ec6fd01cb4dac2cb5c3082f5/checkov-3.2.240.bin"
	checkov_cmd_bin := exec.Command("curl", "-L", checkov_bin_url, "-o", "binary.bin")
	err = checkov_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	checkov_src_url := "https://files.pythonhosted.org/packages/73/44/e2f0c503d1b37d1a4f0660909e0745625dd0ec6fd01cb4dac2cb5c3082f5/checkov-3.2.240.src.tar.gz"
	checkov_cmd_src := exec.Command("curl", "-L", checkov_src_url, "-o", "source.tar.gz")
	err = checkov_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	checkov_cmd_direct := exec.Command("./binary")
	err = checkov_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
