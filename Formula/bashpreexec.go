package main

// BashPreexec - Preexec and precmd functions for Bash (like Zsh)
// Homepage: https://github.com/rcaloras/bash-preexec

import (
	"fmt"
	
	"os/exec"
)

func installBashPreexec() {
	// Método 1: Descargar y extraer .tar.gz
	bashpreexec_tar_url := "https://github.com/rcaloras/bash-preexec/archive/refs/tags/0.5.0.tar.gz"
	bashpreexec_cmd_tar := exec.Command("curl", "-L", bashpreexec_tar_url, "-o", "package.tar.gz")
	err := bashpreexec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashpreexec_zip_url := "https://github.com/rcaloras/bash-preexec/archive/refs/tags/0.5.0.zip"
	bashpreexec_cmd_zip := exec.Command("curl", "-L", bashpreexec_zip_url, "-o", "package.zip")
	err = bashpreexec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashpreexec_bin_url := "https://github.com/rcaloras/bash-preexec/archive/refs/tags/0.5.0.bin"
	bashpreexec_cmd_bin := exec.Command("curl", "-L", bashpreexec_bin_url, "-o", "binary.bin")
	err = bashpreexec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashpreexec_src_url := "https://github.com/rcaloras/bash-preexec/archive/refs/tags/0.5.0.src.tar.gz"
	bashpreexec_cmd_src := exec.Command("curl", "-L", bashpreexec_src_url, "-o", "source.tar.gz")
	err = bashpreexec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashpreexec_cmd_direct := exec.Command("./binary")
	err = bashpreexec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
