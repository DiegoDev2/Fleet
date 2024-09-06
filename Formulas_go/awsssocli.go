package main

// AwsSsoCli - Securely manage AWS API credentials using AWS SSO
// Homepage: https://github.com/synfinatic/aws-sso-cli

import (
	"fmt"
	
	"os/exec"
)

func installAwsSsoCli() {
	// Método 1: Descargar y extraer .tar.gz
	awsssocli_tar_url := "https://github.com/synfinatic/aws-sso-cli/archive/refs/tags/v1.17.0.tar.gz"
	awsssocli_cmd_tar := exec.Command("curl", "-L", awsssocli_tar_url, "-o", "package.tar.gz")
	err := awsssocli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsssocli_zip_url := "https://github.com/synfinatic/aws-sso-cli/archive/refs/tags/v1.17.0.zip"
	awsssocli_cmd_zip := exec.Command("curl", "-L", awsssocli_zip_url, "-o", "package.zip")
	err = awsssocli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsssocli_bin_url := "https://github.com/synfinatic/aws-sso-cli/archive/refs/tags/v1.17.0.bin"
	awsssocli_cmd_bin := exec.Command("curl", "-L", awsssocli_bin_url, "-o", "binary.bin")
	err = awsssocli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsssocli_src_url := "https://github.com/synfinatic/aws-sso-cli/archive/refs/tags/v1.17.0.src.tar.gz"
	awsssocli_cmd_src := exec.Command("curl", "-L", awsssocli_src_url, "-o", "source.tar.gz")
	err = awsssocli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsssocli_cmd_direct := exec.Command("./binary")
	err = awsssocli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
