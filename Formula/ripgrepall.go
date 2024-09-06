package main

// RipgrepAll - Wrapper around ripgrep that adds multiple rich file types
// Homepage: https://github.com/phiresky/ripgrep-all

import (
	"fmt"
	
	"os/exec"
)

func installRipgrepAll() {
	// Método 1: Descargar y extraer .tar.gz
	ripgrepall_tar_url := "https://github.com/phiresky/ripgrep-all/archive/refs/tags/v0.10.6.tar.gz"
	ripgrepall_cmd_tar := exec.Command("curl", "-L", ripgrepall_tar_url, "-o", "package.tar.gz")
	err := ripgrepall_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ripgrepall_zip_url := "https://github.com/phiresky/ripgrep-all/archive/refs/tags/v0.10.6.zip"
	ripgrepall_cmd_zip := exec.Command("curl", "-L", ripgrepall_zip_url, "-o", "package.zip")
	err = ripgrepall_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ripgrepall_bin_url := "https://github.com/phiresky/ripgrep-all/archive/refs/tags/v0.10.6.bin"
	ripgrepall_cmd_bin := exec.Command("curl", "-L", ripgrepall_bin_url, "-o", "binary.bin")
	err = ripgrepall_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ripgrepall_src_url := "https://github.com/phiresky/ripgrep-all/archive/refs/tags/v0.10.6.src.tar.gz"
	ripgrepall_cmd_src := exec.Command("curl", "-L", ripgrepall_src_url, "-o", "source.tar.gz")
	err = ripgrepall_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ripgrepall_cmd_direct := exec.Command("./binary")
	err = ripgrepall_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: ripgrep")
	exec.Command("latte", "install", "ripgrep").Run()
}
