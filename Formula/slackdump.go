package main

// Slackdump - Export Slack data without admin privileges
// Homepage: https://github.com/rusq/slackdump

import (
	"fmt"
	
	"os/exec"
)

func installSlackdump() {
	// Método 1: Descargar y extraer .tar.gz
	slackdump_tar_url := "https://github.com/rusq/slackdump/archive/refs/tags/v2.5.11.tar.gz"
	slackdump_cmd_tar := exec.Command("curl", "-L", slackdump_tar_url, "-o", "package.tar.gz")
	err := slackdump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slackdump_zip_url := "https://github.com/rusq/slackdump/archive/refs/tags/v2.5.11.zip"
	slackdump_cmd_zip := exec.Command("curl", "-L", slackdump_zip_url, "-o", "package.zip")
	err = slackdump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slackdump_bin_url := "https://github.com/rusq/slackdump/archive/refs/tags/v2.5.11.bin"
	slackdump_cmd_bin := exec.Command("curl", "-L", slackdump_bin_url, "-o", "binary.bin")
	err = slackdump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slackdump_src_url := "https://github.com/rusq/slackdump/archive/refs/tags/v2.5.11.src.tar.gz"
	slackdump_cmd_src := exec.Command("curl", "-L", slackdump_src_url, "-o", "source.tar.gz")
	err = slackdump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slackdump_cmd_direct := exec.Command("./binary")
	err = slackdump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
