package main

// WiremockStandalone - Simulator for HTTP-based APIs
// Homepage: https://wiremock.org/docs/running-standalone/

import (
	"fmt"
	
	"os/exec"
)

func installWiremockStandalone() {
	// Método 1: Descargar y extraer .tar.gz
	wiremockstandalone_tar_url := "https://search.maven.org/remotecontent?filepath=org/wiremock/wiremock-standalone/3.9.1/wiremock-standalone-3.9.1.jar"
	wiremockstandalone_cmd_tar := exec.Command("curl", "-L", wiremockstandalone_tar_url, "-o", "package.tar.gz")
	err := wiremockstandalone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wiremockstandalone_zip_url := "https://search.maven.org/remotecontent?filepath=org/wiremock/wiremock-standalone/3.9.1/wiremock-standalone-3.9.1.jar"
	wiremockstandalone_cmd_zip := exec.Command("curl", "-L", wiremockstandalone_zip_url, "-o", "package.zip")
	err = wiremockstandalone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wiremockstandalone_bin_url := "https://search.maven.org/remotecontent?filepath=org/wiremock/wiremock-standalone/3.9.1/wiremock-standalone-3.9.1.jar"
	wiremockstandalone_cmd_bin := exec.Command("curl", "-L", wiremockstandalone_bin_url, "-o", "binary.bin")
	err = wiremockstandalone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wiremockstandalone_src_url := "https://search.maven.org/remotecontent?filepath=org/wiremock/wiremock-standalone/3.9.1/wiremock-standalone-3.9.1.jar"
	wiremockstandalone_cmd_src := exec.Command("curl", "-L", wiremockstandalone_src_url, "-o", "source.tar.gz")
	err = wiremockstandalone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wiremockstandalone_cmd_direct := exec.Command("./binary")
	err = wiremockstandalone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
