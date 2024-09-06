package main

// Pyupgrade - Upgrade syntax for newer versions of Python
// Homepage: https://github.com/asottile/pyupgrade

import (
	"fmt"
	
	"os/exec"
)

func installPyupgrade() {
	// Método 1: Descargar y extraer .tar.gz
	pyupgrade_tar_url := "https://files.pythonhosted.org/packages/7a/79/15cd93e47b5d670f0e32a540eb3f11bac4b5800cf1f796590eb448c6a768/pyupgrade-3.17.0.tar.gz"
	pyupgrade_cmd_tar := exec.Command("curl", "-L", pyupgrade_tar_url, "-o", "package.tar.gz")
	err := pyupgrade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyupgrade_zip_url := "https://files.pythonhosted.org/packages/7a/79/15cd93e47b5d670f0e32a540eb3f11bac4b5800cf1f796590eb448c6a768/pyupgrade-3.17.0.zip"
	pyupgrade_cmd_zip := exec.Command("curl", "-L", pyupgrade_zip_url, "-o", "package.zip")
	err = pyupgrade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyupgrade_bin_url := "https://files.pythonhosted.org/packages/7a/79/15cd93e47b5d670f0e32a540eb3f11bac4b5800cf1f796590eb448c6a768/pyupgrade-3.17.0.bin"
	pyupgrade_cmd_bin := exec.Command("curl", "-L", pyupgrade_bin_url, "-o", "binary.bin")
	err = pyupgrade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyupgrade_src_url := "https://files.pythonhosted.org/packages/7a/79/15cd93e47b5d670f0e32a540eb3f11bac4b5800cf1f796590eb448c6a768/pyupgrade-3.17.0.src.tar.gz"
	pyupgrade_cmd_src := exec.Command("curl", "-L", pyupgrade_src_url, "-o", "source.tar.gz")
	err = pyupgrade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyupgrade_cmd_direct := exec.Command("./binary")
	err = pyupgrade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
