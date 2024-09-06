package main

// Halp - CLI tool to get help with CLI tools
// Homepage: https://halp.cli.rs/

import (
	"fmt"
	
	"os/exec"
)

func installHalp() {
	// Método 1: Descargar y extraer .tar.gz
	halp_tar_url := "https://github.com/orhun/halp/archive/refs/tags/v0.2.0.tar.gz"
	halp_cmd_tar := exec.Command("curl", "-L", halp_tar_url, "-o", "package.tar.gz")
	err := halp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	halp_zip_url := "https://github.com/orhun/halp/archive/refs/tags/v0.2.0.zip"
	halp_cmd_zip := exec.Command("curl", "-L", halp_zip_url, "-o", "package.zip")
	err = halp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	halp_bin_url := "https://github.com/orhun/halp/archive/refs/tags/v0.2.0.bin"
	halp_cmd_bin := exec.Command("curl", "-L", halp_bin_url, "-o", "binary.bin")
	err = halp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	halp_src_url := "https://github.com/orhun/halp/archive/refs/tags/v0.2.0.src.tar.gz"
	halp_cmd_src := exec.Command("curl", "-L", halp_src_url, "-o", "source.tar.gz")
	err = halp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	halp_cmd_direct := exec.Command("./binary")
	err = halp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
