package main

// Liquidprompt - Adaptive prompt for bash and zsh shells
// Homepage: https://liquidprompt.readthedocs.io/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installLiquidprompt() {
	// Método 1: Descargar y extraer .tar.gz
	liquidprompt_tar_url := "https://github.com/liquidprompt/liquidprompt/archive/refs/tags/v2.2.1.tar.gz"
	liquidprompt_cmd_tar := exec.Command("curl", "-L", liquidprompt_tar_url, "-o", "package.tar.gz")
	err := liquidprompt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liquidprompt_zip_url := "https://github.com/liquidprompt/liquidprompt/archive/refs/tags/v2.2.1.zip"
	liquidprompt_cmd_zip := exec.Command("curl", "-L", liquidprompt_zip_url, "-o", "package.zip")
	err = liquidprompt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liquidprompt_bin_url := "https://github.com/liquidprompt/liquidprompt/archive/refs/tags/v2.2.1.bin"
	liquidprompt_cmd_bin := exec.Command("curl", "-L", liquidprompt_bin_url, "-o", "binary.bin")
	err = liquidprompt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liquidprompt_src_url := "https://github.com/liquidprompt/liquidprompt/archive/refs/tags/v2.2.1.src.tar.gz"
	liquidprompt_cmd_src := exec.Command("curl", "-L", liquidprompt_src_url, "-o", "source.tar.gz")
	err = liquidprompt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liquidprompt_cmd_direct := exec.Command("./binary")
	err = liquidprompt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
