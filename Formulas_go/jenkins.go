package main

// Jenkins - Extendable open source continuous integration server
// Homepage: https://www.jenkins.io/

import (
	"fmt"
	
	"os/exec"
)

func installJenkins() {
	// Método 1: Descargar y extraer .tar.gz
	jenkins_tar_url := "https://get.jenkins.io/war/2.475/jenkins.war"
	jenkins_cmd_tar := exec.Command("curl", "-L", jenkins_tar_url, "-o", "package.tar.gz")
	err := jenkins_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jenkins_zip_url := "https://get.jenkins.io/war/2.475/jenkins.war"
	jenkins_cmd_zip := exec.Command("curl", "-L", jenkins_zip_url, "-o", "package.zip")
	err = jenkins_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jenkins_bin_url := "https://get.jenkins.io/war/2.475/jenkins.war"
	jenkins_cmd_bin := exec.Command("curl", "-L", jenkins_bin_url, "-o", "binary.bin")
	err = jenkins_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jenkins_src_url := "https://get.jenkins.io/war/2.475/jenkins.war"
	jenkins_cmd_src := exec.Command("curl", "-L", jenkins_src_url, "-o", "source.tar.gz")
	err = jenkins_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jenkins_cmd_direct := exec.Command("./binary")
	err = jenkins_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: openjdk@21")
exec.Command("latte", "install", "openjdk@21")
}
