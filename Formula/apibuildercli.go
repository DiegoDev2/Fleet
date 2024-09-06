package main

// ApibuilderCli - Command-line interface to generate clients for api builder
// Homepage: https://www.apibuilder.io

import (
	"fmt"
	
	"os/exec"
)

func installApibuilderCli() {
	// Método 1: Descargar y extraer .tar.gz
	apibuildercli_tar_url := "https://github.com/apicollective/apibuilder-cli/archive/refs/tags/0.1.52.tar.gz"
	apibuildercli_cmd_tar := exec.Command("curl", "-L", apibuildercli_tar_url, "-o", "package.tar.gz")
	err := apibuildercli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apibuildercli_zip_url := "https://github.com/apicollective/apibuilder-cli/archive/refs/tags/0.1.52.zip"
	apibuildercli_cmd_zip := exec.Command("curl", "-L", apibuildercli_zip_url, "-o", "package.zip")
	err = apibuildercli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apibuildercli_bin_url := "https://github.com/apicollective/apibuilder-cli/archive/refs/tags/0.1.52.bin"
	apibuildercli_cmd_bin := exec.Command("curl", "-L", apibuildercli_bin_url, "-o", "binary.bin")
	err = apibuildercli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apibuildercli_src_url := "https://github.com/apicollective/apibuilder-cli/archive/refs/tags/0.1.52.src.tar.gz"
	apibuildercli_cmd_src := exec.Command("curl", "-L", apibuildercli_src_url, "-o", "source.tar.gz")
	err = apibuildercli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apibuildercli_cmd_direct := exec.Command("./binary")
	err = apibuildercli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
