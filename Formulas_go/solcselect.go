package main

// SolcSelect - Manage multiple Solidity compiler versions
// Homepage: https://github.com/crytic/solc-select

import (
	"fmt"
	
	"os/exec"
)

func installSolcSelect() {
	// Método 1: Descargar y extraer .tar.gz
	solcselect_tar_url := "https://files.pythonhosted.org/packages/60/a0/2a2bfbbab1d9bd4e1a24e3604c30b5d6f84219238f3c98f06191faf5d019/solc-select-1.0.4.tar.gz"
	solcselect_cmd_tar := exec.Command("curl", "-L", solcselect_tar_url, "-o", "package.tar.gz")
	err := solcselect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solcselect_zip_url := "https://files.pythonhosted.org/packages/60/a0/2a2bfbbab1d9bd4e1a24e3604c30b5d6f84219238f3c98f06191faf5d019/solc-select-1.0.4.zip"
	solcselect_cmd_zip := exec.Command("curl", "-L", solcselect_zip_url, "-o", "package.zip")
	err = solcselect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solcselect_bin_url := "https://files.pythonhosted.org/packages/60/a0/2a2bfbbab1d9bd4e1a24e3604c30b5d6f84219238f3c98f06191faf5d019/solc-select-1.0.4.bin"
	solcselect_cmd_bin := exec.Command("curl", "-L", solcselect_bin_url, "-o", "binary.bin")
	err = solcselect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solcselect_src_url := "https://files.pythonhosted.org/packages/60/a0/2a2bfbbab1d9bd4e1a24e3604c30b5d6f84219238f3c98f06191faf5d019/solc-select-1.0.4.src.tar.gz"
	solcselect_cmd_src := exec.Command("curl", "-L", solcselect_src_url, "-o", "source.tar.gz")
	err = solcselect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solcselect_cmd_direct := exec.Command("./binary")
	err = solcselect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
