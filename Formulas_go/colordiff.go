package main

// Colordiff - Color-highlighted diff(1) output
// Homepage: https://www.colordiff.org/

import (
	"fmt"
	
	"os/exec"
)

func installColordiff() {
	// Método 1: Descargar y extraer .tar.gz
	colordiff_tar_url := "https://www.colordiff.org/colordiff-1.0.21.tar.gz"
	colordiff_cmd_tar := exec.Command("curl", "-L", colordiff_tar_url, "-o", "package.tar.gz")
	err := colordiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colordiff_zip_url := "https://www.colordiff.org/colordiff-1.0.21.zip"
	colordiff_cmd_zip := exec.Command("curl", "-L", colordiff_zip_url, "-o", "package.zip")
	err = colordiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colordiff_bin_url := "https://www.colordiff.org/colordiff-1.0.21.bin"
	colordiff_cmd_bin := exec.Command("curl", "-L", colordiff_bin_url, "-o", "binary.bin")
	err = colordiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colordiff_src_url := "https://www.colordiff.org/colordiff-1.0.21.src.tar.gz"
	colordiff_cmd_src := exec.Command("curl", "-L", colordiff_src_url, "-o", "source.tar.gz")
	err = colordiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colordiff_cmd_direct := exec.Command("./binary")
	err = colordiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
}
