package main

// TCompletion - Completion for CLI power tool for Twitter
// Homepage: https://sferik.github.io/t/

import (
	"fmt"
	
	"os/exec"
)

func installTCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	tcompletion_tar_url := "https://github.com/sferik/t-ruby/archive/refs/tags/v4.1.1.tar.gz"
	tcompletion_cmd_tar := exec.Command("curl", "-L", tcompletion_tar_url, "-o", "package.tar.gz")
	err := tcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcompletion_zip_url := "https://github.com/sferik/t-ruby/archive/refs/tags/v4.1.1.zip"
	tcompletion_cmd_zip := exec.Command("curl", "-L", tcompletion_zip_url, "-o", "package.zip")
	err = tcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcompletion_bin_url := "https://github.com/sferik/t-ruby/archive/refs/tags/v4.1.1.bin"
	tcompletion_cmd_bin := exec.Command("curl", "-L", tcompletion_bin_url, "-o", "binary.bin")
	err = tcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcompletion_src_url := "https://github.com/sferik/t-ruby/archive/refs/tags/v4.1.1.src.tar.gz"
	tcompletion_cmd_src := exec.Command("curl", "-L", tcompletion_src_url, "-o", "source.tar.gz")
	err = tcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcompletion_cmd_direct := exec.Command("./binary")
	err = tcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
