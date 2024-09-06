package main

// Awscli - Official Amazon AWS command-line interface
// Homepage: https://aws.amazon.com/cli/

import (
	"fmt"
	
	"os/exec"
)

func installAwscli() {
	// Método 1: Descargar y extraer .tar.gz
	awscli_tar_url := "https://github.com/aws/aws-cli/archive/refs/tags/2.17.45.tar.gz"
	awscli_cmd_tar := exec.Command("curl", "-L", awscli_tar_url, "-o", "package.tar.gz")
	err := awscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awscli_zip_url := "https://github.com/aws/aws-cli/archive/refs/tags/2.17.45.zip"
	awscli_cmd_zip := exec.Command("curl", "-L", awscli_zip_url, "-o", "package.zip")
	err = awscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awscli_bin_url := "https://github.com/aws/aws-cli/archive/refs/tags/2.17.45.bin"
	awscli_cmd_bin := exec.Command("curl", "-L", awscli_bin_url, "-o", "binary.bin")
	err = awscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awscli_src_url := "https://github.com/aws/aws-cli/archive/refs/tags/2.17.45.src.tar.gz"
	awscli_cmd_src := exec.Command("curl", "-L", awscli_src_url, "-o", "source.tar.gz")
	err = awscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awscli_cmd_direct := exec.Command("./binary")
	err = awscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
}
