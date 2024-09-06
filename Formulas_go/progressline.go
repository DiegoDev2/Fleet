package main

// Progressline - Track commands progress in a compact one-line format
// Homepage: https://github.com/kattouf/ProgressLine

import (
	"fmt"
	
	"os/exec"
)

func installProgressline() {
	// Método 1: Descargar y extraer .tar.gz
	progressline_tar_url := "https://github.com/kattouf/ProgressLine/archive/refs/tags/0.2.2.tar.gz"
	progressline_cmd_tar := exec.Command("curl", "-L", progressline_tar_url, "-o", "package.tar.gz")
	err := progressline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	progressline_zip_url := "https://github.com/kattouf/ProgressLine/archive/refs/tags/0.2.2.zip"
	progressline_cmd_zip := exec.Command("curl", "-L", progressline_zip_url, "-o", "package.zip")
	err = progressline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	progressline_bin_url := "https://github.com/kattouf/ProgressLine/archive/refs/tags/0.2.2.bin"
	progressline_cmd_bin := exec.Command("curl", "-L", progressline_bin_url, "-o", "binary.bin")
	err = progressline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	progressline_src_url := "https://github.com/kattouf/ProgressLine/archive/refs/tags/0.2.2.src.tar.gz"
	progressline_cmd_src := exec.Command("curl", "-L", progressline_src_url, "-o", "source.tar.gz")
	err = progressline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	progressline_cmd_direct := exec.Command("./binary")
	err = progressline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
