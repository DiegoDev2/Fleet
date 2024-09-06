package main

// AwscliAT1 - Official Amazon AWS command-line interface
// Homepage: https://aws.amazon.com/cli/

import (
	"fmt"
	
	"os/exec"
)

func installAwscliAT1() {
	// Método 1: Descargar y extraer .tar.gz
	awscliat1_tar_url := "https://files.pythonhosted.org/packages/6e/26/3a2360d0fb24b12750d38e0852beda98288c090d565ab554a0fc6a7d6bdc/awscli-1.34.10.tar.gz"
	awscliat1_cmd_tar := exec.Command("curl", "-L", awscliat1_tar_url, "-o", "package.tar.gz")
	err := awscliat1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awscliat1_zip_url := "https://files.pythonhosted.org/packages/6e/26/3a2360d0fb24b12750d38e0852beda98288c090d565ab554a0fc6a7d6bdc/awscli-1.34.10.zip"
	awscliat1_cmd_zip := exec.Command("curl", "-L", awscliat1_zip_url, "-o", "package.zip")
	err = awscliat1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awscliat1_bin_url := "https://files.pythonhosted.org/packages/6e/26/3a2360d0fb24b12750d38e0852beda98288c090d565ab554a0fc6a7d6bdc/awscli-1.34.10.bin"
	awscliat1_cmd_bin := exec.Command("curl", "-L", awscliat1_bin_url, "-o", "binary.bin")
	err = awscliat1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awscliat1_src_url := "https://files.pythonhosted.org/packages/6e/26/3a2360d0fb24b12750d38e0852beda98288c090d565ab554a0fc6a7d6bdc/awscli-1.34.10.src.tar.gz"
	awscliat1_cmd_src := exec.Command("curl", "-L", awscliat1_src_url, "-o", "source.tar.gz")
	err = awscliat1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awscliat1_cmd_direct := exec.Command("./binary")
	err = awscliat1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
