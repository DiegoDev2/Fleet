package main

// MCli - Swiss Army Knife for macOS
// Homepage: https://github.com/rgcr/m-cli

import (
	"fmt"
	
	"os/exec"
)

func installMCli() {
	// Método 1: Descargar y extraer .tar.gz
	mcli_tar_url := "https://github.com/rgcr/m-cli/archive/refs/tags/v0.3.0.tar.gz"
	mcli_cmd_tar := exec.Command("curl", "-L", mcli_tar_url, "-o", "package.tar.gz")
	err := mcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mcli_zip_url := "https://github.com/rgcr/m-cli/archive/refs/tags/v0.3.0.zip"
	mcli_cmd_zip := exec.Command("curl", "-L", mcli_zip_url, "-o", "package.zip")
	err = mcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mcli_bin_url := "https://github.com/rgcr/m-cli/archive/refs/tags/v0.3.0.bin"
	mcli_cmd_bin := exec.Command("curl", "-L", mcli_bin_url, "-o", "binary.bin")
	err = mcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mcli_src_url := "https://github.com/rgcr/m-cli/archive/refs/tags/v0.3.0.src.tar.gz"
	mcli_cmd_src := exec.Command("curl", "-L", mcli_src_url, "-o", "source.tar.gz")
	err = mcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mcli_cmd_direct := exec.Command("./binary")
	err = mcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
