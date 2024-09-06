package main

// Twine - Utilities for interacting with PyPI
// Homepage: https://github.com/pypa/twine

import (
	"fmt"
	
	"os/exec"
)

func installTwine() {
	// Método 1: Descargar y extraer .tar.gz
	twine_tar_url := "https://files.pythonhosted.org/packages/77/68/bd982e5e949ef8334e6f7dcf76ae40922a8750aa2e347291ae1477a4782b/twine-5.1.1.tar.gz"
	twine_cmd_tar := exec.Command("curl", "-L", twine_tar_url, "-o", "package.tar.gz")
	err := twine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twine_zip_url := "https://files.pythonhosted.org/packages/77/68/bd982e5e949ef8334e6f7dcf76ae40922a8750aa2e347291ae1477a4782b/twine-5.1.1.zip"
	twine_cmd_zip := exec.Command("curl", "-L", twine_zip_url, "-o", "package.zip")
	err = twine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twine_bin_url := "https://files.pythonhosted.org/packages/77/68/bd982e5e949ef8334e6f7dcf76ae40922a8750aa2e347291ae1477a4782b/twine-5.1.1.bin"
	twine_cmd_bin := exec.Command("curl", "-L", twine_bin_url, "-o", "binary.bin")
	err = twine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twine_src_url := "https://files.pythonhosted.org/packages/77/68/bd982e5e949ef8334e6f7dcf76ae40922a8750aa2e347291ae1477a4782b/twine-5.1.1.src.tar.gz"
	twine_cmd_src := exec.Command("curl", "-L", twine_src_url, "-o", "source.tar.gz")
	err = twine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twine_cmd_direct := exec.Command("./binary")
	err = twine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
