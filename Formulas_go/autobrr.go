package main

// Autobrr - Modern, easy to use download automation for torrents and usenet
// Homepage: https://autobrr.com/

import (
	"fmt"
	
	"os/exec"
)

func installAutobrr() {
	// Método 1: Descargar y extraer .tar.gz
	autobrr_tar_url := "https://github.com/autobrr/autobrr/archive/refs/tags/v1.45.0.tar.gz"
	autobrr_cmd_tar := exec.Command("curl", "-L", autobrr_tar_url, "-o", "package.tar.gz")
	err := autobrr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autobrr_zip_url := "https://github.com/autobrr/autobrr/archive/refs/tags/v1.45.0.zip"
	autobrr_cmd_zip := exec.Command("curl", "-L", autobrr_zip_url, "-o", "package.zip")
	err = autobrr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autobrr_bin_url := "https://github.com/autobrr/autobrr/archive/refs/tags/v1.45.0.bin"
	autobrr_cmd_bin := exec.Command("curl", "-L", autobrr_bin_url, "-o", "binary.bin")
	err = autobrr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autobrr_src_url := "https://github.com/autobrr/autobrr/archive/refs/tags/v1.45.0.src.tar.gz"
	autobrr_cmd_src := exec.Command("curl", "-L", autobrr_src_url, "-o", "source.tar.gz")
	err = autobrr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autobrr_cmd_direct := exec.Command("./binary")
	err = autobrr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: pnpm")
exec.Command("latte", "install", "pnpm")
}
