package main

// Repl - Wrap non-interactive programs with a REPL
// Homepage: https://github.com/defunkt/repl

import (
	"fmt"
	
	"os/exec"
)

func installRepl() {
	// Método 1: Descargar y extraer .tar.gz
	repl_tar_url := "https://github.com/defunkt/repl/archive/refs/tags/v1.0.0.tar.gz"
	repl_cmd_tar := exec.Command("curl", "-L", repl_tar_url, "-o", "package.tar.gz")
	err := repl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	repl_zip_url := "https://github.com/defunkt/repl/archive/refs/tags/v1.0.0.zip"
	repl_cmd_zip := exec.Command("curl", "-L", repl_zip_url, "-o", "package.zip")
	err = repl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	repl_bin_url := "https://github.com/defunkt/repl/archive/refs/tags/v1.0.0.bin"
	repl_cmd_bin := exec.Command("curl", "-L", repl_bin_url, "-o", "binary.bin")
	err = repl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	repl_src_url := "https://github.com/defunkt/repl/archive/refs/tags/v1.0.0.src.tar.gz"
	repl_cmd_src := exec.Command("curl", "-L", repl_src_url, "-o", "source.tar.gz")
	err = repl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	repl_cmd_direct := exec.Command("./binary")
	err = repl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
