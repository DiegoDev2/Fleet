package main

// Fnt - Apt for fonts, the missing font manager for macOS/linux
// Homepage: https://github.com/alexmyczko/fnt

import (
	"fmt"
	
	"os/exec"
)

func installFnt() {
	// Método 1: Descargar y extraer .tar.gz
	fnt_tar_url := "https://github.com/alexmyczko/fnt/archive/refs/tags/1.6.tar.gz"
	fnt_cmd_tar := exec.Command("curl", "-L", fnt_tar_url, "-o", "package.tar.gz")
	err := fnt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fnt_zip_url := "https://github.com/alexmyczko/fnt/archive/refs/tags/1.6.zip"
	fnt_cmd_zip := exec.Command("curl", "-L", fnt_zip_url, "-o", "package.zip")
	err = fnt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fnt_bin_url := "https://github.com/alexmyczko/fnt/archive/refs/tags/1.6.bin"
	fnt_cmd_bin := exec.Command("curl", "-L", fnt_bin_url, "-o", "binary.bin")
	err = fnt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fnt_src_url := "https://github.com/alexmyczko/fnt/archive/refs/tags/1.6.src.tar.gz"
	fnt_cmd_src := exec.Command("curl", "-L", fnt_src_url, "-o", "source.tar.gz")
	err = fnt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fnt_cmd_direct := exec.Command("./binary")
	err = fnt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: chafa")
	exec.Command("latte", "install", "chafa").Run()
	fmt.Println("Instalando dependencia: lcdf-typetools")
	exec.Command("latte", "install", "lcdf-typetools").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
