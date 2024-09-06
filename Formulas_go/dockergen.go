package main

// DockerGen - Generate files from docker container metadata
// Homepage: https://github.com/nginx-proxy/docker-gen

import (
	"fmt"
	
	"os/exec"
)

func installDockerGen() {
	// Método 1: Descargar y extraer .tar.gz
	dockergen_tar_url := "https://github.com/nginx-proxy/docker-gen/archive/refs/tags/0.14.2.tar.gz"
	dockergen_cmd_tar := exec.Command("curl", "-L", dockergen_tar_url, "-o", "package.tar.gz")
	err := dockergen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockergen_zip_url := "https://github.com/nginx-proxy/docker-gen/archive/refs/tags/0.14.2.zip"
	dockergen_cmd_zip := exec.Command("curl", "-L", dockergen_zip_url, "-o", "package.zip")
	err = dockergen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockergen_bin_url := "https://github.com/nginx-proxy/docker-gen/archive/refs/tags/0.14.2.bin"
	dockergen_cmd_bin := exec.Command("curl", "-L", dockergen_bin_url, "-o", "binary.bin")
	err = dockergen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockergen_src_url := "https://github.com/nginx-proxy/docker-gen/archive/refs/tags/0.14.2.src.tar.gz"
	dockergen_cmd_src := exec.Command("curl", "-L", dockergen_src_url, "-o", "source.tar.gz")
	err = dockergen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockergen_cmd_direct := exec.Command("./binary")
	err = dockergen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
