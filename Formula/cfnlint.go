package main

// CfnLint - Validate CloudFormation templates against the CloudFormation spec
// Homepage: https://github.com/aws-cloudformation/cfn-lint/

import (
	"fmt"
	
	"os/exec"
)

func installCfnLint() {
	// Método 1: Descargar y extraer .tar.gz
	cfnlint_tar_url := "https://files.pythonhosted.org/packages/ae/e9/c4bba9c23e3cc1584fd1896f284eac42464e3a832588be619fdd28ecc3d5/cfn_lint-1.12.3.tar.gz"
	cfnlint_cmd_tar := exec.Command("curl", "-L", cfnlint_tar_url, "-o", "package.tar.gz")
	err := cfnlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfnlint_zip_url := "https://files.pythonhosted.org/packages/ae/e9/c4bba9c23e3cc1584fd1896f284eac42464e3a832588be619fdd28ecc3d5/cfn_lint-1.12.3.zip"
	cfnlint_cmd_zip := exec.Command("curl", "-L", cfnlint_zip_url, "-o", "package.zip")
	err = cfnlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfnlint_bin_url := "https://files.pythonhosted.org/packages/ae/e9/c4bba9c23e3cc1584fd1896f284eac42464e3a832588be619fdd28ecc3d5/cfn_lint-1.12.3.bin"
	cfnlint_cmd_bin := exec.Command("curl", "-L", cfnlint_bin_url, "-o", "binary.bin")
	err = cfnlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfnlint_src_url := "https://files.pythonhosted.org/packages/ae/e9/c4bba9c23e3cc1584fd1896f284eac42464e3a832588be619fdd28ecc3d5/cfn_lint-1.12.3.src.tar.gz"
	cfnlint_cmd_src := exec.Command("curl", "-L", cfnlint_src_url, "-o", "source.tar.gz")
	err = cfnlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfnlint_cmd_direct := exec.Command("./binary")
	err = cfnlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
