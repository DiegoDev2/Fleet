package main

// CloudformationCli - CloudFormation Provider Development Toolkit
// Homepage: https://github.com/aws-cloudformation/cloudformation-cli/

import (
	"fmt"
	
	"os/exec"
)

func installCloudformationCli() {
	// Método 1: Descargar y extraer .tar.gz
	cloudformationcli_tar_url := "https://files.pythonhosted.org/packages/51/46/e1a0bad01e7e1f14b7e3190d73818c4a903f44aeb228dfcabd18d197a0bf/cloudformation-cli-0.2.38.tar.gz"
	cloudformationcli_cmd_tar := exec.Command("curl", "-L", cloudformationcli_tar_url, "-o", "package.tar.gz")
	err := cloudformationcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudformationcli_zip_url := "https://files.pythonhosted.org/packages/51/46/e1a0bad01e7e1f14b7e3190d73818c4a903f44aeb228dfcabd18d197a0bf/cloudformation-cli-0.2.38.zip"
	cloudformationcli_cmd_zip := exec.Command("curl", "-L", cloudformationcli_zip_url, "-o", "package.zip")
	err = cloudformationcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudformationcli_bin_url := "https://files.pythonhosted.org/packages/51/46/e1a0bad01e7e1f14b7e3190d73818c4a903f44aeb228dfcabd18d197a0bf/cloudformation-cli-0.2.38.bin"
	cloudformationcli_cmd_bin := exec.Command("curl", "-L", cloudformationcli_bin_url, "-o", "binary.bin")
	err = cloudformationcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudformationcli_src_url := "https://files.pythonhosted.org/packages/51/46/e1a0bad01e7e1f14b7e3190d73818c4a903f44aeb228dfcabd18d197a0bf/cloudformation-cli-0.2.38.src.tar.gz"
	cloudformationcli_cmd_src := exec.Command("curl", "-L", cloudformationcli_src_url, "-o", "source.tar.gz")
	err = cloudformationcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudformationcli_cmd_direct := exec.Command("./binary")
	err = cloudformationcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
