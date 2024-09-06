package main

// SphinxDoc - Tool to create intelligent and beautiful documentation
// Homepage: https://www.sphinx-doc.org/

import (
	"fmt"
	
	"os/exec"
)

func installSphinxDoc() {
	// Método 1: Descargar y extraer .tar.gz
	sphinxdoc_tar_url := "https://files.pythonhosted.org/packages/25/a7/3cc3d6dcad70aba2e32a3ae8de5a90026a0a2fdaaa0756925e3a120249b6/sphinx-8.0.2.tar.gz"
	sphinxdoc_cmd_tar := exec.Command("curl", "-L", sphinxdoc_tar_url, "-o", "package.tar.gz")
	err := sphinxdoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sphinxdoc_zip_url := "https://files.pythonhosted.org/packages/25/a7/3cc3d6dcad70aba2e32a3ae8de5a90026a0a2fdaaa0756925e3a120249b6/sphinx-8.0.2.zip"
	sphinxdoc_cmd_zip := exec.Command("curl", "-L", sphinxdoc_zip_url, "-o", "package.zip")
	err = sphinxdoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sphinxdoc_bin_url := "https://files.pythonhosted.org/packages/25/a7/3cc3d6dcad70aba2e32a3ae8de5a90026a0a2fdaaa0756925e3a120249b6/sphinx-8.0.2.bin"
	sphinxdoc_cmd_bin := exec.Command("curl", "-L", sphinxdoc_bin_url, "-o", "binary.bin")
	err = sphinxdoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sphinxdoc_src_url := "https://files.pythonhosted.org/packages/25/a7/3cc3d6dcad70aba2e32a3ae8de5a90026a0a2fdaaa0756925e3a120249b6/sphinx-8.0.2.src.tar.gz"
	sphinxdoc_cmd_src := exec.Command("curl", "-L", sphinxdoc_src_url, "-o", "source.tar.gz")
	err = sphinxdoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sphinxdoc_cmd_direct := exec.Command("./binary")
	err = sphinxdoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
