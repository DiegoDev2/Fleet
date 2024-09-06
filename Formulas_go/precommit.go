package main

// PreCommit - Framework for managing multi-language pre-commit hooks
// Homepage: https://pre-commit.com/

import (
	"fmt"
	
	"os/exec"
)

func installPreCommit() {
	// Método 1: Descargar y extraer .tar.gz
	precommit_tar_url := "https://files.pythonhosted.org/packages/64/10/97ee2fa54dff1e9da9badbc5e35d0bbaef0776271ea5907eccf64140f72f/pre_commit-3.8.0.tar.gz"
	precommit_cmd_tar := exec.Command("curl", "-L", precommit_tar_url, "-o", "package.tar.gz")
	err := precommit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	precommit_zip_url := "https://files.pythonhosted.org/packages/64/10/97ee2fa54dff1e9da9badbc5e35d0bbaef0776271ea5907eccf64140f72f/pre_commit-3.8.0.zip"
	precommit_cmd_zip := exec.Command("curl", "-L", precommit_zip_url, "-o", "package.zip")
	err = precommit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	precommit_bin_url := "https://files.pythonhosted.org/packages/64/10/97ee2fa54dff1e9da9badbc5e35d0bbaef0776271ea5907eccf64140f72f/pre_commit-3.8.0.bin"
	precommit_cmd_bin := exec.Command("curl", "-L", precommit_bin_url, "-o", "binary.bin")
	err = precommit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	precommit_src_url := "https://files.pythonhosted.org/packages/64/10/97ee2fa54dff1e9da9badbc5e35d0bbaef0776271ea5907eccf64140f72f/pre_commit-3.8.0.src.tar.gz"
	precommit_cmd_src := exec.Command("curl", "-L", precommit_src_url, "-o", "source.tar.gz")
	err = precommit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	precommit_cmd_direct := exec.Command("./binary")
	err = precommit_cmd_direct.Run()
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
