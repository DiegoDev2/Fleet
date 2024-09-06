package main

// Topgrade - Upgrade all the things
// Homepage: https://github.com/topgrade-rs/topgrade

import (
	"fmt"
	
	"os/exec"
)

func installTopgrade() {
	// Método 1: Descargar y extraer .tar.gz
	topgrade_tar_url := "https://github.com/topgrade-rs/topgrade/archive/refs/tags/v15.0.0.tar.gz"
	topgrade_cmd_tar := exec.Command("curl", "-L", topgrade_tar_url, "-o", "package.tar.gz")
	err := topgrade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	topgrade_zip_url := "https://github.com/topgrade-rs/topgrade/archive/refs/tags/v15.0.0.zip"
	topgrade_cmd_zip := exec.Command("curl", "-L", topgrade_zip_url, "-o", "package.zip")
	err = topgrade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	topgrade_bin_url := "https://github.com/topgrade-rs/topgrade/archive/refs/tags/v15.0.0.bin"
	topgrade_cmd_bin := exec.Command("curl", "-L", topgrade_bin_url, "-o", "binary.bin")
	err = topgrade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	topgrade_src_url := "https://github.com/topgrade-rs/topgrade/archive/refs/tags/v15.0.0.src.tar.gz"
	topgrade_cmd_src := exec.Command("curl", "-L", topgrade_src_url, "-o", "source.tar.gz")
	err = topgrade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	topgrade_cmd_direct := exec.Command("./binary")
	err = topgrade_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
