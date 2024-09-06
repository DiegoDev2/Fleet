package main

// Glances - Alternative to top/htop
// Homepage: https://nicolargo.github.io/glances/

import (
	"fmt"
	
	"os/exec"
)

func installGlances() {
	// Método 1: Descargar y extraer .tar.gz
	glances_tar_url := "https://files.pythonhosted.org/packages/1b/98/1ee2cf6c1c3d84f69ba23d5cd77973d04e8bf7136fe7a44416a408e05ff0/glances-4.1.2.tar.gz"
	glances_cmd_tar := exec.Command("curl", "-L", glances_tar_url, "-o", "package.tar.gz")
	err := glances_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glances_zip_url := "https://files.pythonhosted.org/packages/1b/98/1ee2cf6c1c3d84f69ba23d5cd77973d04e8bf7136fe7a44416a408e05ff0/glances-4.1.2.zip"
	glances_cmd_zip := exec.Command("curl", "-L", glances_zip_url, "-o", "package.zip")
	err = glances_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glances_bin_url := "https://files.pythonhosted.org/packages/1b/98/1ee2cf6c1c3d84f69ba23d5cd77973d04e8bf7136fe7a44416a408e05ff0/glances-4.1.2.bin"
	glances_cmd_bin := exec.Command("curl", "-L", glances_bin_url, "-o", "binary.bin")
	err = glances_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glances_src_url := "https://files.pythonhosted.org/packages/1b/98/1ee2cf6c1c3d84f69ba23d5cd77973d04e8bf7136fe7a44416a408e05ff0/glances-4.1.2.src.tar.gz"
	glances_cmd_src := exec.Command("curl", "-L", glances_src_url, "-o", "source.tar.gz")
	err = glances_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glances_cmd_direct := exec.Command("./binary")
	err = glances_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
