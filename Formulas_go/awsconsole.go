package main

// AwsConsole - Command-line to use AWS CLI credentials to launch the AWS console in a browser
// Homepage: https://github.com/aws-cloudformation/rain

import (
	"fmt"
	
	"os/exec"
)

func installAwsConsole() {
	// Método 1: Descargar y extraer .tar.gz
	awsconsole_tar_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.tar.gz"
	awsconsole_cmd_tar := exec.Command("curl", "-L", awsconsole_tar_url, "-o", "package.tar.gz")
	err := awsconsole_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsconsole_zip_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.zip"
	awsconsole_cmd_zip := exec.Command("curl", "-L", awsconsole_zip_url, "-o", "package.zip")
	err = awsconsole_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsconsole_bin_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.bin"
	awsconsole_cmd_bin := exec.Command("curl", "-L", awsconsole_bin_url, "-o", "binary.bin")
	err = awsconsole_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsconsole_src_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.src.tar.gz"
	awsconsole_cmd_src := exec.Command("curl", "-L", awsconsole_src_url, "-o", "source.tar.gz")
	err = awsconsole_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsconsole_cmd_direct := exec.Command("./binary")
	err = awsconsole_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
