package main

// Matlab2tikz - Convert MATLAB(R) figures into TikZ/Pgfplots figures
// Homepage: https://github.com/matlab2tikz/matlab2tikz

import (
	"fmt"
	
	"os/exec"
)

func installMatlab2tikz() {
	// Método 1: Descargar y extraer .tar.gz
	matlab2tikz_tar_url := "https://github.com/matlab2tikz/matlab2tikz/archive/refs/tags/v1.1.0.tar.gz"
	matlab2tikz_cmd_tar := exec.Command("curl", "-L", matlab2tikz_tar_url, "-o", "package.tar.gz")
	err := matlab2tikz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	matlab2tikz_zip_url := "https://github.com/matlab2tikz/matlab2tikz/archive/refs/tags/v1.1.0.zip"
	matlab2tikz_cmd_zip := exec.Command("curl", "-L", matlab2tikz_zip_url, "-o", "package.zip")
	err = matlab2tikz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	matlab2tikz_bin_url := "https://github.com/matlab2tikz/matlab2tikz/archive/refs/tags/v1.1.0.bin"
	matlab2tikz_cmd_bin := exec.Command("curl", "-L", matlab2tikz_bin_url, "-o", "binary.bin")
	err = matlab2tikz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	matlab2tikz_src_url := "https://github.com/matlab2tikz/matlab2tikz/archive/refs/tags/v1.1.0.src.tar.gz"
	matlab2tikz_cmd_src := exec.Command("curl", "-L", matlab2tikz_src_url, "-o", "source.tar.gz")
	err = matlab2tikz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	matlab2tikz_cmd_direct := exec.Command("./binary")
	err = matlab2tikz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
