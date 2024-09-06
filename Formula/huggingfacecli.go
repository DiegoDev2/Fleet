package main

// HuggingfaceCli - Client library for huggingface.co hub
// Homepage: https://huggingface.co/docs/huggingface_hub/index

import (
	"fmt"
	
	"os/exec"
)

func installHuggingfaceCli() {
	// Método 1: Descargar y extraer .tar.gz
	huggingfacecli_tar_url := "https://files.pythonhosted.org/packages/65/24/b98fce967b7d63700e5805b915012ba25bb538a81fcf11e97f3cc3f4f012/huggingface_hub-0.24.6.tar.gz"
	huggingfacecli_cmd_tar := exec.Command("curl", "-L", huggingfacecli_tar_url, "-o", "package.tar.gz")
	err := huggingfacecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	huggingfacecli_zip_url := "https://files.pythonhosted.org/packages/65/24/b98fce967b7d63700e5805b915012ba25bb538a81fcf11e97f3cc3f4f012/huggingface_hub-0.24.6.zip"
	huggingfacecli_cmd_zip := exec.Command("curl", "-L", huggingfacecli_zip_url, "-o", "package.zip")
	err = huggingfacecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	huggingfacecli_bin_url := "https://files.pythonhosted.org/packages/65/24/b98fce967b7d63700e5805b915012ba25bb538a81fcf11e97f3cc3f4f012/huggingface_hub-0.24.6.bin"
	huggingfacecli_cmd_bin := exec.Command("curl", "-L", huggingfacecli_bin_url, "-o", "binary.bin")
	err = huggingfacecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	huggingfacecli_src_url := "https://files.pythonhosted.org/packages/65/24/b98fce967b7d63700e5805b915012ba25bb538a81fcf11e97f3cc3f4f012/huggingface_hub-0.24.6.src.tar.gz"
	huggingfacecli_cmd_src := exec.Command("curl", "-L", huggingfacecli_src_url, "-o", "source.tar.gz")
	err = huggingfacecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	huggingfacecli_cmd_direct := exec.Command("./binary")
	err = huggingfacecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: git-lfs")
	exec.Command("latte", "install", "git-lfs").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
