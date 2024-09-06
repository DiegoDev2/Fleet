package main

// Peru - Dependency retriever for version control and archives
// Homepage: https://github.com/buildinspace/peru

import (
	"fmt"
	
	"os/exec"
)

func installPeru() {
	// Método 1: Descargar y extraer .tar.gz
	peru_tar_url := "https://files.pythonhosted.org/packages/8e/c7/c451e70443c0b82440384d51f4b9517b921d4fe44172d63dc10a09da114f/peru-1.3.1.tar.gz"
	peru_cmd_tar := exec.Command("curl", "-L", peru_tar_url, "-o", "package.tar.gz")
	err := peru_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	peru_zip_url := "https://files.pythonhosted.org/packages/8e/c7/c451e70443c0b82440384d51f4b9517b921d4fe44172d63dc10a09da114f/peru-1.3.1.zip"
	peru_cmd_zip := exec.Command("curl", "-L", peru_zip_url, "-o", "package.zip")
	err = peru_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	peru_bin_url := "https://files.pythonhosted.org/packages/8e/c7/c451e70443c0b82440384d51f4b9517b921d4fe44172d63dc10a09da114f/peru-1.3.1.bin"
	peru_cmd_bin := exec.Command("curl", "-L", peru_bin_url, "-o", "binary.bin")
	err = peru_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	peru_src_url := "https://files.pythonhosted.org/packages/8e/c7/c451e70443c0b82440384d51f4b9517b921d4fe44172d63dc10a09da114f/peru-1.3.1.src.tar.gz"
	peru_cmd_src := exec.Command("curl", "-L", peru_src_url, "-o", "source.tar.gz")
	err = peru_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	peru_cmd_direct := exec.Command("./binary")
	err = peru_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
