package main

// Llm - Access large language models from the command-line
// Homepage: https://llm.datasette.io/

import (
	"fmt"
	
	"os/exec"
)

func installLlm() {
	// Método 1: Descargar y extraer .tar.gz
	llm_tar_url := "https://files.pythonhosted.org/packages/88/06/437d8ffaf329a05fedbccc994e0a9f57c1ea7c23a1c07e113b813a6652a6/llm-0.15.tar.gz"
	llm_cmd_tar := exec.Command("curl", "-L", llm_tar_url, "-o", "package.tar.gz")
	err := llm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	llm_zip_url := "https://files.pythonhosted.org/packages/88/06/437d8ffaf329a05fedbccc994e0a9f57c1ea7c23a1c07e113b813a6652a6/llm-0.15.zip"
	llm_cmd_zip := exec.Command("curl", "-L", llm_zip_url, "-o", "package.zip")
	err = llm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	llm_bin_url := "https://files.pythonhosted.org/packages/88/06/437d8ffaf329a05fedbccc994e0a9f57c1ea7c23a1c07e113b813a6652a6/llm-0.15.bin"
	llm_cmd_bin := exec.Command("curl", "-L", llm_bin_url, "-o", "binary.bin")
	err = llm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	llm_src_url := "https://files.pythonhosted.org/packages/88/06/437d8ffaf329a05fedbccc994e0a9f57c1ea7c23a1c07e113b813a6652a6/llm-0.15.src.tar.gz"
	llm_cmd_src := exec.Command("curl", "-L", llm_src_url, "-o", "source.tar.gz")
	err = llm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	llm_cmd_direct := exec.Command("./binary")
	err = llm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
