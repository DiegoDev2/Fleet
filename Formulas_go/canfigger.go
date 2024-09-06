package main

// Canfigger - Simple configuration file parser library
// Homepage: https://github.com/andy5995/canfigger/

import (
	"fmt"
	
	"os/exec"
)

func installCanfigger() {
	// Método 1: Descargar y extraer .tar.gz
	canfigger_tar_url := "https://github.com/andy5995/canfigger/releases/download/v0.3.0/canfigger-0.3.0.tar.xz"
	canfigger_cmd_tar := exec.Command("curl", "-L", canfigger_tar_url, "-o", "package.tar.gz")
	err := canfigger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	canfigger_zip_url := "https://github.com/andy5995/canfigger/releases/download/v0.3.0/canfigger-0.3.0.tar.xz"
	canfigger_cmd_zip := exec.Command("curl", "-L", canfigger_zip_url, "-o", "package.zip")
	err = canfigger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	canfigger_bin_url := "https://github.com/andy5995/canfigger/releases/download/v0.3.0/canfigger-0.3.0.tar.xz"
	canfigger_cmd_bin := exec.Command("curl", "-L", canfigger_bin_url, "-o", "binary.bin")
	err = canfigger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	canfigger_src_url := "https://github.com/andy5995/canfigger/releases/download/v0.3.0/canfigger-0.3.0.tar.xz"
	canfigger_cmd_src := exec.Command("curl", "-L", canfigger_src_url, "-o", "source.tar.gz")
	err = canfigger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	canfigger_cmd_direct := exec.Command("./binary")
	err = canfigger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
