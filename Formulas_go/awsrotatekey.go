package main

// AwsRotateKey - Easily rotate your AWS access key
// Homepage: https://github.com/stefansundin/aws-rotate-key

import (
	"fmt"
	
	"os/exec"
)

func installAwsRotateKey() {
	// Método 1: Descargar y extraer .tar.gz
	awsrotatekey_tar_url := "https://github.com/stefansundin/aws-rotate-key/archive/refs/tags/v1.1.0.tar.gz"
	awsrotatekey_cmd_tar := exec.Command("curl", "-L", awsrotatekey_tar_url, "-o", "package.tar.gz")
	err := awsrotatekey_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsrotatekey_zip_url := "https://github.com/stefansundin/aws-rotate-key/archive/refs/tags/v1.1.0.zip"
	awsrotatekey_cmd_zip := exec.Command("curl", "-L", awsrotatekey_zip_url, "-o", "package.zip")
	err = awsrotatekey_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsrotatekey_bin_url := "https://github.com/stefansundin/aws-rotate-key/archive/refs/tags/v1.1.0.bin"
	awsrotatekey_cmd_bin := exec.Command("curl", "-L", awsrotatekey_bin_url, "-o", "binary.bin")
	err = awsrotatekey_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsrotatekey_src_url := "https://github.com/stefansundin/aws-rotate-key/archive/refs/tags/v1.1.0.src.tar.gz"
	awsrotatekey_cmd_src := exec.Command("curl", "-L", awsrotatekey_src_url, "-o", "source.tar.gz")
	err = awsrotatekey_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsrotatekey_cmd_direct := exec.Command("./binary")
	err = awsrotatekey_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
