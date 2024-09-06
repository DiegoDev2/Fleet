package main

// Prettyping - Wrapper to colorize and simplify ping's output
// Homepage: https://denilsonsa.github.io/prettyping/

import (
	"fmt"
	
	"os/exec"
)

func installPrettyping() {
	// Método 1: Descargar y extraer .tar.gz
	prettyping_tar_url := "https://github.com/denilsonsa/prettyping/archive/refs/tags/v1.0.1.tar.gz"
	prettyping_cmd_tar := exec.Command("curl", "-L", prettyping_tar_url, "-o", "package.tar.gz")
	err := prettyping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prettyping_zip_url := "https://github.com/denilsonsa/prettyping/archive/refs/tags/v1.0.1.zip"
	prettyping_cmd_zip := exec.Command("curl", "-L", prettyping_zip_url, "-o", "package.zip")
	err = prettyping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prettyping_bin_url := "https://github.com/denilsonsa/prettyping/archive/refs/tags/v1.0.1.bin"
	prettyping_cmd_bin := exec.Command("curl", "-L", prettyping_bin_url, "-o", "binary.bin")
	err = prettyping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prettyping_src_url := "https://github.com/denilsonsa/prettyping/archive/refs/tags/v1.0.1.src.tar.gz"
	prettyping_cmd_src := exec.Command("curl", "-L", prettyping_src_url, "-o", "source.tar.gz")
	err = prettyping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prettyping_cmd_direct := exec.Command("./binary")
	err = prettyping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
