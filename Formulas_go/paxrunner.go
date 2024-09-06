package main

// PaxRunner - Tool to provision OSGi bundles
// Homepage: https://ops4j1.jira.com/wiki/spaces/paxrunner/overview

import (
	"fmt"
	
	"os/exec"
)

func installPaxRunner() {
	// Método 1: Descargar y extraer .tar.gz
	paxrunner_tar_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/runner/pax-runner-assembly/1.9.0/pax-runner-assembly-1.9.0-jdk15.tar.gz"
	paxrunner_cmd_tar := exec.Command("curl", "-L", paxrunner_tar_url, "-o", "package.tar.gz")
	err := paxrunner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	paxrunner_zip_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/runner/pax-runner-assembly/1.9.0/pax-runner-assembly-1.9.0-jdk15.zip"
	paxrunner_cmd_zip := exec.Command("curl", "-L", paxrunner_zip_url, "-o", "package.zip")
	err = paxrunner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	paxrunner_bin_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/runner/pax-runner-assembly/1.9.0/pax-runner-assembly-1.9.0-jdk15.bin"
	paxrunner_cmd_bin := exec.Command("curl", "-L", paxrunner_bin_url, "-o", "binary.bin")
	err = paxrunner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	paxrunner_src_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/runner/pax-runner-assembly/1.9.0/pax-runner-assembly-1.9.0-jdk15.src.tar.gz"
	paxrunner_cmd_src := exec.Command("curl", "-L", paxrunner_src_url, "-o", "source.tar.gz")
	err = paxrunner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	paxrunner_cmd_direct := exec.Command("./binary")
	err = paxrunner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
