package main

// Gator - CLI Utility for Open Policy Agent Gatekeeper
// Homepage: https://open-policy-agent.github.io/gatekeeper/website/docs/gator

import (
	"fmt"
	
	"os/exec"
)

func installGator() {
	// Método 1: Descargar y extraer .tar.gz
	gator_tar_url := "https://github.com/open-policy-agent/gatekeeper/archive/refs/tags/v3.17.0.tar.gz"
	gator_cmd_tar := exec.Command("curl", "-L", gator_tar_url, "-o", "package.tar.gz")
	err := gator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gator_zip_url := "https://github.com/open-policy-agent/gatekeeper/archive/refs/tags/v3.17.0.zip"
	gator_cmd_zip := exec.Command("curl", "-L", gator_zip_url, "-o", "package.zip")
	err = gator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gator_bin_url := "https://github.com/open-policy-agent/gatekeeper/archive/refs/tags/v3.17.0.bin"
	gator_cmd_bin := exec.Command("curl", "-L", gator_bin_url, "-o", "binary.bin")
	err = gator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gator_src_url := "https://github.com/open-policy-agent/gatekeeper/archive/refs/tags/v3.17.0.src.tar.gz"
	gator_cmd_src := exec.Command("curl", "-L", gator_src_url, "-o", "source.tar.gz")
	err = gator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gator_cmd_direct := exec.Command("./binary")
	err = gator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
