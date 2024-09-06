package main

// WrenCli - Simple REPL and CLI tool for running Wren scripts
// Homepage: https://github.com/wren-lang/wren-cli

import (
	"fmt"
	
	"os/exec"
)

func installWrenCli() {
	// Método 1: Descargar y extraer .tar.gz
	wrencli_tar_url := "https://github.com/wren-lang/wren-cli/archive/refs/tags/0.4.0.tar.gz"
	wrencli_cmd_tar := exec.Command("curl", "-L", wrencli_tar_url, "-o", "package.tar.gz")
	err := wrencli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wrencli_zip_url := "https://github.com/wren-lang/wren-cli/archive/refs/tags/0.4.0.zip"
	wrencli_cmd_zip := exec.Command("curl", "-L", wrencli_zip_url, "-o", "package.zip")
	err = wrencli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wrencli_bin_url := "https://github.com/wren-lang/wren-cli/archive/refs/tags/0.4.0.bin"
	wrencli_cmd_bin := exec.Command("curl", "-L", wrencli_bin_url, "-o", "binary.bin")
	err = wrencli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wrencli_src_url := "https://github.com/wren-lang/wren-cli/archive/refs/tags/0.4.0.src.tar.gz"
	wrencli_cmd_src := exec.Command("curl", "-L", wrencli_src_url, "-o", "source.tar.gz")
	err = wrencli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wrencli_cmd_direct := exec.Command("./binary")
	err = wrencli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
