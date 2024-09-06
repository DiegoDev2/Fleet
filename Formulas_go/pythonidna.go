package main

// PythonIdna - Internationalized Domain Names in Applications (IDNA)
// Homepage: https://github.com/kjd/idna

import (
	"fmt"
	
	"os/exec"
)

func installPythonIdna() {
	// Método 1: Descargar y extraer .tar.gz
	pythonidna_tar_url := "https://files.pythonhosted.org/packages/bf/3f/ea4b9117521a1e9c50344b909be7886dd00a519552724809bb1f486986c2/idna-3.6.tar.gz"
	pythonidna_cmd_tar := exec.Command("curl", "-L", pythonidna_tar_url, "-o", "package.tar.gz")
	err := pythonidna_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonidna_zip_url := "https://files.pythonhosted.org/packages/bf/3f/ea4b9117521a1e9c50344b909be7886dd00a519552724809bb1f486986c2/idna-3.6.zip"
	pythonidna_cmd_zip := exec.Command("curl", "-L", pythonidna_zip_url, "-o", "package.zip")
	err = pythonidna_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonidna_bin_url := "https://files.pythonhosted.org/packages/bf/3f/ea4b9117521a1e9c50344b909be7886dd00a519552724809bb1f486986c2/idna-3.6.bin"
	pythonidna_cmd_bin := exec.Command("curl", "-L", pythonidna_bin_url, "-o", "binary.bin")
	err = pythonidna_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonidna_src_url := "https://files.pythonhosted.org/packages/bf/3f/ea4b9117521a1e9c50344b909be7886dd00a519552724809bb1f486986c2/idna-3.6.src.tar.gz"
	pythonidna_cmd_src := exec.Command("curl", "-L", pythonidna_src_url, "-o", "source.tar.gz")
	err = pythonidna_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonidna_cmd_direct := exec.Command("./binary")
	err = pythonidna_cmd_direct.Run()
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
