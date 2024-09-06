package main

// Yapf - Formatter for python code
// Homepage: https://github.com/google/yapf

import (
	"fmt"
	
	"os/exec"
)

func installYapf() {
	// Método 1: Descargar y extraer .tar.gz
	yapf_tar_url := "https://files.pythonhosted.org/packages/b9/14/c1f0ebd083fddd38a7c832d5ffde343150bd465689d12c549c303fbcd0f5/yapf-0.40.2.tar.gz"
	yapf_cmd_tar := exec.Command("curl", "-L", yapf_tar_url, "-o", "package.tar.gz")
	err := yapf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yapf_zip_url := "https://files.pythonhosted.org/packages/b9/14/c1f0ebd083fddd38a7c832d5ffde343150bd465689d12c549c303fbcd0f5/yapf-0.40.2.zip"
	yapf_cmd_zip := exec.Command("curl", "-L", yapf_zip_url, "-o", "package.zip")
	err = yapf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yapf_bin_url := "https://files.pythonhosted.org/packages/b9/14/c1f0ebd083fddd38a7c832d5ffde343150bd465689d12c549c303fbcd0f5/yapf-0.40.2.bin"
	yapf_cmd_bin := exec.Command("curl", "-L", yapf_bin_url, "-o", "binary.bin")
	err = yapf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yapf_src_url := "https://files.pythonhosted.org/packages/b9/14/c1f0ebd083fddd38a7c832d5ffde343150bd465689d12c549c303fbcd0f5/yapf-0.40.2.src.tar.gz"
	yapf_cmd_src := exec.Command("curl", "-L", yapf_src_url, "-o", "source.tar.gz")
	err = yapf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yapf_cmd_direct := exec.Command("./binary")
	err = yapf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
