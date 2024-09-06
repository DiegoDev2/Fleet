package main

// Mathlibtools - Lean prover mathlib supporting tools
// Homepage: https://pypi.org/project/mathlibtools

import (
	"fmt"
	
	"os/exec"
)

func installMathlibtools() {
	// Método 1: Descargar y extraer .tar.gz
	mathlibtools_tar_url := "https://files.pythonhosted.org/packages/ae/6a/815d7f65dc853973b13be082fefe797074e633407ef1262a62bc0be84203/mathlibtools-1.3.2.tar.gz"
	mathlibtools_cmd_tar := exec.Command("curl", "-L", mathlibtools_tar_url, "-o", "package.tar.gz")
	err := mathlibtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mathlibtools_zip_url := "https://files.pythonhosted.org/packages/ae/6a/815d7f65dc853973b13be082fefe797074e633407ef1262a62bc0be84203/mathlibtools-1.3.2.zip"
	mathlibtools_cmd_zip := exec.Command("curl", "-L", mathlibtools_zip_url, "-o", "package.zip")
	err = mathlibtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mathlibtools_bin_url := "https://files.pythonhosted.org/packages/ae/6a/815d7f65dc853973b13be082fefe797074e633407ef1262a62bc0be84203/mathlibtools-1.3.2.bin"
	mathlibtools_cmd_bin := exec.Command("curl", "-L", mathlibtools_bin_url, "-o", "binary.bin")
	err = mathlibtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mathlibtools_src_url := "https://files.pythonhosted.org/packages/ae/6a/815d7f65dc853973b13be082fefe797074e633407ef1262a62bc0be84203/mathlibtools-1.3.2.src.tar.gz"
	mathlibtools_cmd_src := exec.Command("curl", "-L", mathlibtools_src_url, "-o", "source.tar.gz")
	err = mathlibtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mathlibtools_cmd_direct := exec.Command("./binary")
	err = mathlibtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
