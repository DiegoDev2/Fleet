package main

// Sd - Intuitive find & replace CLI
// Homepage: https://github.com/chmln/sd

import (
	"fmt"
	
	"os/exec"
)

func installSd() {
	// Método 1: Descargar y extraer .tar.gz
	sd_tar_url := "https://github.com/chmln/sd/archive/refs/tags/v1.0.0.tar.gz"
	sd_cmd_tar := exec.Command("curl", "-L", sd_tar_url, "-o", "package.tar.gz")
	err := sd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sd_zip_url := "https://github.com/chmln/sd/archive/refs/tags/v1.0.0.zip"
	sd_cmd_zip := exec.Command("curl", "-L", sd_zip_url, "-o", "package.zip")
	err = sd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sd_bin_url := "https://github.com/chmln/sd/archive/refs/tags/v1.0.0.bin"
	sd_cmd_bin := exec.Command("curl", "-L", sd_bin_url, "-o", "binary.bin")
	err = sd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sd_src_url := "https://github.com/chmln/sd/archive/refs/tags/v1.0.0.src.tar.gz"
	sd_cmd_src := exec.Command("curl", "-L", sd_src_url, "-o", "source.tar.gz")
	err = sd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sd_cmd_direct := exec.Command("./binary")
	err = sd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
