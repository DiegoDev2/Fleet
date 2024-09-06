package main

// Diceware - Passphrases to remember
// Homepage: https://github.com/ulif/diceware

import (
	"fmt"
	
	"os/exec"
)

func installDiceware() {
	// Método 1: Descargar y extraer .tar.gz
	diceware_tar_url := "https://files.pythonhosted.org/packages/2f/7b/2ebe60ee2360170d93f1c3f1e4429353c8445992fc2bc501e98013697c71/diceware-0.10.tar.gz"
	diceware_cmd_tar := exec.Command("curl", "-L", diceware_tar_url, "-o", "package.tar.gz")
	err := diceware_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diceware_zip_url := "https://files.pythonhosted.org/packages/2f/7b/2ebe60ee2360170d93f1c3f1e4429353c8445992fc2bc501e98013697c71/diceware-0.10.zip"
	diceware_cmd_zip := exec.Command("curl", "-L", diceware_zip_url, "-o", "package.zip")
	err = diceware_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diceware_bin_url := "https://files.pythonhosted.org/packages/2f/7b/2ebe60ee2360170d93f1c3f1e4429353c8445992fc2bc501e98013697c71/diceware-0.10.bin"
	diceware_cmd_bin := exec.Command("curl", "-L", diceware_bin_url, "-o", "binary.bin")
	err = diceware_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diceware_src_url := "https://files.pythonhosted.org/packages/2f/7b/2ebe60ee2360170d93f1c3f1e4429353c8445992fc2bc501e98013697c71/diceware-0.10.src.tar.gz"
	diceware_cmd_src := exec.Command("curl", "-L", diceware_src_url, "-o", "source.tar.gz")
	err = diceware_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diceware_cmd_direct := exec.Command("./binary")
	err = diceware_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
