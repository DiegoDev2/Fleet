package main

// CfnFlip - Convert AWS CloudFormation templates between JSON and YAML formats
// Homepage: https://github.com/awslabs/aws-cfn-template-flip

import (
	"fmt"
	
	"os/exec"
)

func installCfnFlip() {
	// Método 1: Descargar y extraer .tar.gz
	cfnflip_tar_url := "https://files.pythonhosted.org/packages/ca/75/8eba0bb52a6c58e347bc4c839b249d9f42380de93ed12a14eba4355387b4/cfn_flip-1.3.0.tar.gz"
	cfnflip_cmd_tar := exec.Command("curl", "-L", cfnflip_tar_url, "-o", "package.tar.gz")
	err := cfnflip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfnflip_zip_url := "https://files.pythonhosted.org/packages/ca/75/8eba0bb52a6c58e347bc4c839b249d9f42380de93ed12a14eba4355387b4/cfn_flip-1.3.0.zip"
	cfnflip_cmd_zip := exec.Command("curl", "-L", cfnflip_zip_url, "-o", "package.zip")
	err = cfnflip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfnflip_bin_url := "https://files.pythonhosted.org/packages/ca/75/8eba0bb52a6c58e347bc4c839b249d9f42380de93ed12a14eba4355387b4/cfn_flip-1.3.0.bin"
	cfnflip_cmd_bin := exec.Command("curl", "-L", cfnflip_bin_url, "-o", "binary.bin")
	err = cfnflip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfnflip_src_url := "https://files.pythonhosted.org/packages/ca/75/8eba0bb52a6c58e347bc4c839b249d9f42380de93ed12a14eba4355387b4/cfn_flip-1.3.0.src.tar.gz"
	cfnflip_cmd_src := exec.Command("curl", "-L", cfnflip_src_url, "-o", "source.tar.gz")
	err = cfnflip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfnflip_cmd_direct := exec.Command("./binary")
	err = cfnflip_cmd_direct.Run()
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
