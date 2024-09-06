package main

// Yank - Copy terminal output to clipboard
// Homepage: https://github.com/mptre/yank

import (
	"fmt"
	
	"os/exec"
)

func installYank() {
	// Método 1: Descargar y extraer .tar.gz
	yank_tar_url := "https://github.com/mptre/yank/archive/refs/tags/v1.3.0.tar.gz"
	yank_cmd_tar := exec.Command("curl", "-L", yank_tar_url, "-o", "package.tar.gz")
	err := yank_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yank_zip_url := "https://github.com/mptre/yank/archive/refs/tags/v1.3.0.zip"
	yank_cmd_zip := exec.Command("curl", "-L", yank_zip_url, "-o", "package.zip")
	err = yank_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yank_bin_url := "https://github.com/mptre/yank/archive/refs/tags/v1.3.0.bin"
	yank_cmd_bin := exec.Command("curl", "-L", yank_bin_url, "-o", "binary.bin")
	err = yank_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yank_src_url := "https://github.com/mptre/yank/archive/refs/tags/v1.3.0.src.tar.gz"
	yank_cmd_src := exec.Command("curl", "-L", yank_src_url, "-o", "source.tar.gz")
	err = yank_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yank_cmd_direct := exec.Command("./binary")
	err = yank_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
