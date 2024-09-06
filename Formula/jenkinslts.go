package main

// JenkinsLts - Extendable open source continuous integration server
// Homepage: https://www.jenkins.io/

import (
	"fmt"
	
	"os/exec"
)

func installJenkinsLts() {
	// Método 1: Descargar y extraer .tar.gz
	jenkinslts_tar_url := "https://get.jenkins.io/war-stable/2.462.2/jenkins.war"
	jenkinslts_cmd_tar := exec.Command("curl", "-L", jenkinslts_tar_url, "-o", "package.tar.gz")
	err := jenkinslts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jenkinslts_zip_url := "https://get.jenkins.io/war-stable/2.462.2/jenkins.war"
	jenkinslts_cmd_zip := exec.Command("curl", "-L", jenkinslts_zip_url, "-o", "package.zip")
	err = jenkinslts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jenkinslts_bin_url := "https://get.jenkins.io/war-stable/2.462.2/jenkins.war"
	jenkinslts_cmd_bin := exec.Command("curl", "-L", jenkinslts_bin_url, "-o", "binary.bin")
	err = jenkinslts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jenkinslts_src_url := "https://get.jenkins.io/war-stable/2.462.2/jenkins.war"
	jenkinslts_cmd_src := exec.Command("curl", "-L", jenkinslts_src_url, "-o", "source.tar.gz")
	err = jenkinslts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jenkinslts_cmd_direct := exec.Command("./binary")
	err = jenkinslts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
