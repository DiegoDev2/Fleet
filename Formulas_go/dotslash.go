package main

// Dotslash - Simplified executable deployment
// Homepage: https://dotslash-cli.com

import (
	"fmt"
	
	"os/exec"
)

func installDotslash() {
	// Método 1: Descargar y extraer .tar.gz
	dotslash_tar_url := "https://github.com/facebook/dotslash/archive/refs/tags/v0.4.1.tar.gz"
	dotslash_cmd_tar := exec.Command("curl", "-L", dotslash_tar_url, "-o", "package.tar.gz")
	err := dotslash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dotslash_zip_url := "https://github.com/facebook/dotslash/archive/refs/tags/v0.4.1.zip"
	dotslash_cmd_zip := exec.Command("curl", "-L", dotslash_zip_url, "-o", "package.zip")
	err = dotslash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dotslash_bin_url := "https://github.com/facebook/dotslash/archive/refs/tags/v0.4.1.bin"
	dotslash_cmd_bin := exec.Command("curl", "-L", dotslash_bin_url, "-o", "binary.bin")
	err = dotslash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dotslash_src_url := "https://github.com/facebook/dotslash/archive/refs/tags/v0.4.1.src.tar.gz"
	dotslash_cmd_src := exec.Command("curl", "-L", dotslash_src_url, "-o", "source.tar.gz")
	err = dotslash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dotslash_cmd_direct := exec.Command("./binary")
	err = dotslash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
