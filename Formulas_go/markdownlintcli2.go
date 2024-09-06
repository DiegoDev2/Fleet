package main

// MarkdownlintCli2 - Fast, flexible, config-based cli for linting Markdown/CommonMark files
// Homepage: https://github.com/DavidAnson/markdownlint-cli2

import (
	"fmt"
	
	"os/exec"
)

func installMarkdownlintCli2() {
	// Método 1: Descargar y extraer .tar.gz
	markdownlintcli2_tar_url := "https://registry.npmjs.org/markdownlint-cli2/-/markdownlint-cli2-0.13.0.tgz"
	markdownlintcli2_cmd_tar := exec.Command("curl", "-L", markdownlintcli2_tar_url, "-o", "package.tar.gz")
	err := markdownlintcli2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	markdownlintcli2_zip_url := "https://registry.npmjs.org/markdownlint-cli2/-/markdownlint-cli2-0.13.0.tgz"
	markdownlintcli2_cmd_zip := exec.Command("curl", "-L", markdownlintcli2_zip_url, "-o", "package.zip")
	err = markdownlintcli2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	markdownlintcli2_bin_url := "https://registry.npmjs.org/markdownlint-cli2/-/markdownlint-cli2-0.13.0.tgz"
	markdownlintcli2_cmd_bin := exec.Command("curl", "-L", markdownlintcli2_bin_url, "-o", "binary.bin")
	err = markdownlintcli2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	markdownlintcli2_src_url := "https://registry.npmjs.org/markdownlint-cli2/-/markdownlint-cli2-0.13.0.tgz"
	markdownlintcli2_cmd_src := exec.Command("curl", "-L", markdownlintcli2_src_url, "-o", "source.tar.gz")
	err = markdownlintcli2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	markdownlintcli2_cmd_direct := exec.Command("./binary")
	err = markdownlintcli2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
