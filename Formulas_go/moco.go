package main

// Moco - Stub server with Maven, Gradle, Scala, and shell integration
// Homepage: https://github.com/dreamhead/moco

import (
	"fmt"
	
	"os/exec"
)

func installMoco() {
	// Método 1: Descargar y extraer .tar.gz
	moco_tar_url := "https://search.maven.org/remotecontent?filepath=com/github/dreamhead/moco-runner/1.5.0/moco-runner-1.5.0-standalone.jar"
	moco_cmd_tar := exec.Command("curl", "-L", moco_tar_url, "-o", "package.tar.gz")
	err := moco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moco_zip_url := "https://search.maven.org/remotecontent?filepath=com/github/dreamhead/moco-runner/1.5.0/moco-runner-1.5.0-standalone.jar"
	moco_cmd_zip := exec.Command("curl", "-L", moco_zip_url, "-o", "package.zip")
	err = moco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moco_bin_url := "https://search.maven.org/remotecontent?filepath=com/github/dreamhead/moco-runner/1.5.0/moco-runner-1.5.0-standalone.jar"
	moco_cmd_bin := exec.Command("curl", "-L", moco_bin_url, "-o", "binary.bin")
	err = moco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moco_src_url := "https://search.maven.org/remotecontent?filepath=com/github/dreamhead/moco-runner/1.5.0/moco-runner-1.5.0-standalone.jar"
	moco_cmd_src := exec.Command("curl", "-L", moco_src_url, "-o", "source.tar.gz")
	err = moco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moco_cmd_direct := exec.Command("./binary")
	err = moco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
