package main

// EulerPy - Project Euler command-line tool written in Python
// Homepage: https://github.com/iKevinY/EulerPy

import (
	"fmt"
	
	"os/exec"
)

func installEulerPy() {
	// Método 1: Descargar y extraer .tar.gz
	eulerpy_tar_url := "https://files.pythonhosted.org/packages/a6/41/f074081bc036fbe2f066746e44020947ecf06ac53b6319a826023b8b5333/EulerPy-1.4.0.tar.gz"
	eulerpy_cmd_tar := exec.Command("curl", "-L", eulerpy_tar_url, "-o", "package.tar.gz")
	err := eulerpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eulerpy_zip_url := "https://files.pythonhosted.org/packages/a6/41/f074081bc036fbe2f066746e44020947ecf06ac53b6319a826023b8b5333/EulerPy-1.4.0.zip"
	eulerpy_cmd_zip := exec.Command("curl", "-L", eulerpy_zip_url, "-o", "package.zip")
	err = eulerpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eulerpy_bin_url := "https://files.pythonhosted.org/packages/a6/41/f074081bc036fbe2f066746e44020947ecf06ac53b6319a826023b8b5333/EulerPy-1.4.0.bin"
	eulerpy_cmd_bin := exec.Command("curl", "-L", eulerpy_bin_url, "-o", "binary.bin")
	err = eulerpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eulerpy_src_url := "https://files.pythonhosted.org/packages/a6/41/f074081bc036fbe2f066746e44020947ecf06ac53b6319a826023b8b5333/EulerPy-1.4.0.src.tar.gz"
	eulerpy_cmd_src := exec.Command("curl", "-L", eulerpy_src_url, "-o", "source.tar.gz")
	err = eulerpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eulerpy_cmd_direct := exec.Command("./binary")
	err = eulerpy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
