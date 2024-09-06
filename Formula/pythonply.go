package main

// PythonPly - Python Lex & Yacc
// Homepage: https://github.com/dabeaz/ply

import (
	"fmt"
	
	"os/exec"
)

func installPythonPly() {
	// Método 1: Descargar y extraer .tar.gz
	pythonply_tar_url := "https://files.pythonhosted.org/packages/e5/69/882ee5c9d017149285cab114ebeab373308ef0f874fcdac9beb90e0ac4da/ply-3.11.tar.gz"
	pythonply_cmd_tar := exec.Command("curl", "-L", pythonply_tar_url, "-o", "package.tar.gz")
	err := pythonply_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonply_zip_url := "https://files.pythonhosted.org/packages/e5/69/882ee5c9d017149285cab114ebeab373308ef0f874fcdac9beb90e0ac4da/ply-3.11.zip"
	pythonply_cmd_zip := exec.Command("curl", "-L", pythonply_zip_url, "-o", "package.zip")
	err = pythonply_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonply_bin_url := "https://files.pythonhosted.org/packages/e5/69/882ee5c9d017149285cab114ebeab373308ef0f874fcdac9beb90e0ac4da/ply-3.11.bin"
	pythonply_cmd_bin := exec.Command("curl", "-L", pythonply_bin_url, "-o", "binary.bin")
	err = pythonply_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonply_src_url := "https://files.pythonhosted.org/packages/e5/69/882ee5c9d017149285cab114ebeab373308ef0f874fcdac9beb90e0ac4da/ply-3.11.src.tar.gz"
	pythonply_cmd_src := exec.Command("curl", "-L", pythonply_src_url, "-o", "source.tar.gz")
	err = pythonply_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonply_cmd_direct := exec.Command("./binary")
	err = pythonply_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
