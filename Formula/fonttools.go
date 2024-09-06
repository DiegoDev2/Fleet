package main

// Fonttools - Library for manipulating fonts
// Homepage: https://github.com/fonttools/fonttools

import (
	"fmt"
	
	"os/exec"
)

func installFonttools() {
	// Método 1: Descargar y extraer .tar.gz
	fonttools_tar_url := "https://files.pythonhosted.org/packages/c6/cb/cd80a0da995adde8ade6044a8744aee0da5efea01301cadf770f7fbe7dcc/fonttools-4.53.1.tar.gz"
	fonttools_cmd_tar := exec.Command("curl", "-L", fonttools_tar_url, "-o", "package.tar.gz")
	err := fonttools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fonttools_zip_url := "https://files.pythonhosted.org/packages/c6/cb/cd80a0da995adde8ade6044a8744aee0da5efea01301cadf770f7fbe7dcc/fonttools-4.53.1.zip"
	fonttools_cmd_zip := exec.Command("curl", "-L", fonttools_zip_url, "-o", "package.zip")
	err = fonttools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fonttools_bin_url := "https://files.pythonhosted.org/packages/c6/cb/cd80a0da995adde8ade6044a8744aee0da5efea01301cadf770f7fbe7dcc/fonttools-4.53.1.bin"
	fonttools_cmd_bin := exec.Command("curl", "-L", fonttools_bin_url, "-o", "binary.bin")
	err = fonttools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fonttools_src_url := "https://files.pythonhosted.org/packages/c6/cb/cd80a0da995adde8ade6044a8744aee0da5efea01301cadf770f7fbe7dcc/fonttools-4.53.1.src.tar.gz"
	fonttools_cmd_src := exec.Command("curl", "-L", fonttools_src_url, "-o", "source.tar.gz")
	err = fonttools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fonttools_cmd_direct := exec.Command("./binary")
	err = fonttools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
