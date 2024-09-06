package main

// Grc - Colorize logfiles and command output
// Homepage: https://kassiopeia.juls.savba.sk/~garabik/software/grc.html

import (
	"fmt"
	
	"os/exec"
)

func installGrc() {
	// Método 1: Descargar y extraer .tar.gz
	grc_tar_url := "https://github.com/garabik/grc/archive/refs/tags/v1.13.tar.gz"
	grc_cmd_tar := exec.Command("curl", "-L", grc_tar_url, "-o", "package.tar.gz")
	err := grc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grc_zip_url := "https://github.com/garabik/grc/archive/refs/tags/v1.13.zip"
	grc_cmd_zip := exec.Command("curl", "-L", grc_zip_url, "-o", "package.zip")
	err = grc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grc_bin_url := "https://github.com/garabik/grc/archive/refs/tags/v1.13.bin"
	grc_cmd_bin := exec.Command("curl", "-L", grc_bin_url, "-o", "binary.bin")
	err = grc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grc_src_url := "https://github.com/garabik/grc/archive/refs/tags/v1.13.src.tar.gz"
	grc_cmd_src := exec.Command("curl", "-L", grc_src_url, "-o", "source.tar.gz")
	err = grc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grc_cmd_direct := exec.Command("./binary")
	err = grc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
