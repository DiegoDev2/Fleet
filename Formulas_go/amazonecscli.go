package main

// AmazonEcsCli - CLI for Amazon ECS to manage clusters and tasks for development
// Homepage: https://aws.amazon.com/ecs

import (
	"fmt"
	
	"os/exec"
)

func installAmazonEcsCli() {
	// Método 1: Descargar y extraer .tar.gz
	amazonecscli_tar_url := "https://github.com/aws/amazon-ecs-cli/archive/refs/tags/v1.21.0.tar.gz"
	amazonecscli_cmd_tar := exec.Command("curl", "-L", amazonecscli_tar_url, "-o", "package.tar.gz")
	err := amazonecscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amazonecscli_zip_url := "https://github.com/aws/amazon-ecs-cli/archive/refs/tags/v1.21.0.zip"
	amazonecscli_cmd_zip := exec.Command("curl", "-L", amazonecscli_zip_url, "-o", "package.zip")
	err = amazonecscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amazonecscli_bin_url := "https://github.com/aws/amazon-ecs-cli/archive/refs/tags/v1.21.0.bin"
	amazonecscli_cmd_bin := exec.Command("curl", "-L", amazonecscli_bin_url, "-o", "binary.bin")
	err = amazonecscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amazonecscli_src_url := "https://github.com/aws/amazon-ecs-cli/archive/refs/tags/v1.21.0.src.tar.gz"
	amazonecscli_cmd_src := exec.Command("curl", "-L", amazonecscli_src_url, "-o", "source.tar.gz")
	err = amazonecscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amazonecscli_cmd_direct := exec.Command("./binary")
	err = amazonecscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
