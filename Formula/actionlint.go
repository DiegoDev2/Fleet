package main

// Actionlint - Static checker for GitHub Actions workflow files
// Homepage: https://rhysd.github.io/actionlint/

import (
	"fmt"
	
	"os/exec"
)

func installActionlint() {
	// Método 1: Descargar y extraer .tar.gz
	actionlint_tar_url := "https://github.com/rhysd/actionlint/archive/refs/tags/v1.7.1.tar.gz"
	actionlint_cmd_tar := exec.Command("curl", "-L", actionlint_tar_url, "-o", "package.tar.gz")
	err := actionlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	actionlint_zip_url := "https://github.com/rhysd/actionlint/archive/refs/tags/v1.7.1.zip"
	actionlint_cmd_zip := exec.Command("curl", "-L", actionlint_zip_url, "-o", "package.zip")
	err = actionlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	actionlint_bin_url := "https://github.com/rhysd/actionlint/archive/refs/tags/v1.7.1.bin"
	actionlint_cmd_bin := exec.Command("curl", "-L", actionlint_bin_url, "-o", "binary.bin")
	err = actionlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	actionlint_src_url := "https://github.com/rhysd/actionlint/archive/refs/tags/v1.7.1.src.tar.gz"
	actionlint_cmd_src := exec.Command("curl", "-L", actionlint_src_url, "-o", "source.tar.gz")
	err = actionlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	actionlint_cmd_direct := exec.Command("./binary")
	err = actionlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: ronn")
	exec.Command("latte", "install", "ronn").Run()
}
