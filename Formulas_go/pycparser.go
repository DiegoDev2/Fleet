package main

// Pycparser - C parser in Python
// Homepage: https://github.com/eliben/pycparser

import (
	"fmt"
	
	"os/exec"
)

func installPycparser() {
	// Método 1: Descargar y extraer .tar.gz
	pycparser_tar_url := "https://files.pythonhosted.org/packages/1d/b2/31537cf4b1ca988837256c910a668b553fceb8f069bedc4b1c826024b52c/pycparser-2.22.tar.gz"
	pycparser_cmd_tar := exec.Command("curl", "-L", pycparser_tar_url, "-o", "package.tar.gz")
	err := pycparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pycparser_zip_url := "https://files.pythonhosted.org/packages/1d/b2/31537cf4b1ca988837256c910a668b553fceb8f069bedc4b1c826024b52c/pycparser-2.22.zip"
	pycparser_cmd_zip := exec.Command("curl", "-L", pycparser_zip_url, "-o", "package.zip")
	err = pycparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pycparser_bin_url := "https://files.pythonhosted.org/packages/1d/b2/31537cf4b1ca988837256c910a668b553fceb8f069bedc4b1c826024b52c/pycparser-2.22.bin"
	pycparser_cmd_bin := exec.Command("curl", "-L", pycparser_bin_url, "-o", "binary.bin")
	err = pycparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pycparser_src_url := "https://files.pythonhosted.org/packages/1d/b2/31537cf4b1ca988837256c910a668b553fceb8f069bedc4b1c826024b52c/pycparser-2.22.src.tar.gz"
	pycparser_cmd_src := exec.Command("curl", "-L", pycparser_src_url, "-o", "source.tar.gz")
	err = pycparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pycparser_cmd_direct := exec.Command("./binary")
	err = pycparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
