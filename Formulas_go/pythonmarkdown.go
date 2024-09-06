package main

// PythonMarkdown - Python implementation of Markdown
// Homepage: https://python-markdown.github.io

import (
	"fmt"
	
	"os/exec"
)

func installPythonMarkdown() {
	// Método 1: Descargar y extraer .tar.gz
	pythonmarkdown_tar_url := "https://files.pythonhosted.org/packages/54/28/3af612670f82f4c056911fbbbb42760255801b3068c48de792d354ff4472/markdown-3.7.tar.gz"
	pythonmarkdown_cmd_tar := exec.Command("curl", "-L", pythonmarkdown_tar_url, "-o", "package.tar.gz")
	err := pythonmarkdown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonmarkdown_zip_url := "https://files.pythonhosted.org/packages/54/28/3af612670f82f4c056911fbbbb42760255801b3068c48de792d354ff4472/markdown-3.7.zip"
	pythonmarkdown_cmd_zip := exec.Command("curl", "-L", pythonmarkdown_zip_url, "-o", "package.zip")
	err = pythonmarkdown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonmarkdown_bin_url := "https://files.pythonhosted.org/packages/54/28/3af612670f82f4c056911fbbbb42760255801b3068c48de792d354ff4472/markdown-3.7.bin"
	pythonmarkdown_cmd_bin := exec.Command("curl", "-L", pythonmarkdown_bin_url, "-o", "binary.bin")
	err = pythonmarkdown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonmarkdown_src_url := "https://files.pythonhosted.org/packages/54/28/3af612670f82f4c056911fbbbb42760255801b3068c48de792d354ff4472/markdown-3.7.src.tar.gz"
	pythonmarkdown_cmd_src := exec.Command("curl", "-L", pythonmarkdown_src_url, "-o", "source.tar.gz")
	err = pythonmarkdown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonmarkdown_cmd_direct := exec.Command("./binary")
	err = pythonmarkdown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
