package main

// DockerSquash - Docker image squashing tool
// Homepage: https://github.com/goldmann/docker-squash

import (
	"fmt"
	
	"os/exec"
)

func installDockerSquash() {
	// Método 1: Descargar y extraer .tar.gz
	dockersquash_tar_url := "https://files.pythonhosted.org/packages/c9/10/4bf67bb4e15414f2a5e2a7d20fe773c96f471223b1a4d6dd60f1fe2e6365/docker_squash-1.2.1.tar.gz"
	dockersquash_cmd_tar := exec.Command("curl", "-L", dockersquash_tar_url, "-o", "package.tar.gz")
	err := dockersquash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockersquash_zip_url := "https://files.pythonhosted.org/packages/c9/10/4bf67bb4e15414f2a5e2a7d20fe773c96f471223b1a4d6dd60f1fe2e6365/docker_squash-1.2.1.zip"
	dockersquash_cmd_zip := exec.Command("curl", "-L", dockersquash_zip_url, "-o", "package.zip")
	err = dockersquash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockersquash_bin_url := "https://files.pythonhosted.org/packages/c9/10/4bf67bb4e15414f2a5e2a7d20fe773c96f471223b1a4d6dd60f1fe2e6365/docker_squash-1.2.1.bin"
	dockersquash_cmd_bin := exec.Command("curl", "-L", dockersquash_bin_url, "-o", "binary.bin")
	err = dockersquash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockersquash_src_url := "https://files.pythonhosted.org/packages/c9/10/4bf67bb4e15414f2a5e2a7d20fe773c96f471223b1a4d6dd60f1fe2e6365/docker_squash-1.2.1.src.tar.gz"
	dockersquash_cmd_src := exec.Command("curl", "-L", dockersquash_src_url, "-o", "source.tar.gz")
	err = dockersquash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockersquash_cmd_direct := exec.Command("./binary")
	err = dockersquash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
