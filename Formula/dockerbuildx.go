package main

// DockerBuildx - Docker CLI plugin for extended build capabilities with BuildKit
// Homepage: https://docs.docker.com/buildx/working-with-buildx/

import (
	"fmt"
	
	"os/exec"
)

func installDockerBuildx() {
	// Método 1: Descargar y extraer .tar.gz
	dockerbuildx_tar_url := "https://github.com/docker/buildx/archive/refs/tags/v0.16.2.tar.gz"
	dockerbuildx_cmd_tar := exec.Command("curl", "-L", dockerbuildx_tar_url, "-o", "package.tar.gz")
	err := dockerbuildx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockerbuildx_zip_url := "https://github.com/docker/buildx/archive/refs/tags/v0.16.2.zip"
	dockerbuildx_cmd_zip := exec.Command("curl", "-L", dockerbuildx_zip_url, "-o", "package.zip")
	err = dockerbuildx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockerbuildx_bin_url := "https://github.com/docker/buildx/archive/refs/tags/v0.16.2.bin"
	dockerbuildx_cmd_bin := exec.Command("curl", "-L", dockerbuildx_bin_url, "-o", "binary.bin")
	err = dockerbuildx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockerbuildx_src_url := "https://github.com/docker/buildx/archive/refs/tags/v0.16.2.src.tar.gz"
	dockerbuildx_cmd_src := exec.Command("curl", "-L", dockerbuildx_src_url, "-o", "source.tar.gz")
	err = dockerbuildx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockerbuildx_cmd_direct := exec.Command("./binary")
	err = dockerbuildx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
