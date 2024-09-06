package main

// DockerMachine - Create Docker hosts locally and on cloud providers
// Homepage: https://docs.docker.com/machine

import (
	"fmt"
	
	"os/exec"
)

func installDockerMachine() {
	// Método 1: Descargar y extraer .tar.gz
	dockermachine_tar_url := "https://gitlab.com/gitlab-org/ci-cd/docker-machine/-/archive/v0.16.2-gitlab.28/docker-machine-v0.16.2-gitlab.28.tar.bz2"
	dockermachine_cmd_tar := exec.Command("curl", "-L", dockermachine_tar_url, "-o", "package.tar.gz")
	err := dockermachine_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockermachine_zip_url := "https://gitlab.com/gitlab-org/ci-cd/docker-machine/-/archive/v0.16.2-gitlab.28/docker-machine-v0.16.2-gitlab.28.tar.bz2"
	dockermachine_cmd_zip := exec.Command("curl", "-L", dockermachine_zip_url, "-o", "package.zip")
	err = dockermachine_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockermachine_bin_url := "https://gitlab.com/gitlab-org/ci-cd/docker-machine/-/archive/v0.16.2-gitlab.28/docker-machine-v0.16.2-gitlab.28.tar.bz2"
	dockermachine_cmd_bin := exec.Command("curl", "-L", dockermachine_bin_url, "-o", "binary.bin")
	err = dockermachine_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockermachine_src_url := "https://gitlab.com/gitlab-org/ci-cd/docker-machine/-/archive/v0.16.2-gitlab.28/docker-machine-v0.16.2-gitlab.28.tar.bz2"
	dockermachine_cmd_src := exec.Command("curl", "-L", dockermachine_src_url, "-o", "source.tar.gz")
	err = dockermachine_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockermachine_cmd_direct := exec.Command("./binary")
	err = dockermachine_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
