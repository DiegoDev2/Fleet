package main

// CfnFormat - Command-line tool for formatting AWS CloudFormation templates
// Homepage: https://github.com/aws-cloudformation/rain

import (
	"fmt"
	
	"os/exec"
)

func installCfnFormat() {
	// Método 1: Descargar y extraer .tar.gz
	cfnformat_tar_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.tar.gz"
	cfnformat_cmd_tar := exec.Command("curl", "-L", cfnformat_tar_url, "-o", "package.tar.gz")
	err := cfnformat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfnformat_zip_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.zip"
	cfnformat_cmd_zip := exec.Command("curl", "-L", cfnformat_zip_url, "-o", "package.zip")
	err = cfnformat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfnformat_bin_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.bin"
	cfnformat_cmd_bin := exec.Command("curl", "-L", cfnformat_bin_url, "-o", "binary.bin")
	err = cfnformat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfnformat_src_url := "https://github.com/aws-cloudformation/rain/archive/refs/tags/v1.15.0.src.tar.gz"
	cfnformat_cmd_src := exec.Command("curl", "-L", cfnformat_src_url, "-o", "source.tar.gz")
	err = cfnformat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfnformat_cmd_direct := exec.Command("./binary")
	err = cfnformat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
