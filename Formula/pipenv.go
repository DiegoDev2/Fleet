package main

// Pipenv - Python dependency management tool
// Homepage: https://github.com/pypa/pipenv

import (
	"fmt"
	
	"os/exec"
)

func installPipenv() {
	// Método 1: Descargar y extraer .tar.gz
	pipenv_tar_url := "https://files.pythonhosted.org/packages/d1/67/c29cb9081e5648b754b7ec95482e348b4d616681a3f0ee402ca082b9be02/pipenv-2024.0.1.tar.gz"
	pipenv_cmd_tar := exec.Command("curl", "-L", pipenv_tar_url, "-o", "package.tar.gz")
	err := pipenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipenv_zip_url := "https://files.pythonhosted.org/packages/d1/67/c29cb9081e5648b754b7ec95482e348b4d616681a3f0ee402ca082b9be02/pipenv-2024.0.1.zip"
	pipenv_cmd_zip := exec.Command("curl", "-L", pipenv_zip_url, "-o", "package.zip")
	err = pipenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipenv_bin_url := "https://files.pythonhosted.org/packages/d1/67/c29cb9081e5648b754b7ec95482e348b4d616681a3f0ee402ca082b9be02/pipenv-2024.0.1.bin"
	pipenv_cmd_bin := exec.Command("curl", "-L", pipenv_bin_url, "-o", "binary.bin")
	err = pipenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipenv_src_url := "https://files.pythonhosted.org/packages/d1/67/c29cb9081e5648b754b7ec95482e348b4d616681a3f0ee402ca082b9be02/pipenv-2024.0.1.src.tar.gz"
	pipenv_cmd_src := exec.Command("curl", "-L", pipenv_src_url, "-o", "source.tar.gz")
	err = pipenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipenv_cmd_direct := exec.Command("./binary")
	err = pipenv_cmd_direct.Run()
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
