package main

// PaxConstruct - Tools to setup and develop OSGi projects quickly
// Homepage: https://ops4j1.jira.com/wiki/spaces/paxconstruct/overview

import (
	"fmt"
	
	"os/exec"
)

func installPaxConstruct() {
	// Método 1: Descargar y extraer .tar.gz
	paxconstruct_tar_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/construct/scripts/1.6.0/scripts-1.6.0.zip"
	paxconstruct_cmd_tar := exec.Command("curl", "-L", paxconstruct_tar_url, "-o", "package.tar.gz")
	err := paxconstruct_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	paxconstruct_zip_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/construct/scripts/1.6.0/scripts-1.6.0.zip"
	paxconstruct_cmd_zip := exec.Command("curl", "-L", paxconstruct_zip_url, "-o", "package.zip")
	err = paxconstruct_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	paxconstruct_bin_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/construct/scripts/1.6.0/scripts-1.6.0.zip"
	paxconstruct_cmd_bin := exec.Command("curl", "-L", paxconstruct_bin_url, "-o", "binary.bin")
	err = paxconstruct_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	paxconstruct_src_url := "https://search.maven.org/remotecontent?filepath=org/ops4j/pax/construct/scripts/1.6.0/scripts-1.6.0.zip"
	paxconstruct_cmd_src := exec.Command("curl", "-L", paxconstruct_src_url, "-o", "source.tar.gz")
	err = paxconstruct_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	paxconstruct_cmd_direct := exec.Command("./binary")
	err = paxconstruct_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
}
