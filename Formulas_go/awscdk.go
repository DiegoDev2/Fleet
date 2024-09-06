package main

// AwsCdk - AWS Cloud Development Kit - framework for defining AWS infra as code
// Homepage: https://github.com/aws/aws-cdk

import (
	"fmt"
	
	"os/exec"
)

func installAwsCdk() {
	// Método 1: Descargar y extraer .tar.gz
	awscdk_tar_url := "https://registry.npmjs.org/aws-cdk/-/aws-cdk-2.155.0.tgz"
	awscdk_cmd_tar := exec.Command("curl", "-L", awscdk_tar_url, "-o", "package.tar.gz")
	err := awscdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awscdk_zip_url := "https://registry.npmjs.org/aws-cdk/-/aws-cdk-2.155.0.tgz"
	awscdk_cmd_zip := exec.Command("curl", "-L", awscdk_zip_url, "-o", "package.zip")
	err = awscdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awscdk_bin_url := "https://registry.npmjs.org/aws-cdk/-/aws-cdk-2.155.0.tgz"
	awscdk_cmd_bin := exec.Command("curl", "-L", awscdk_bin_url, "-o", "binary.bin")
	err = awscdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awscdk_src_url := "https://registry.npmjs.org/aws-cdk/-/aws-cdk-2.155.0.tgz"
	awscdk_cmd_src := exec.Command("curl", "-L", awscdk_src_url, "-o", "source.tar.gz")
	err = awscdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awscdk_cmd_direct := exec.Command("./binary")
	err = awscdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
