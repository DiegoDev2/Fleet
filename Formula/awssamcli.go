package main

// AwsSamCli - CLI tool to build, test, debug, and deploy Serverless applications using AWS SAM
// Homepage: https://aws.amazon.com/serverless/sam/

import (
	"fmt"
	
	"os/exec"
)

func installAwsSamCli() {
	// Método 1: Descargar y extraer .tar.gz
	awssamcli_tar_url := "https://files.pythonhosted.org/packages/67/5d/0a5159e7a752a6d969a22ce48213ddf56231c38f5835dc9e411772c0f777/aws_sam_cli-1.123.0.tar.gz"
	awssamcli_cmd_tar := exec.Command("curl", "-L", awssamcli_tar_url, "-o", "package.tar.gz")
	err := awssamcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awssamcli_zip_url := "https://files.pythonhosted.org/packages/67/5d/0a5159e7a752a6d969a22ce48213ddf56231c38f5835dc9e411772c0f777/aws_sam_cli-1.123.0.zip"
	awssamcli_cmd_zip := exec.Command("curl", "-L", awssamcli_zip_url, "-o", "package.zip")
	err = awssamcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awssamcli_bin_url := "https://files.pythonhosted.org/packages/67/5d/0a5159e7a752a6d969a22ce48213ddf56231c38f5835dc9e411772c0f777/aws_sam_cli-1.123.0.bin"
	awssamcli_cmd_bin := exec.Command("curl", "-L", awssamcli_bin_url, "-o", "binary.bin")
	err = awssamcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awssamcli_src_url := "https://files.pythonhosted.org/packages/67/5d/0a5159e7a752a6d969a22ce48213ddf56231c38f5835dc9e411772c0f777/aws_sam_cli-1.123.0.src.tar.gz"
	awssamcli_cmd_src := exec.Command("curl", "-L", awssamcli_src_url, "-o", "source.tar.gz")
	err = awssamcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awssamcli_cmd_direct := exec.Command("./binary")
	err = awssamcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
