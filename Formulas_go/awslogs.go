package main

// Awslogs - Simple command-line tool to read AWS CloudWatch logs
// Homepage: https://github.com/jorgebastida/awslogs

import (
	"fmt"
	
	"os/exec"
)

func installAwslogs() {
	// Método 1: Descargar y extraer .tar.gz
	awslogs_tar_url := "https://files.pythonhosted.org/packages/15/f5/8f3bd0f4a927b1fbb3a5e6a5b036f29e4263977fb167b301bc4a5f4db2b9/awslogs-0.15.0.tar.gz"
	awslogs_cmd_tar := exec.Command("curl", "-L", awslogs_tar_url, "-o", "package.tar.gz")
	err := awslogs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awslogs_zip_url := "https://files.pythonhosted.org/packages/15/f5/8f3bd0f4a927b1fbb3a5e6a5b036f29e4263977fb167b301bc4a5f4db2b9/awslogs-0.15.0.zip"
	awslogs_cmd_zip := exec.Command("curl", "-L", awslogs_zip_url, "-o", "package.zip")
	err = awslogs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awslogs_bin_url := "https://files.pythonhosted.org/packages/15/f5/8f3bd0f4a927b1fbb3a5e6a5b036f29e4263977fb167b301bc4a5f4db2b9/awslogs-0.15.0.bin"
	awslogs_cmd_bin := exec.Command("curl", "-L", awslogs_bin_url, "-o", "binary.bin")
	err = awslogs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awslogs_src_url := "https://files.pythonhosted.org/packages/15/f5/8f3bd0f4a927b1fbb3a5e6a5b036f29e4263977fb167b301bc4a5f4db2b9/awslogs-0.15.0.src.tar.gz"
	awslogs_cmd_src := exec.Command("curl", "-L", awslogs_src_url, "-o", "source.tar.gz")
	err = awslogs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awslogs_cmd_direct := exec.Command("./binary")
	err = awslogs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
