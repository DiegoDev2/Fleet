package main

// Athenacli - CLI tool for AWS Athena service
// Homepage: https://athenacli.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installAthenacli() {
	// Método 1: Descargar y extraer .tar.gz
	athenacli_tar_url := "https://files.pythonhosted.org/packages/38/1a/d9cd6c68a4a1cd2ce779b163f8cec390ae82c684caa920d0360094886b1f/athenacli-1.6.8.tar.gz"
	athenacli_cmd_tar := exec.Command("curl", "-L", athenacli_tar_url, "-o", "package.tar.gz")
	err := athenacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	athenacli_zip_url := "https://files.pythonhosted.org/packages/38/1a/d9cd6c68a4a1cd2ce779b163f8cec390ae82c684caa920d0360094886b1f/athenacli-1.6.8.zip"
	athenacli_cmd_zip := exec.Command("curl", "-L", athenacli_zip_url, "-o", "package.zip")
	err = athenacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	athenacli_bin_url := "https://files.pythonhosted.org/packages/38/1a/d9cd6c68a4a1cd2ce779b163f8cec390ae82c684caa920d0360094886b1f/athenacli-1.6.8.bin"
	athenacli_cmd_bin := exec.Command("curl", "-L", athenacli_bin_url, "-o", "binary.bin")
	err = athenacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	athenacli_src_url := "https://files.pythonhosted.org/packages/38/1a/d9cd6c68a4a1cd2ce779b163f8cec390ae82c684caa920d0360094886b1f/athenacli-1.6.8.src.tar.gz"
	athenacli_cmd_src := exec.Command("curl", "-L", athenacli_src_url, "-o", "source.tar.gz")
	err = athenacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	athenacli_cmd_direct := exec.Command("./binary")
	err = athenacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
