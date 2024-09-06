package main

// Byobu - Text-based window manager and terminal multiplexer
// Homepage: https://github.com/dustinkirkland/byobu

import (
	"fmt"
	
	"os/exec"
)

func installByobu() {
	// Método 1: Descargar y extraer .tar.gz
	byobu_tar_url := "https://github.com/dustinkirkland/byobu/archive/refs/tags/6.12.tar.gz"
	byobu_cmd_tar := exec.Command("curl", "-L", byobu_tar_url, "-o", "package.tar.gz")
	err := byobu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	byobu_zip_url := "https://github.com/dustinkirkland/byobu/archive/refs/tags/6.12.zip"
	byobu_cmd_zip := exec.Command("curl", "-L", byobu_zip_url, "-o", "package.zip")
	err = byobu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	byobu_bin_url := "https://github.com/dustinkirkland/byobu/archive/refs/tags/6.12.bin"
	byobu_cmd_bin := exec.Command("curl", "-L", byobu_bin_url, "-o", "binary.bin")
	err = byobu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	byobu_src_url := "https://github.com/dustinkirkland/byobu/archive/refs/tags/6.12.src.tar.gz"
	byobu_cmd_src := exec.Command("curl", "-L", byobu_src_url, "-o", "source.tar.gz")
	err = byobu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	byobu_cmd_direct := exec.Command("./binary")
	err = byobu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: newt")
exec.Command("latte", "install", "newt")
	fmt.Println("Instalando dependencia: tmux")
exec.Command("latte", "install", "tmux")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
}
