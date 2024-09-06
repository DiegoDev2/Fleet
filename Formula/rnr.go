package main

// Rnr - Command-line tool to batch rename files and directories
// Homepage: https://github.com/ismaelgv/rnr

import (
	"fmt"
	
	"os/exec"
)

func installRnr() {
	// Método 1: Descargar y extraer .tar.gz
	rnr_tar_url := "https://github.com/ismaelgv/rnr/archive/refs/tags/v0.4.2.tar.gz"
	rnr_cmd_tar := exec.Command("curl", "-L", rnr_tar_url, "-o", "package.tar.gz")
	err := rnr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rnr_zip_url := "https://github.com/ismaelgv/rnr/archive/refs/tags/v0.4.2.zip"
	rnr_cmd_zip := exec.Command("curl", "-L", rnr_zip_url, "-o", "package.zip")
	err = rnr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rnr_bin_url := "https://github.com/ismaelgv/rnr/archive/refs/tags/v0.4.2.bin"
	rnr_cmd_bin := exec.Command("curl", "-L", rnr_bin_url, "-o", "binary.bin")
	err = rnr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rnr_src_url := "https://github.com/ismaelgv/rnr/archive/refs/tags/v0.4.2.src.tar.gz"
	rnr_cmd_src := exec.Command("curl", "-L", rnr_src_url, "-o", "source.tar.gz")
	err = rnr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rnr_cmd_direct := exec.Command("./binary")
	err = rnr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
