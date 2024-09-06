package main

// OrcTools - ORC java command-line tools and utilities
// Homepage: https://orc.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installOrcTools() {
	// Método 1: Descargar y extraer .tar.gz
	orctools_tar_url := "https://search.maven.org/remotecontent?filepath=org/apache/orc/orc-tools/2.0.2/orc-tools-2.0.2-uber.jar"
	orctools_cmd_tar := exec.Command("curl", "-L", orctools_tar_url, "-o", "package.tar.gz")
	err := orctools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orctools_zip_url := "https://search.maven.org/remotecontent?filepath=org/apache/orc/orc-tools/2.0.2/orc-tools-2.0.2-uber.jar"
	orctools_cmd_zip := exec.Command("curl", "-L", orctools_zip_url, "-o", "package.zip")
	err = orctools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orctools_bin_url := "https://search.maven.org/remotecontent?filepath=org/apache/orc/orc-tools/2.0.2/orc-tools-2.0.2-uber.jar"
	orctools_cmd_bin := exec.Command("curl", "-L", orctools_bin_url, "-o", "binary.bin")
	err = orctools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orctools_src_url := "https://search.maven.org/remotecontent?filepath=org/apache/orc/orc-tools/2.0.2/orc-tools-2.0.2-uber.jar"
	orctools_cmd_src := exec.Command("curl", "-L", orctools_src_url, "-o", "source.tar.gz")
	err = orctools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orctools_cmd_direct := exec.Command("./binary")
	err = orctools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
