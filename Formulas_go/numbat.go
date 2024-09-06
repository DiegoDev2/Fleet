package main

// Numbat - Statically typed programming language for scientific computations
// Homepage: https://github.com/sharkdp/numbat

import (
	"fmt"
	
	"os/exec"
)

func installNumbat() {
	// Método 1: Descargar y extraer .tar.gz
	numbat_tar_url := "https://github.com/sharkdp/numbat/archive/refs/tags/v1.13.0.tar.gz"
	numbat_cmd_tar := exec.Command("curl", "-L", numbat_tar_url, "-o", "package.tar.gz")
	err := numbat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	numbat_zip_url := "https://github.com/sharkdp/numbat/archive/refs/tags/v1.13.0.zip"
	numbat_cmd_zip := exec.Command("curl", "-L", numbat_zip_url, "-o", "package.zip")
	err = numbat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	numbat_bin_url := "https://github.com/sharkdp/numbat/archive/refs/tags/v1.13.0.bin"
	numbat_cmd_bin := exec.Command("curl", "-L", numbat_bin_url, "-o", "binary.bin")
	err = numbat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	numbat_src_url := "https://github.com/sharkdp/numbat/archive/refs/tags/v1.13.0.src.tar.gz"
	numbat_cmd_src := exec.Command("curl", "-L", numbat_src_url, "-o", "source.tar.gz")
	err = numbat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	numbat_cmd_direct := exec.Command("./binary")
	err = numbat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
