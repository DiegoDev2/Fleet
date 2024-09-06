package main

// DockerCredentialHelperEcr - Docker Credential Helper for Amazon ECR
// Homepage: https://github.com/awslabs/amazon-ecr-credential-helper

import (
	"fmt"
	
	"os/exec"
)

func installDockerCredentialHelperEcr() {
	// Método 1: Descargar y extraer .tar.gz
	dockercredentialhelperecr_tar_url := "https://github.com/awslabs/amazon-ecr-credential-helper/archive/refs/tags/v0.8.0.tar.gz"
	dockercredentialhelperecr_cmd_tar := exec.Command("curl", "-L", dockercredentialhelperecr_tar_url, "-o", "package.tar.gz")
	err := dockercredentialhelperecr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockercredentialhelperecr_zip_url := "https://github.com/awslabs/amazon-ecr-credential-helper/archive/refs/tags/v0.8.0.zip"
	dockercredentialhelperecr_cmd_zip := exec.Command("curl", "-L", dockercredentialhelperecr_zip_url, "-o", "package.zip")
	err = dockercredentialhelperecr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockercredentialhelperecr_bin_url := "https://github.com/awslabs/amazon-ecr-credential-helper/archive/refs/tags/v0.8.0.bin"
	dockercredentialhelperecr_cmd_bin := exec.Command("curl", "-L", dockercredentialhelperecr_bin_url, "-o", "binary.bin")
	err = dockercredentialhelperecr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockercredentialhelperecr_src_url := "https://github.com/awslabs/amazon-ecr-credential-helper/archive/refs/tags/v0.8.0.src.tar.gz"
	dockercredentialhelperecr_cmd_src := exec.Command("curl", "-L", dockercredentialhelperecr_src_url, "-o", "source.tar.gz")
	err = dockercredentialhelperecr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockercredentialhelperecr_cmd_direct := exec.Command("./binary")
	err = dockercredentialhelperecr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
