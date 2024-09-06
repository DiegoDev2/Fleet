package main

// Pastel - Command-line tool to generate, analyze, convert and manipulate colors
// Homepage: https://github.com/sharkdp/pastel

import (
	"fmt"
	
	"os/exec"
)

func installPastel() {
	// Método 1: Descargar y extraer .tar.gz
	pastel_tar_url := "https://github.com/sharkdp/pastel/archive/refs/tags/v0.9.0.tar.gz"
	pastel_cmd_tar := exec.Command("curl", "-L", pastel_tar_url, "-o", "package.tar.gz")
	err := pastel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pastel_zip_url := "https://github.com/sharkdp/pastel/archive/refs/tags/v0.9.0.zip"
	pastel_cmd_zip := exec.Command("curl", "-L", pastel_zip_url, "-o", "package.zip")
	err = pastel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pastel_bin_url := "https://github.com/sharkdp/pastel/archive/refs/tags/v0.9.0.bin"
	pastel_cmd_bin := exec.Command("curl", "-L", pastel_bin_url, "-o", "binary.bin")
	err = pastel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pastel_src_url := "https://github.com/sharkdp/pastel/archive/refs/tags/v0.9.0.src.tar.gz"
	pastel_cmd_src := exec.Command("curl", "-L", pastel_src_url, "-o", "source.tar.gz")
	err = pastel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pastel_cmd_direct := exec.Command("./binary")
	err = pastel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
