package main

// Localstack - Fully functional local AWS cloud stack
// Homepage: https://localstack.cloud/

import (
	"fmt"
	
	"os/exec"
)

func installLocalstack() {
	// Método 1: Descargar y extraer .tar.gz
	localstack_tar_url := "https://files.pythonhosted.org/packages/f4/4d/6f15ff84f6375a9e1d9d7484a8ffcfd91c71d17923bf816f192465917429/localstack-3.7.1.tar.gz"
	localstack_cmd_tar := exec.Command("curl", "-L", localstack_tar_url, "-o", "package.tar.gz")
	err := localstack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	localstack_zip_url := "https://files.pythonhosted.org/packages/f4/4d/6f15ff84f6375a9e1d9d7484a8ffcfd91c71d17923bf816f192465917429/localstack-3.7.1.zip"
	localstack_cmd_zip := exec.Command("curl", "-L", localstack_zip_url, "-o", "package.zip")
	err = localstack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	localstack_bin_url := "https://files.pythonhosted.org/packages/f4/4d/6f15ff84f6375a9e1d9d7484a8ffcfd91c71d17923bf816f192465917429/localstack-3.7.1.bin"
	localstack_cmd_bin := exec.Command("curl", "-L", localstack_bin_url, "-o", "binary.bin")
	err = localstack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	localstack_src_url := "https://files.pythonhosted.org/packages/f4/4d/6f15ff84f6375a9e1d9d7484a8ffcfd91c71d17923bf816f192465917429/localstack-3.7.1.src.tar.gz"
	localstack_cmd_src := exec.Command("curl", "-L", localstack_src_url, "-o", "source.tar.gz")
	err = localstack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	localstack_cmd_direct := exec.Command("./binary")
	err = localstack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docker")
exec.Command("latte", "install", "docker")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
