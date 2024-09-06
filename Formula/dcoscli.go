package main

// DcosCli - Command-line interface for managing DC/OS clusters
// Homepage: https://docs.d2iq.com/mesosphere/dcos/latest/cli

import (
	"fmt"
	
	"os/exec"
)

func installDcosCli() {
	// Método 1: Descargar y extraer .tar.gz
	dcoscli_tar_url := "https://github.com/dcos/dcos-cli/archive/refs/tags/1.2.0.tar.gz"
	dcoscli_cmd_tar := exec.Command("curl", "-L", dcoscli_tar_url, "-o", "package.tar.gz")
	err := dcoscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dcoscli_zip_url := "https://github.com/dcos/dcos-cli/archive/refs/tags/1.2.0.zip"
	dcoscli_cmd_zip := exec.Command("curl", "-L", dcoscli_zip_url, "-o", "package.zip")
	err = dcoscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dcoscli_bin_url := "https://github.com/dcos/dcos-cli/archive/refs/tags/1.2.0.bin"
	dcoscli_cmd_bin := exec.Command("curl", "-L", dcoscli_bin_url, "-o", "binary.bin")
	err = dcoscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dcoscli_src_url := "https://github.com/dcos/dcos-cli/archive/refs/tags/1.2.0.src.tar.gz"
	dcoscli_cmd_src := exec.Command("curl", "-L", dcoscli_src_url, "-o", "source.tar.gz")
	err = dcoscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dcoscli_cmd_direct := exec.Command("./binary")
	err = dcoscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
