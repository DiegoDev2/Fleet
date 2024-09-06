package main

// Slackcat - Command-line utility for posting snippets to Slack
// Homepage: https://github.com/bcicen/slackcat

import (
	"fmt"
	
	"os/exec"
)

func installSlackcat() {
	// Método 1: Descargar y extraer .tar.gz
	slackcat_tar_url := "https://github.com/bcicen/slackcat/archive/refs/tags/1.7.3.tar.gz"
	slackcat_cmd_tar := exec.Command("curl", "-L", slackcat_tar_url, "-o", "package.tar.gz")
	err := slackcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slackcat_zip_url := "https://github.com/bcicen/slackcat/archive/refs/tags/1.7.3.zip"
	slackcat_cmd_zip := exec.Command("curl", "-L", slackcat_zip_url, "-o", "package.zip")
	err = slackcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slackcat_bin_url := "https://github.com/bcicen/slackcat/archive/refs/tags/1.7.3.bin"
	slackcat_cmd_bin := exec.Command("curl", "-L", slackcat_bin_url, "-o", "binary.bin")
	err = slackcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slackcat_src_url := "https://github.com/bcicen/slackcat/archive/refs/tags/1.7.3.src.tar.gz"
	slackcat_cmd_src := exec.Command("curl", "-L", slackcat_src_url, "-o", "source.tar.gz")
	err = slackcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slackcat_cmd_direct := exec.Command("./binary")
	err = slackcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
