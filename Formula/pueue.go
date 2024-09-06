package main

// Pueue - Command-line tool for managing long-running shell commands
// Homepage: https://github.com/Nukesor/pueue

import (
	"fmt"
	
	"os/exec"
)

func installPueue() {
	// Método 1: Descargar y extraer .tar.gz
	pueue_tar_url := "https://github.com/Nukesor/pueue/archive/refs/tags/v3.4.1.tar.gz"
	pueue_cmd_tar := exec.Command("curl", "-L", pueue_tar_url, "-o", "package.tar.gz")
	err := pueue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pueue_zip_url := "https://github.com/Nukesor/pueue/archive/refs/tags/v3.4.1.zip"
	pueue_cmd_zip := exec.Command("curl", "-L", pueue_zip_url, "-o", "package.zip")
	err = pueue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pueue_bin_url := "https://github.com/Nukesor/pueue/archive/refs/tags/v3.4.1.bin"
	pueue_cmd_bin := exec.Command("curl", "-L", pueue_bin_url, "-o", "binary.bin")
	err = pueue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pueue_src_url := "https://github.com/Nukesor/pueue/archive/refs/tags/v3.4.1.src.tar.gz"
	pueue_cmd_src := exec.Command("curl", "-L", pueue_src_url, "-o", "source.tar.gz")
	err = pueue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pueue_cmd_direct := exec.Command("./binary")
	err = pueue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
