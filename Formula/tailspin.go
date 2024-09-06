package main

// Tailspin - Log file highlighter
// Homepage: https://github.com/bensadeh/tailspin

import (
	"fmt"
	
	"os/exec"
)

func installTailspin() {
	// Método 1: Descargar y extraer .tar.gz
	tailspin_tar_url := "https://github.com/bensadeh/tailspin/archive/refs/tags/3.0.2.tar.gz"
	tailspin_cmd_tar := exec.Command("curl", "-L", tailspin_tar_url, "-o", "package.tar.gz")
	err := tailspin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tailspin_zip_url := "https://github.com/bensadeh/tailspin/archive/refs/tags/3.0.2.zip"
	tailspin_cmd_zip := exec.Command("curl", "-L", tailspin_zip_url, "-o", "package.zip")
	err = tailspin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tailspin_bin_url := "https://github.com/bensadeh/tailspin/archive/refs/tags/3.0.2.bin"
	tailspin_cmd_bin := exec.Command("curl", "-L", tailspin_bin_url, "-o", "binary.bin")
	err = tailspin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tailspin_src_url := "https://github.com/bensadeh/tailspin/archive/refs/tags/3.0.2.src.tar.gz"
	tailspin_cmd_src := exec.Command("curl", "-L", tailspin_src_url, "-o", "source.tar.gz")
	err = tailspin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tailspin_cmd_direct := exec.Command("./binary")
	err = tailspin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
