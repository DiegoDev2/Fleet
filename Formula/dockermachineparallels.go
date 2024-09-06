package main

// DockerMachineParallels - Parallels Driver for Docker Machine
// Homepage: https://github.com/Parallels/docker-machine-parallels

import (
	"fmt"
	
	"os/exec"
)

func installDockerMachineParallels() {
	// Método 1: Descargar y extraer .tar.gz
	dockermachineparallels_tar_url := "https://github.com/Parallels/docker-machine-parallels.git"
	dockermachineparallels_cmd_tar := exec.Command("curl", "-L", dockermachineparallels_tar_url, "-o", "package.tar.gz")
	err := dockermachineparallels_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockermachineparallels_zip_url := "https://github.com/Parallels/docker-machine-parallels.git"
	dockermachineparallels_cmd_zip := exec.Command("curl", "-L", dockermachineparallels_zip_url, "-o", "package.zip")
	err = dockermachineparallels_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockermachineparallels_bin_url := "https://github.com/Parallels/docker-machine-parallels.git"
	dockermachineparallels_cmd_bin := exec.Command("curl", "-L", dockermachineparallels_bin_url, "-o", "binary.bin")
	err = dockermachineparallels_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockermachineparallels_src_url := "https://github.com/Parallels/docker-machine-parallels.git"
	dockermachineparallels_cmd_src := exec.Command("curl", "-L", dockermachineparallels_src_url, "-o", "source.tar.gz")
	err = dockermachineparallels_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockermachineparallels_cmd_direct := exec.Command("./binary")
	err = dockermachineparallels_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: docker-machine")
	exec.Command("latte", "install", "docker-machine").Run()
}
