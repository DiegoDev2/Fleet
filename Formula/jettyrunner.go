package main

// JettyRunner - Use Jetty without an installed distribution
// Homepage: https://jetty.org/

import (
	"fmt"
	
	"os/exec"
)

func installJettyRunner() {
	// Método 1: Descargar y extraer .tar.gz
	jettyrunner_tar_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-runner/11.0.24/jetty-runner-11.0.24.jar"
	jettyrunner_cmd_tar := exec.Command("curl", "-L", jettyrunner_tar_url, "-o", "package.tar.gz")
	err := jettyrunner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jettyrunner_zip_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-runner/11.0.24/jetty-runner-11.0.24.jar"
	jettyrunner_cmd_zip := exec.Command("curl", "-L", jettyrunner_zip_url, "-o", "package.zip")
	err = jettyrunner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jettyrunner_bin_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-runner/11.0.24/jetty-runner-11.0.24.jar"
	jettyrunner_cmd_bin := exec.Command("curl", "-L", jettyrunner_bin_url, "-o", "binary.bin")
	err = jettyrunner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jettyrunner_src_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-runner/11.0.24/jetty-runner-11.0.24.jar"
	jettyrunner_cmd_src := exec.Command("curl", "-L", jettyrunner_src_url, "-o", "source.tar.gz")
	err = jettyrunner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jettyrunner_cmd_direct := exec.Command("./binary")
	err = jettyrunner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
