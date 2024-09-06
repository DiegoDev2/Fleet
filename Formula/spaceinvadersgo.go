package main

// SpaceinvadersGo - Space Invaders in your terminal written in Go
// Homepage: https://github.com/asib/spaceinvaders

import (
	"fmt"
	
	"os/exec"
)

func installSpaceinvadersGo() {
	// Método 1: Descargar y extraer .tar.gz
	spaceinvadersgo_tar_url := "https://github.com/asib/spaceinvaders/archive/refs/tags/v1.2.1.tar.gz"
	spaceinvadersgo_cmd_tar := exec.Command("curl", "-L", spaceinvadersgo_tar_url, "-o", "package.tar.gz")
	err := spaceinvadersgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spaceinvadersgo_zip_url := "https://github.com/asib/spaceinvaders/archive/refs/tags/v1.2.1.zip"
	spaceinvadersgo_cmd_zip := exec.Command("curl", "-L", spaceinvadersgo_zip_url, "-o", "package.zip")
	err = spaceinvadersgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spaceinvadersgo_bin_url := "https://github.com/asib/spaceinvaders/archive/refs/tags/v1.2.1.bin"
	spaceinvadersgo_cmd_bin := exec.Command("curl", "-L", spaceinvadersgo_bin_url, "-o", "binary.bin")
	err = spaceinvadersgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spaceinvadersgo_src_url := "https://github.com/asib/spaceinvaders/archive/refs/tags/v1.2.1.src.tar.gz"
	spaceinvadersgo_cmd_src := exec.Command("curl", "-L", spaceinvadersgo_src_url, "-o", "source.tar.gz")
	err = spaceinvadersgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spaceinvadersgo_cmd_direct := exec.Command("./binary")
	err = spaceinvadersgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
