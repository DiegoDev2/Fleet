package main

// Poetry - Python package management tool
// Homepage: https://python-poetry.org/

import (
	"fmt"
	
	"os/exec"
)

func installPoetry() {
	// Método 1: Descargar y extraer .tar.gz
	poetry_tar_url := "https://files.pythonhosted.org/packages/07/c7/41108195c39ac1010054ef6b3b445894cee79e8ec73f086b73da94a01901/poetry-1.8.3.tar.gz"
	poetry_cmd_tar := exec.Command("curl", "-L", poetry_tar_url, "-o", "package.tar.gz")
	err := poetry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poetry_zip_url := "https://files.pythonhosted.org/packages/07/c7/41108195c39ac1010054ef6b3b445894cee79e8ec73f086b73da94a01901/poetry-1.8.3.zip"
	poetry_cmd_zip := exec.Command("curl", "-L", poetry_zip_url, "-o", "package.zip")
	err = poetry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poetry_bin_url := "https://files.pythonhosted.org/packages/07/c7/41108195c39ac1010054ef6b3b445894cee79e8ec73f086b73da94a01901/poetry-1.8.3.bin"
	poetry_cmd_bin := exec.Command("curl", "-L", poetry_bin_url, "-o", "binary.bin")
	err = poetry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poetry_src_url := "https://files.pythonhosted.org/packages/07/c7/41108195c39ac1010054ef6b3b445894cee79e8ec73f086b73da94a01901/poetry-1.8.3.src.tar.gz"
	poetry_cmd_src := exec.Command("curl", "-L", poetry_src_url, "-o", "source.tar.gz")
	err = poetry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poetry_cmd_direct := exec.Command("./binary")
	err = poetry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cffi")
exec.Command("latte", "install", "cffi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
}
