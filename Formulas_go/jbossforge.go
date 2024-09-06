package main

// JbossForge - Tools to help set up and configure a project
// Homepage: https://forge.jboss.org/

import (
	"fmt"
	
	"os/exec"
)

func installJbossForge() {
	// Método 1: Descargar y extraer .tar.gz
	jbossforge_tar_url := "https://downloads.jboss.org/forge/releases/3.10.0.Final/forge-distribution-3.10.0.Final-offline.zip"
	jbossforge_cmd_tar := exec.Command("curl", "-L", jbossforge_tar_url, "-o", "package.tar.gz")
	err := jbossforge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jbossforge_zip_url := "https://downloads.jboss.org/forge/releases/3.10.0.Final/forge-distribution-3.10.0.Final-offline.zip"
	jbossforge_cmd_zip := exec.Command("curl", "-L", jbossforge_zip_url, "-o", "package.zip")
	err = jbossforge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jbossforge_bin_url := "https://downloads.jboss.org/forge/releases/3.10.0.Final/forge-distribution-3.10.0.Final-offline.zip"
	jbossforge_cmd_bin := exec.Command("curl", "-L", jbossforge_bin_url, "-o", "binary.bin")
	err = jbossforge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jbossforge_src_url := "https://downloads.jboss.org/forge/releases/3.10.0.Final/forge-distribution-3.10.0.Final-offline.zip"
	jbossforge_cmd_src := exec.Command("curl", "-L", jbossforge_src_url, "-o", "source.tar.gz")
	err = jbossforge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jbossforge_cmd_direct := exec.Command("./binary")
	err = jbossforge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
