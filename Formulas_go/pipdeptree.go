package main

// Pipdeptree - CLI to display dependency tree of the installed Python packages
// Homepage: https://github.com/tox-dev/pipdeptree

import (
	"fmt"
	
	"os/exec"
)

func installPipdeptree() {
	// Método 1: Descargar y extraer .tar.gz
	pipdeptree_tar_url := "https://files.pythonhosted.org/packages/2b/1f/9bbf4f001befee00c050051f8f6e7971b41e2716784949989c7e856bf625/pipdeptree-2.23.1.tar.gz"
	pipdeptree_cmd_tar := exec.Command("curl", "-L", pipdeptree_tar_url, "-o", "package.tar.gz")
	err := pipdeptree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipdeptree_zip_url := "https://files.pythonhosted.org/packages/2b/1f/9bbf4f001befee00c050051f8f6e7971b41e2716784949989c7e856bf625/pipdeptree-2.23.1.zip"
	pipdeptree_cmd_zip := exec.Command("curl", "-L", pipdeptree_zip_url, "-o", "package.zip")
	err = pipdeptree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipdeptree_bin_url := "https://files.pythonhosted.org/packages/2b/1f/9bbf4f001befee00c050051f8f6e7971b41e2716784949989c7e856bf625/pipdeptree-2.23.1.bin"
	pipdeptree_cmd_bin := exec.Command("curl", "-L", pipdeptree_bin_url, "-o", "binary.bin")
	err = pipdeptree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipdeptree_src_url := "https://files.pythonhosted.org/packages/2b/1f/9bbf4f001befee00c050051f8f6e7971b41e2716784949989c7e856bf625/pipdeptree-2.23.1.src.tar.gz"
	pipdeptree_cmd_src := exec.Command("curl", "-L", pipdeptree_src_url, "-o", "source.tar.gz")
	err = pipdeptree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipdeptree_cmd_direct := exec.Command("./binary")
	err = pipdeptree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
