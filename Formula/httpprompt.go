package main

// HttpPrompt - Interactive command-line HTTP client with autocomplete and syntax highlighting
// Homepage: https://http-prompt.com

import (
	"fmt"
	
	"os/exec"
)

func installHttpPrompt() {
	// Método 1: Descargar y extraer .tar.gz
	httpprompt_tar_url := "https://files.pythonhosted.org/packages/bf/e2/bc5b0df107afcac65fde7015df48cbe9b4d877d1d0818203544ed1a41d4c/http-prompt-2.1.0.tar.gz"
	httpprompt_cmd_tar := exec.Command("curl", "-L", httpprompt_tar_url, "-o", "package.tar.gz")
	err := httpprompt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpprompt_zip_url := "https://files.pythonhosted.org/packages/bf/e2/bc5b0df107afcac65fde7015df48cbe9b4d877d1d0818203544ed1a41d4c/http-prompt-2.1.0.zip"
	httpprompt_cmd_zip := exec.Command("curl", "-L", httpprompt_zip_url, "-o", "package.zip")
	err = httpprompt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpprompt_bin_url := "https://files.pythonhosted.org/packages/bf/e2/bc5b0df107afcac65fde7015df48cbe9b4d877d1d0818203544ed1a41d4c/http-prompt-2.1.0.bin"
	httpprompt_cmd_bin := exec.Command("curl", "-L", httpprompt_bin_url, "-o", "binary.bin")
	err = httpprompt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpprompt_src_url := "https://files.pythonhosted.org/packages/bf/e2/bc5b0df107afcac65fde7015df48cbe9b4d877d1d0818203544ed1a41d4c/http-prompt-2.1.0.src.tar.gz"
	httpprompt_cmd_src := exec.Command("curl", "-L", httpprompt_src_url, "-o", "source.tar.gz")
	err = httpprompt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpprompt_cmd_direct := exec.Command("./binary")
	err = httpprompt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
