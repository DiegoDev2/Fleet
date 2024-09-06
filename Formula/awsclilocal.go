package main

// AwscliLocal - Thin wrapper around the `aws` command-line interface for use with LocalStack
// Homepage: https://localstack.cloud/

import (
	"fmt"
	
	"os/exec"
)

func installAwscliLocal() {
	// Método 1: Descargar y extraer .tar.gz
	awsclilocal_tar_url := "https://files.pythonhosted.org/packages/25/f9/023c80ea27d67b0930f116597fd55a93f84de9b05d18b38c7d2d5d75c1c9/awscli-local-0.22.0.tar.gz"
	awsclilocal_cmd_tar := exec.Command("curl", "-L", awsclilocal_tar_url, "-o", "package.tar.gz")
	err := awsclilocal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsclilocal_zip_url := "https://files.pythonhosted.org/packages/25/f9/023c80ea27d67b0930f116597fd55a93f84de9b05d18b38c7d2d5d75c1c9/awscli-local-0.22.0.zip"
	awsclilocal_cmd_zip := exec.Command("curl", "-L", awsclilocal_zip_url, "-o", "package.zip")
	err = awsclilocal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsclilocal_bin_url := "https://files.pythonhosted.org/packages/25/f9/023c80ea27d67b0930f116597fd55a93f84de9b05d18b38c7d2d5d75c1c9/awscli-local-0.22.0.bin"
	awsclilocal_cmd_bin := exec.Command("curl", "-L", awsclilocal_bin_url, "-o", "binary.bin")
	err = awsclilocal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsclilocal_src_url := "https://files.pythonhosted.org/packages/25/f9/023c80ea27d67b0930f116597fd55a93f84de9b05d18b38c7d2d5d75c1c9/awscli-local-0.22.0.src.tar.gz"
	awsclilocal_cmd_src := exec.Command("curl", "-L", awsclilocal_src_url, "-o", "source.tar.gz")
	err = awsclilocal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsclilocal_cmd_direct := exec.Command("./binary")
	err = awsclilocal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: awscli")
	exec.Command("latte", "install", "awscli").Run()
	fmt.Println("Instalando dependencia: localstack")
	exec.Command("latte", "install", "localstack").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
