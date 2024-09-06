package main

// Wordgrinder - Unicode-aware word processor that runs in a terminal
// Homepage: https://cowlark.com/wordgrinder

import (
	"fmt"
	
	"os/exec"
)

func installWordgrinder() {
	// Método 1: Descargar y extraer .tar.gz
	wordgrinder_tar_url := "https://github.com/davidgiven/wordgrinder/archive/refs/tags/0.8.tar.gz"
	wordgrinder_cmd_tar := exec.Command("curl", "-L", wordgrinder_tar_url, "-o", "package.tar.gz")
	err := wordgrinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wordgrinder_zip_url := "https://github.com/davidgiven/wordgrinder/archive/refs/tags/0.8.zip"
	wordgrinder_cmd_zip := exec.Command("curl", "-L", wordgrinder_zip_url, "-o", "package.zip")
	err = wordgrinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wordgrinder_bin_url := "https://github.com/davidgiven/wordgrinder/archive/refs/tags/0.8.bin"
	wordgrinder_cmd_bin := exec.Command("curl", "-L", wordgrinder_bin_url, "-o", "binary.bin")
	err = wordgrinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wordgrinder_src_url := "https://github.com/davidgiven/wordgrinder/archive/refs/tags/0.8.src.tar.gz"
	wordgrinder_cmd_src := exec.Command("curl", "-L", wordgrinder_src_url, "-o", "source.tar.gz")
	err = wordgrinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wordgrinder_cmd_direct := exec.Command("./binary")
	err = wordgrinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
