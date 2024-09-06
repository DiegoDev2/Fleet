package main

// Commitizen - Defines a standard way of committing rules and communicating it
// Homepage: https://commitizen-tools.github.io/commitizen/

import (
	"fmt"
	
	"os/exec"
)

func installCommitizen() {
	// Método 1: Descargar y extraer .tar.gz
	commitizen_tar_url := "https://files.pythonhosted.org/packages/7b/7b/9bd4792a523b4a5b7262d64a2ae9f63bea27469eeecaa0dabf737f40c143/commitizen-3.29.0.tar.gz"
	commitizen_cmd_tar := exec.Command("curl", "-L", commitizen_tar_url, "-o", "package.tar.gz")
	err := commitizen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	commitizen_zip_url := "https://files.pythonhosted.org/packages/7b/7b/9bd4792a523b4a5b7262d64a2ae9f63bea27469eeecaa0dabf737f40c143/commitizen-3.29.0.zip"
	commitizen_cmd_zip := exec.Command("curl", "-L", commitizen_zip_url, "-o", "package.zip")
	err = commitizen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	commitizen_bin_url := "https://files.pythonhosted.org/packages/7b/7b/9bd4792a523b4a5b7262d64a2ae9f63bea27469eeecaa0dabf737f40c143/commitizen-3.29.0.bin"
	commitizen_cmd_bin := exec.Command("curl", "-L", commitizen_bin_url, "-o", "binary.bin")
	err = commitizen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	commitizen_src_url := "https://files.pythonhosted.org/packages/7b/7b/9bd4792a523b4a5b7262d64a2ae9f63bea27469eeecaa0dabf737f40c143/commitizen-3.29.0.src.tar.gz"
	commitizen_cmd_src := exec.Command("curl", "-L", commitizen_src_url, "-o", "source.tar.gz")
	err = commitizen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	commitizen_cmd_direct := exec.Command("./binary")
	err = commitizen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
