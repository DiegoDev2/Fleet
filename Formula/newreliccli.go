package main

// NewrelicCli - Command-line interface for New Relic
// Homepage: https://github.com/newrelic/newrelic-cli

import (
	"fmt"
	
	"os/exec"
)

func installNewrelicCli() {
	// Método 1: Descargar y extraer .tar.gz
	newreliccli_tar_url := "https://github.com/newrelic/newrelic-cli/archive/refs/tags/v0.93.4.tar.gz"
	newreliccli_cmd_tar := exec.Command("curl", "-L", newreliccli_tar_url, "-o", "package.tar.gz")
	err := newreliccli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	newreliccli_zip_url := "https://github.com/newrelic/newrelic-cli/archive/refs/tags/v0.93.4.zip"
	newreliccli_cmd_zip := exec.Command("curl", "-L", newreliccli_zip_url, "-o", "package.zip")
	err = newreliccli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	newreliccli_bin_url := "https://github.com/newrelic/newrelic-cli/archive/refs/tags/v0.93.4.bin"
	newreliccli_cmd_bin := exec.Command("curl", "-L", newreliccli_bin_url, "-o", "binary.bin")
	err = newreliccli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	newreliccli_src_url := "https://github.com/newrelic/newrelic-cli/archive/refs/tags/v0.93.4.src.tar.gz"
	newreliccli_cmd_src := exec.Command("curl", "-L", newreliccli_src_url, "-o", "source.tar.gz")
	err = newreliccli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	newreliccli_cmd_direct := exec.Command("./binary")
	err = newreliccli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
