package main

// DockerMachineDriverVultr - Docker Machine driver plugin for Vultr Cloud
// Homepage: https://github.com/vultr/docker-machine-driver-vultr

import (
	"fmt"
	
	"os/exec"
)

func installDockerMachineDriverVultr() {
	// Método 1: Descargar y extraer .tar.gz
	dockermachinedrivervultr_tar_url := "https://github.com/vultr/docker-machine-driver-vultr/archive/refs/tags/v2.1.0.tar.gz"
	dockermachinedrivervultr_cmd_tar := exec.Command("curl", "-L", dockermachinedrivervultr_tar_url, "-o", "package.tar.gz")
	err := dockermachinedrivervultr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockermachinedrivervultr_zip_url := "https://github.com/vultr/docker-machine-driver-vultr/archive/refs/tags/v2.1.0.zip"
	dockermachinedrivervultr_cmd_zip := exec.Command("curl", "-L", dockermachinedrivervultr_zip_url, "-o", "package.zip")
	err = dockermachinedrivervultr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockermachinedrivervultr_bin_url := "https://github.com/vultr/docker-machine-driver-vultr/archive/refs/tags/v2.1.0.bin"
	dockermachinedrivervultr_cmd_bin := exec.Command("curl", "-L", dockermachinedrivervultr_bin_url, "-o", "binary.bin")
	err = dockermachinedrivervultr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockermachinedrivervultr_src_url := "https://github.com/vultr/docker-machine-driver-vultr/archive/refs/tags/v2.1.0.src.tar.gz"
	dockermachinedrivervultr_cmd_src := exec.Command("curl", "-L", dockermachinedrivervultr_src_url, "-o", "source.tar.gz")
	err = dockermachinedrivervultr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockermachinedrivervultr_cmd_direct := exec.Command("./binary")
	err = dockermachinedrivervultr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: docker-machine")
exec.Command("latte", "install", "docker-machine")
}
