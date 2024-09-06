package main

// AwsGoogleAuth - Acquire AWS credentials using Google Apps
// Homepage: https://github.com/cevoaustralia/aws-google-auth

import (
	"fmt"
	
	"os/exec"
)

func installAwsGoogleAuth() {
	// Método 1: Descargar y extraer .tar.gz
	awsgoogleauth_tar_url := "https://files.pythonhosted.org/packages/32/4c/3a1dd1781c9d3bb4a85921b3d3e6e32fc0f0bad61ace6a8e1bd1a59c5ba0/aws-google-auth-0.0.38.tar.gz"
	awsgoogleauth_cmd_tar := exec.Command("curl", "-L", awsgoogleauth_tar_url, "-o", "package.tar.gz")
	err := awsgoogleauth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsgoogleauth_zip_url := "https://files.pythonhosted.org/packages/32/4c/3a1dd1781c9d3bb4a85921b3d3e6e32fc0f0bad61ace6a8e1bd1a59c5ba0/aws-google-auth-0.0.38.zip"
	awsgoogleauth_cmd_zip := exec.Command("curl", "-L", awsgoogleauth_zip_url, "-o", "package.zip")
	err = awsgoogleauth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsgoogleauth_bin_url := "https://files.pythonhosted.org/packages/32/4c/3a1dd1781c9d3bb4a85921b3d3e6e32fc0f0bad61ace6a8e1bd1a59c5ba0/aws-google-auth-0.0.38.bin"
	awsgoogleauth_cmd_bin := exec.Command("curl", "-L", awsgoogleauth_bin_url, "-o", "binary.bin")
	err = awsgoogleauth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsgoogleauth_src_url := "https://files.pythonhosted.org/packages/32/4c/3a1dd1781c9d3bb4a85921b3d3e6e32fc0f0bad61ace6a8e1bd1a59c5ba0/aws-google-auth-0.0.38.src.tar.gz"
	awsgoogleauth_cmd_src := exec.Command("curl", "-L", awsgoogleauth_src_url, "-o", "source.tar.gz")
	err = awsgoogleauth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsgoogleauth_cmd_direct := exec.Command("./binary")
	err = awsgoogleauth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
}
