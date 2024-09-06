package main

// Shellspec - BDD unit testing framework for dash, bash, ksh, zsh and all POSIX shells
// Homepage: https://shellspec.info/

import (
	"fmt"
	
	"os/exec"
)

func installShellspec() {
	// Método 1: Descargar y extraer .tar.gz
	shellspec_tar_url := "https://github.com/shellspec/shellspec/archive/refs/tags/0.28.1.tar.gz"
	shellspec_cmd_tar := exec.Command("curl", "-L", shellspec_tar_url, "-o", "package.tar.gz")
	err := shellspec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shellspec_zip_url := "https://github.com/shellspec/shellspec/archive/refs/tags/0.28.1.zip"
	shellspec_cmd_zip := exec.Command("curl", "-L", shellspec_zip_url, "-o", "package.zip")
	err = shellspec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shellspec_bin_url := "https://github.com/shellspec/shellspec/archive/refs/tags/0.28.1.bin"
	shellspec_cmd_bin := exec.Command("curl", "-L", shellspec_bin_url, "-o", "binary.bin")
	err = shellspec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shellspec_src_url := "https://github.com/shellspec/shellspec/archive/refs/tags/0.28.1.src.tar.gz"
	shellspec_cmd_src := exec.Command("curl", "-L", shellspec_src_url, "-o", "source.tar.gz")
	err = shellspec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shellspec_cmd_direct := exec.Command("./binary")
	err = shellspec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
