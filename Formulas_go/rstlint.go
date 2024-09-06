package main

// RstLint - ReStructuredText linter
// Homepage: https://github.com/twolfson/restructuredtext-lint

import (
	"fmt"
	
	"os/exec"
)

func installRstLint() {
	// Método 1: Descargar y extraer .tar.gz
	rstlint_tar_url := "https://files.pythonhosted.org/packages/48/9c/6d8035cafa2d2d314f34e6cd9313a299de095b26e96f1c7312878f988eec/restructuredtext_lint-1.4.0.tar.gz"
	rstlint_cmd_tar := exec.Command("curl", "-L", rstlint_tar_url, "-o", "package.tar.gz")
	err := rstlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rstlint_zip_url := "https://files.pythonhosted.org/packages/48/9c/6d8035cafa2d2d314f34e6cd9313a299de095b26e96f1c7312878f988eec/restructuredtext_lint-1.4.0.zip"
	rstlint_cmd_zip := exec.Command("curl", "-L", rstlint_zip_url, "-o", "package.zip")
	err = rstlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rstlint_bin_url := "https://files.pythonhosted.org/packages/48/9c/6d8035cafa2d2d314f34e6cd9313a299de095b26e96f1c7312878f988eec/restructuredtext_lint-1.4.0.bin"
	rstlint_cmd_bin := exec.Command("curl", "-L", rstlint_bin_url, "-o", "binary.bin")
	err = rstlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rstlint_src_url := "https://files.pythonhosted.org/packages/48/9c/6d8035cafa2d2d314f34e6cd9313a299de095b26e96f1c7312878f988eec/restructuredtext_lint-1.4.0.src.tar.gz"
	rstlint_cmd_src := exec.Command("curl", "-L", rstlint_src_url, "-o", "source.tar.gz")
	err = rstlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rstlint_cmd_direct := exec.Command("./binary")
	err = rstlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
