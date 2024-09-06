package main

// Temporal - Command-line interface for running and interacting with Temporal Server and UI
// Homepage: https://temporal.io/

import (
	"fmt"
	
	"os/exec"
)

func installTemporal() {
	// Método 1: Descargar y extraer .tar.gz
	temporal_tar_url := "https://github.com/temporalio/cli/archive/refs/tags/v1.0.0.tar.gz"
	temporal_cmd_tar := exec.Command("curl", "-L", temporal_tar_url, "-o", "package.tar.gz")
	err := temporal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	temporal_zip_url := "https://github.com/temporalio/cli/archive/refs/tags/v1.0.0.zip"
	temporal_cmd_zip := exec.Command("curl", "-L", temporal_zip_url, "-o", "package.zip")
	err = temporal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	temporal_bin_url := "https://github.com/temporalio/cli/archive/refs/tags/v1.0.0.bin"
	temporal_cmd_bin := exec.Command("curl", "-L", temporal_bin_url, "-o", "binary.bin")
	err = temporal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	temporal_src_url := "https://github.com/temporalio/cli/archive/refs/tags/v1.0.0.src.tar.gz"
	temporal_cmd_src := exec.Command("curl", "-L", temporal_src_url, "-o", "source.tar.gz")
	err = temporal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	temporal_cmd_direct := exec.Command("./binary")
	err = temporal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
