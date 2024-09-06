package main

// DockerCompletion - Bash, Zsh and Fish completion for Docker
// Homepage: https://www.docker.com/

import (
	"fmt"
	
	"os/exec"
)

func installDockerCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	dockercompletion_tar_url := "https://github.com/docker/cli/archive/refs/tags/v27.2.0.tar.gz"
	dockercompletion_cmd_tar := exec.Command("curl", "-L", dockercompletion_tar_url, "-o", "package.tar.gz")
	err := dockercompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockercompletion_zip_url := "https://github.com/docker/cli/archive/refs/tags/v27.2.0.zip"
	dockercompletion_cmd_zip := exec.Command("curl", "-L", dockercompletion_zip_url, "-o", "package.zip")
	err = dockercompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockercompletion_bin_url := "https://github.com/docker/cli/archive/refs/tags/v27.2.0.bin"
	dockercompletion_cmd_bin := exec.Command("curl", "-L", dockercompletion_bin_url, "-o", "binary.bin")
	err = dockercompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockercompletion_src_url := "https://github.com/docker/cli/archive/refs/tags/v27.2.0.src.tar.gz"
	dockercompletion_cmd_src := exec.Command("curl", "-L", dockercompletion_src_url, "-o", "source.tar.gz")
	err = dockercompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockercompletion_cmd_direct := exec.Command("./binary")
	err = dockercompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
