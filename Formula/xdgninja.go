package main

// XdgNinja - Check your $HOME for unwanted files and directories
// Homepage: https://github.com/b3nj5m1n/xdg-ninja

import (
	"fmt"
	
	"os/exec"
)

func installXdgNinja() {
	// Método 1: Descargar y extraer .tar.gz
	xdgninja_tar_url := "https://github.com/b3nj5m1n/xdg-ninja/archive/refs/tags/v0.2.0.2.tar.gz"
	xdgninja_cmd_tar := exec.Command("curl", "-L", xdgninja_tar_url, "-o", "package.tar.gz")
	err := xdgninja_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xdgninja_zip_url := "https://github.com/b3nj5m1n/xdg-ninja/archive/refs/tags/v0.2.0.2.zip"
	xdgninja_cmd_zip := exec.Command("curl", "-L", xdgninja_zip_url, "-o", "package.zip")
	err = xdgninja_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xdgninja_bin_url := "https://github.com/b3nj5m1n/xdg-ninja/archive/refs/tags/v0.2.0.2.bin"
	xdgninja_cmd_bin := exec.Command("curl", "-L", xdgninja_bin_url, "-o", "binary.bin")
	err = xdgninja_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xdgninja_src_url := "https://github.com/b3nj5m1n/xdg-ninja/archive/refs/tags/v0.2.0.2.src.tar.gz"
	xdgninja_cmd_src := exec.Command("curl", "-L", xdgninja_src_url, "-o", "source.tar.gz")
	err = xdgninja_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xdgninja_cmd_direct := exec.Command("./binary")
	err = xdgninja_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glow")
	exec.Command("latte", "install", "glow").Run()
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
}
