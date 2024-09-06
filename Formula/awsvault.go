package main

// AwsVault - Securely store and access AWS credentials in development environments
// Homepage: https://github.com/99designs/aws-vault

import (
	"fmt"
	
	"os/exec"
)

func installAwsVault() {
	// Método 1: Descargar y extraer .tar.gz
	awsvault_tar_url := "https://github.com/99designs/aws-vault/archive/refs/tags/v7.2.0.tar.gz"
	awsvault_cmd_tar := exec.Command("curl", "-L", awsvault_tar_url, "-o", "package.tar.gz")
	err := awsvault_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsvault_zip_url := "https://github.com/99designs/aws-vault/archive/refs/tags/v7.2.0.zip"
	awsvault_cmd_zip := exec.Command("curl", "-L", awsvault_zip_url, "-o", "package.zip")
	err = awsvault_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsvault_bin_url := "https://github.com/99designs/aws-vault/archive/refs/tags/v7.2.0.bin"
	awsvault_cmd_bin := exec.Command("curl", "-L", awsvault_bin_url, "-o", "binary.bin")
	err = awsvault_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsvault_src_url := "https://github.com/99designs/aws-vault/archive/refs/tags/v7.2.0.src.tar.gz"
	awsvault_cmd_src := exec.Command("curl", "-L", awsvault_src_url, "-o", "source.tar.gz")
	err = awsvault_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsvault_cmd_direct := exec.Command("./binary")
	err = awsvault_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
