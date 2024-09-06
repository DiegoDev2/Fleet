package main

// RustcCompletion - Bash completion for rustc
// Homepage: https://github.com/roshan/rust-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installRustcCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	rustccompletion_tar_url := "https://github.com/roshan/rust-bash-completion/archive/refs/tags/0.12.1.tar.gz"
	rustccompletion_cmd_tar := exec.Command("curl", "-L", rustccompletion_tar_url, "-o", "package.tar.gz")
	err := rustccompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustccompletion_zip_url := "https://github.com/roshan/rust-bash-completion/archive/refs/tags/0.12.1.zip"
	rustccompletion_cmd_zip := exec.Command("curl", "-L", rustccompletion_zip_url, "-o", "package.zip")
	err = rustccompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustccompletion_bin_url := "https://github.com/roshan/rust-bash-completion/archive/refs/tags/0.12.1.bin"
	rustccompletion_cmd_bin := exec.Command("curl", "-L", rustccompletion_bin_url, "-o", "binary.bin")
	err = rustccompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustccompletion_src_url := "https://github.com/roshan/rust-bash-completion/archive/refs/tags/0.12.1.src.tar.gz"
	rustccompletion_cmd_src := exec.Command("curl", "-L", rustccompletion_src_url, "-o", "source.tar.gz")
	err = rustccompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustccompletion_cmd_direct := exec.Command("./binary")
	err = rustccompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
