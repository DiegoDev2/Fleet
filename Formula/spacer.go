package main

// Spacer - Small command-line utility for adding spacers to command output
// Homepage: https://github.com/samwho/spacer

import (
	"fmt"
	
	"os/exec"
)

func installSpacer() {
	// Método 1: Descargar y extraer .tar.gz
	spacer_tar_url := "https://github.com/samwho/spacer/archive/refs/tags/v0.3.0.tar.gz"
	spacer_cmd_tar := exec.Command("curl", "-L", spacer_tar_url, "-o", "package.tar.gz")
	err := spacer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spacer_zip_url := "https://github.com/samwho/spacer/archive/refs/tags/v0.3.0.zip"
	spacer_cmd_zip := exec.Command("curl", "-L", spacer_zip_url, "-o", "package.zip")
	err = spacer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spacer_bin_url := "https://github.com/samwho/spacer/archive/refs/tags/v0.3.0.bin"
	spacer_cmd_bin := exec.Command("curl", "-L", spacer_bin_url, "-o", "binary.bin")
	err = spacer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spacer_src_url := "https://github.com/samwho/spacer/archive/refs/tags/v0.3.0.src.tar.gz"
	spacer_cmd_src := exec.Command("curl", "-L", spacer_src_url, "-o", "source.tar.gz")
	err = spacer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spacer_cmd_direct := exec.Command("./binary")
	err = spacer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
