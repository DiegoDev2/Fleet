package main

// DockerCompose - Isolated development environments using Docker
// Homepage: https://docs.docker.com/compose/

import (
	"fmt"
	
	"os/exec"
)

func installDockerCompose() {
	// Método 1: Descargar y extraer .tar.gz
	dockercompose_tar_url := "https://github.com/docker/compose/archive/refs/tags/v2.29.2.tar.gz"
	dockercompose_cmd_tar := exec.Command("curl", "-L", dockercompose_tar_url, "-o", "package.tar.gz")
	err := dockercompose_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockercompose_zip_url := "https://github.com/docker/compose/archive/refs/tags/v2.29.2.zip"
	dockercompose_cmd_zip := exec.Command("curl", "-L", dockercompose_zip_url, "-o", "package.zip")
	err = dockercompose_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockercompose_bin_url := "https://github.com/docker/compose/archive/refs/tags/v2.29.2.bin"
	dockercompose_cmd_bin := exec.Command("curl", "-L", dockercompose_bin_url, "-o", "binary.bin")
	err = dockercompose_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockercompose_src_url := "https://github.com/docker/compose/archive/refs/tags/v2.29.2.src.tar.gz"
	dockercompose_cmd_src := exec.Command("curl", "-L", dockercompose_src_url, "-o", "source.tar.gz")
	err = dockercompose_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockercompose_cmd_direct := exec.Command("./binary")
	err = dockercompose_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
