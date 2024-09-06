package main

// Pdm - Modern Python package and dependency manager supporting the latest PEP standards
// Homepage: https://pdm.fming.dev

import (
	"fmt"
	
	"os/exec"
)

func installPdm() {
	// Método 1: Descargar y extraer .tar.gz
	pdm_tar_url := "https://files.pythonhosted.org/packages/a7/4f/6636a4aca3293294d26a240b618ea9813d1214ed77d093f92ed56ca8ebf2/pdm-2.18.1.tar.gz"
	pdm_cmd_tar := exec.Command("curl", "-L", pdm_tar_url, "-o", "package.tar.gz")
	err := pdm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdm_zip_url := "https://files.pythonhosted.org/packages/a7/4f/6636a4aca3293294d26a240b618ea9813d1214ed77d093f92ed56ca8ebf2/pdm-2.18.1.zip"
	pdm_cmd_zip := exec.Command("curl", "-L", pdm_zip_url, "-o", "package.zip")
	err = pdm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdm_bin_url := "https://files.pythonhosted.org/packages/a7/4f/6636a4aca3293294d26a240b618ea9813d1214ed77d093f92ed56ca8ebf2/pdm-2.18.1.bin"
	pdm_cmd_bin := exec.Command("curl", "-L", pdm_bin_url, "-o", "binary.bin")
	err = pdm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdm_src_url := "https://files.pythonhosted.org/packages/a7/4f/6636a4aca3293294d26a240b618ea9813d1214ed77d093f92ed56ca8ebf2/pdm-2.18.1.src.tar.gz"
	pdm_cmd_src := exec.Command("curl", "-L", pdm_src_url, "-o", "source.tar.gz")
	err = pdm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdm_cmd_direct := exec.Command("./binary")
	err = pdm_cmd_direct.Run()
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
