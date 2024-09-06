package main

// Gitlint - Linting for your git commit messages
// Homepage: https://jorisroovers.com/gitlint/

import (
	"fmt"
	
	"os/exec"
)

func installGitlint() {
	// Método 1: Descargar y extraer .tar.gz
	gitlint_tar_url := "https://files.pythonhosted.org/packages/73/51/b59270264aabcab5b933f3eb9bfb022464ca9205b04feef1bdc1635fd9b4/gitlint_core-0.19.1.tar.gz"
	gitlint_cmd_tar := exec.Command("curl", "-L", gitlint_tar_url, "-o", "package.tar.gz")
	err := gitlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitlint_zip_url := "https://files.pythonhosted.org/packages/73/51/b59270264aabcab5b933f3eb9bfb022464ca9205b04feef1bdc1635fd9b4/gitlint_core-0.19.1.zip"
	gitlint_cmd_zip := exec.Command("curl", "-L", gitlint_zip_url, "-o", "package.zip")
	err = gitlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitlint_bin_url := "https://files.pythonhosted.org/packages/73/51/b59270264aabcab5b933f3eb9bfb022464ca9205b04feef1bdc1635fd9b4/gitlint_core-0.19.1.bin"
	gitlint_cmd_bin := exec.Command("curl", "-L", gitlint_bin_url, "-o", "binary.bin")
	err = gitlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitlint_src_url := "https://files.pythonhosted.org/packages/73/51/b59270264aabcab5b933f3eb9bfb022464ca9205b04feef1bdc1635fd9b4/gitlint_core-0.19.1.src.tar.gz"
	gitlint_cmd_src := exec.Command("curl", "-L", gitlint_src_url, "-o", "source.tar.gz")
	err = gitlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitlint_cmd_direct := exec.Command("./binary")
	err = gitlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
