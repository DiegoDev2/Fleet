package main

// Riff - Diff filter highlighting which line parts have changed
// Homepage: https://github.com/walles/riff

import (
	"fmt"
	
	"os/exec"
)

func installRiff() {
	// Método 1: Descargar y extraer .tar.gz
	riff_tar_url := "https://github.com/walles/riff/archive/refs/tags/3.2.0.tar.gz"
	riff_cmd_tar := exec.Command("curl", "-L", riff_tar_url, "-o", "package.tar.gz")
	err := riff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	riff_zip_url := "https://github.com/walles/riff/archive/refs/tags/3.2.0.zip"
	riff_cmd_zip := exec.Command("curl", "-L", riff_zip_url, "-o", "package.zip")
	err = riff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	riff_bin_url := "https://github.com/walles/riff/archive/refs/tags/3.2.0.bin"
	riff_cmd_bin := exec.Command("curl", "-L", riff_bin_url, "-o", "binary.bin")
	err = riff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	riff_src_url := "https://github.com/walles/riff/archive/refs/tags/3.2.0.src.tar.gz"
	riff_cmd_src := exec.Command("curl", "-L", riff_src_url, "-o", "source.tar.gz")
	err = riff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	riff_cmd_direct := exec.Command("./binary")
	err = riff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
