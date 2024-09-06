package main

// DockerSlim - Minify and secure Docker images
// Homepage: https://slimtoolkit.org/

import (
	"fmt"
	
	"os/exec"
)

func installDockerSlim() {
	// Método 1: Descargar y extraer .tar.gz
	dockerslim_tar_url := "https://github.com/slimtoolkit/slim/archive/refs/tags/1.40.11.tar.gz"
	dockerslim_cmd_tar := exec.Command("curl", "-L", dockerslim_tar_url, "-o", "package.tar.gz")
	err := dockerslim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockerslim_zip_url := "https://github.com/slimtoolkit/slim/archive/refs/tags/1.40.11.zip"
	dockerslim_cmd_zip := exec.Command("curl", "-L", dockerslim_zip_url, "-o", "package.zip")
	err = dockerslim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockerslim_bin_url := "https://github.com/slimtoolkit/slim/archive/refs/tags/1.40.11.bin"
	dockerslim_cmd_bin := exec.Command("curl", "-L", dockerslim_bin_url, "-o", "binary.bin")
	err = dockerslim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockerslim_src_url := "https://github.com/slimtoolkit/slim/archive/refs/tags/1.40.11.src.tar.gz"
	dockerslim_cmd_src := exec.Command("curl", "-L", dockerslim_src_url, "-o", "source.tar.gz")
	err = dockerslim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockerslim_cmd_direct := exec.Command("./binary")
	err = dockerslim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
