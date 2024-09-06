package main

// Kumo - Word Clouds in Java
// Homepage: https://github.com/kennycason/kumo

import (
	"fmt"
	
	"os/exec"
)

func installKumo() {
	// Método 1: Descargar y extraer .tar.gz
	kumo_tar_url := "https://search.maven.org/remotecontent?filepath=com/kennycason/kumo-cli/1.28/kumo-cli-1.28.jar"
	kumo_cmd_tar := exec.Command("curl", "-L", kumo_tar_url, "-o", "package.tar.gz")
	err := kumo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kumo_zip_url := "https://search.maven.org/remotecontent?filepath=com/kennycason/kumo-cli/1.28/kumo-cli-1.28.jar"
	kumo_cmd_zip := exec.Command("curl", "-L", kumo_zip_url, "-o", "package.zip")
	err = kumo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kumo_bin_url := "https://search.maven.org/remotecontent?filepath=com/kennycason/kumo-cli/1.28/kumo-cli-1.28.jar"
	kumo_cmd_bin := exec.Command("curl", "-L", kumo_bin_url, "-o", "binary.bin")
	err = kumo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kumo_src_url := "https://search.maven.org/remotecontent?filepath=com/kennycason/kumo-cli/1.28/kumo-cli-1.28.jar"
	kumo_cmd_src := exec.Command("curl", "-L", kumo_src_url, "-o", "source.tar.gz")
	err = kumo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kumo_cmd_direct := exec.Command("./binary")
	err = kumo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
