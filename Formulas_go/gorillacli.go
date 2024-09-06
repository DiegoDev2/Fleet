package main

// GorillaCli - LLMs for your CLI
// Homepage: https://gorilla.cs.berkeley.edu/

import (
	"fmt"
	
	"os/exec"
)

func installGorillaCli() {
	// Método 1: Descargar y extraer .tar.gz
	gorillacli_tar_url := "https://files.pythonhosted.org/packages/cd/2b/7a64f9ad59009e72ddf73d055195b4bf23e15599a61e66f1458b4025b9e5/gorilla-cli-0.0.10.tar.gz"
	gorillacli_cmd_tar := exec.Command("curl", "-L", gorillacli_tar_url, "-o", "package.tar.gz")
	err := gorillacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gorillacli_zip_url := "https://files.pythonhosted.org/packages/cd/2b/7a64f9ad59009e72ddf73d055195b4bf23e15599a61e66f1458b4025b9e5/gorilla-cli-0.0.10.zip"
	gorillacli_cmd_zip := exec.Command("curl", "-L", gorillacli_zip_url, "-o", "package.zip")
	err = gorillacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gorillacli_bin_url := "https://files.pythonhosted.org/packages/cd/2b/7a64f9ad59009e72ddf73d055195b4bf23e15599a61e66f1458b4025b9e5/gorilla-cli-0.0.10.bin"
	gorillacli_cmd_bin := exec.Command("curl", "-L", gorillacli_bin_url, "-o", "binary.bin")
	err = gorillacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gorillacli_src_url := "https://files.pythonhosted.org/packages/cd/2b/7a64f9ad59009e72ddf73d055195b4bf23e15599a61e66f1458b4025b9e5/gorilla-cli-0.0.10.src.tar.gz"
	gorillacli_cmd_src := exec.Command("curl", "-L", gorillacli_src_url, "-o", "source.tar.gz")
	err = gorillacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gorillacli_cmd_direct := exec.Command("./binary")
	err = gorillacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
