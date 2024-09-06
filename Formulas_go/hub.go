package main

// Hub - Add GitHub support to git on the command-line
// Homepage: https://hub.github.com/

import (
	"fmt"
	
	"os/exec"
)

func installHub() {
	// Método 1: Descargar y extraer .tar.gz
	hub_tar_url := "https://github.com/github/hub/archive/refs/tags/v2.14.2.tar.gz"
	hub_cmd_tar := exec.Command("curl", "-L", hub_tar_url, "-o", "package.tar.gz")
	err := hub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hub_zip_url := "https://github.com/github/hub/archive/refs/tags/v2.14.2.zip"
	hub_cmd_zip := exec.Command("curl", "-L", hub_zip_url, "-o", "package.zip")
	err = hub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hub_bin_url := "https://github.com/github/hub/archive/refs/tags/v2.14.2.bin"
	hub_cmd_bin := exec.Command("curl", "-L", hub_bin_url, "-o", "binary.bin")
	err = hub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hub_src_url := "https://github.com/github/hub/archive/refs/tags/v2.14.2.src.tar.gz"
	hub_cmd_src := exec.Command("curl", "-L", hub_src_url, "-o", "source.tar.gz")
	err = hub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hub_cmd_direct := exec.Command("./binary")
	err = hub_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: groff")
exec.Command("latte", "install", "groff")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
