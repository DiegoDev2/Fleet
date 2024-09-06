package main

// Codespell - Fix common misspellings in source code and text files
// Homepage: https://github.com/codespell-project/codespell

import (
	"fmt"
	
	"os/exec"
)

func installCodespell() {
	// Método 1: Descargar y extraer .tar.gz
	codespell_tar_url := "https://files.pythonhosted.org/packages/a0/a9/98353dfc7afcdf18cffd2dd3e959a25eaaf2728cf450caa59af89648a8e4/codespell-2.3.0.tar.gz"
	codespell_cmd_tar := exec.Command("curl", "-L", codespell_tar_url, "-o", "package.tar.gz")
	err := codespell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codespell_zip_url := "https://files.pythonhosted.org/packages/a0/a9/98353dfc7afcdf18cffd2dd3e959a25eaaf2728cf450caa59af89648a8e4/codespell-2.3.0.zip"
	codespell_cmd_zip := exec.Command("curl", "-L", codespell_zip_url, "-o", "package.zip")
	err = codespell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codespell_bin_url := "https://files.pythonhosted.org/packages/a0/a9/98353dfc7afcdf18cffd2dd3e959a25eaaf2728cf450caa59af89648a8e4/codespell-2.3.0.bin"
	codespell_cmd_bin := exec.Command("curl", "-L", codespell_bin_url, "-o", "binary.bin")
	err = codespell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codespell_src_url := "https://files.pythonhosted.org/packages/a0/a9/98353dfc7afcdf18cffd2dd3e959a25eaaf2728cf450caa59af89648a8e4/codespell-2.3.0.src.tar.gz"
	codespell_cmd_src := exec.Command("curl", "-L", codespell_src_url, "-o", "source.tar.gz")
	err = codespell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codespell_cmd_direct := exec.Command("./binary")
	err = codespell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
