package main

// Cpplint - Static code checker for C++
// Homepage: https://pypi.org/project/cpplint/

import (
	"fmt"
	
	"os/exec"
)

func installCpplint() {
	// Método 1: Descargar y extraer .tar.gz
	cpplint_tar_url := "https://files.pythonhosted.org/packages/18/72/ea0f4035bcf35d8f8df053657d7f3370d56ff4d4e6617021b6544b9958d4/cpplint-1.6.1.tar.gz"
	cpplint_cmd_tar := exec.Command("curl", "-L", cpplint_tar_url, "-o", "package.tar.gz")
	err := cpplint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpplint_zip_url := "https://files.pythonhosted.org/packages/18/72/ea0f4035bcf35d8f8df053657d7f3370d56ff4d4e6617021b6544b9958d4/cpplint-1.6.1.zip"
	cpplint_cmd_zip := exec.Command("curl", "-L", cpplint_zip_url, "-o", "package.zip")
	err = cpplint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpplint_bin_url := "https://files.pythonhosted.org/packages/18/72/ea0f4035bcf35d8f8df053657d7f3370d56ff4d4e6617021b6544b9958d4/cpplint-1.6.1.bin"
	cpplint_cmd_bin := exec.Command("curl", "-L", cpplint_bin_url, "-o", "binary.bin")
	err = cpplint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpplint_src_url := "https://files.pythonhosted.org/packages/18/72/ea0f4035bcf35d8f8df053657d7f3370d56ff4d4e6617021b6544b9958d4/cpplint-1.6.1.src.tar.gz"
	cpplint_cmd_src := exec.Command("curl", "-L", cpplint_src_url, "-o", "source.tar.gz")
	err = cpplint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpplint_cmd_direct := exec.Command("./binary")
	err = cpplint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
