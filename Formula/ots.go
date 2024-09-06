package main

// Ots - Share end-to-end encrypted secrets with others via a one-time URL
// Homepage: https://ots.sniptt.com

import (
	"fmt"
	
	"os/exec"
)

func installOts() {
	// Método 1: Descargar y extraer .tar.gz
	ots_tar_url := "https://github.com/sniptt-official/ots/archive/refs/tags/v0.2.0.tar.gz"
	ots_cmd_tar := exec.Command("curl", "-L", ots_tar_url, "-o", "package.tar.gz")
	err := ots_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ots_zip_url := "https://github.com/sniptt-official/ots/archive/refs/tags/v0.2.0.zip"
	ots_cmd_zip := exec.Command("curl", "-L", ots_zip_url, "-o", "package.zip")
	err = ots_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ots_bin_url := "https://github.com/sniptt-official/ots/archive/refs/tags/v0.2.0.bin"
	ots_cmd_bin := exec.Command("curl", "-L", ots_bin_url, "-o", "binary.bin")
	err = ots_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ots_src_url := "https://github.com/sniptt-official/ots/archive/refs/tags/v0.2.0.src.tar.gz"
	ots_cmd_src := exec.Command("curl", "-L", ots_src_url, "-o", "source.tar.gz")
	err = ots_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ots_cmd_direct := exec.Command("./binary")
	err = ots_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
