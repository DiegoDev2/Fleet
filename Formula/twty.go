package main

// Twty - Command-line twitter client written in golang
// Homepage: https://github.com/mattn/twty/

import (
	"fmt"
	
	"os/exec"
)

func installTwty() {
	// Método 1: Descargar y extraer .tar.gz
	twty_tar_url := "https://github.com/mattn/twty/archive/refs/tags/v0.0.13.tar.gz"
	twty_cmd_tar := exec.Command("curl", "-L", twty_tar_url, "-o", "package.tar.gz")
	err := twty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twty_zip_url := "https://github.com/mattn/twty/archive/refs/tags/v0.0.13.zip"
	twty_cmd_zip := exec.Command("curl", "-L", twty_zip_url, "-o", "package.zip")
	err = twty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twty_bin_url := "https://github.com/mattn/twty/archive/refs/tags/v0.0.13.bin"
	twty_cmd_bin := exec.Command("curl", "-L", twty_bin_url, "-o", "binary.bin")
	err = twty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twty_src_url := "https://github.com/mattn/twty/archive/refs/tags/v0.0.13.src.tar.gz"
	twty_cmd_src := exec.Command("curl", "-L", twty_src_url, "-o", "source.tar.gz")
	err = twty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twty_cmd_direct := exec.Command("./binary")
	err = twty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
