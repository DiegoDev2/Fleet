package main

// DockerClean - Clean Docker containers, images, networks, and volumes
// Homepage: https://github.com/ZZROTDesign/docker-clean

import (
	"fmt"
	
	"os/exec"
)

func installDockerClean() {
	// Método 1: Descargar y extraer .tar.gz
	dockerclean_tar_url := "https://github.com/ZZROTDesign/docker-clean/archive/refs/tags/v2.0.4.tar.gz"
	dockerclean_cmd_tar := exec.Command("curl", "-L", dockerclean_tar_url, "-o", "package.tar.gz")
	err := dockerclean_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockerclean_zip_url := "https://github.com/ZZROTDesign/docker-clean/archive/refs/tags/v2.0.4.zip"
	dockerclean_cmd_zip := exec.Command("curl", "-L", dockerclean_zip_url, "-o", "package.zip")
	err = dockerclean_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockerclean_bin_url := "https://github.com/ZZROTDesign/docker-clean/archive/refs/tags/v2.0.4.bin"
	dockerclean_cmd_bin := exec.Command("curl", "-L", dockerclean_bin_url, "-o", "binary.bin")
	err = dockerclean_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockerclean_src_url := "https://github.com/ZZROTDesign/docker-clean/archive/refs/tags/v2.0.4.src.tar.gz"
	dockerclean_cmd_src := exec.Command("curl", "-L", dockerclean_src_url, "-o", "source.tar.gz")
	err = dockerclean_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockerclean_cmd_direct := exec.Command("./binary")
	err = dockerclean_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
