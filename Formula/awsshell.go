package main

// AwsShell - Integrated shell for working with the AWS CLI
// Homepage: https://github.com/awslabs/aws-shell

import (
	"fmt"
	
	"os/exec"
)

func installAwsShell() {
	// Método 1: Descargar y extraer .tar.gz
	awsshell_tar_url := "https://files.pythonhosted.org/packages/01/31/ee166a91c865a855af4f15e393974eadf57762629fc2a163a3eb3f470ac5/aws-shell-0.2.2.tar.gz"
	awsshell_cmd_tar := exec.Command("curl", "-L", awsshell_tar_url, "-o", "package.tar.gz")
	err := awsshell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsshell_zip_url := "https://files.pythonhosted.org/packages/01/31/ee166a91c865a855af4f15e393974eadf57762629fc2a163a3eb3f470ac5/aws-shell-0.2.2.zip"
	awsshell_cmd_zip := exec.Command("curl", "-L", awsshell_zip_url, "-o", "package.zip")
	err = awsshell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsshell_bin_url := "https://files.pythonhosted.org/packages/01/31/ee166a91c865a855af4f15e393974eadf57762629fc2a163a3eb3f470ac5/aws-shell-0.2.2.bin"
	awsshell_cmd_bin := exec.Command("curl", "-L", awsshell_bin_url, "-o", "binary.bin")
	err = awsshell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsshell_src_url := "https://files.pythonhosted.org/packages/01/31/ee166a91c865a855af4f15e393974eadf57762629fc2a163a3eb3f470ac5/aws-shell-0.2.2.src.tar.gz"
	awsshell_cmd_src := exec.Command("curl", "-L", awsshell_src_url, "-o", "source.tar.gz")
	err = awsshell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsshell_cmd_direct := exec.Command("./binary")
	err = awsshell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
