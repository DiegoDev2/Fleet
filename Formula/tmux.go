package main

// Tmux - Terminal multiplexer
// Homepage: https://tmux.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installTmux() {
	// Método 1: Descargar y extraer .tar.gz
	tmux_tar_url := "https://github.com/tmux/tmux/releases/download/3.4/tmux-3.4.tar.gz"
	tmux_cmd_tar := exec.Command("curl", "-L", tmux_tar_url, "-o", "package.tar.gz")
	err := tmux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmux_zip_url := "https://github.com/tmux/tmux/releases/download/3.4/tmux-3.4.zip"
	tmux_cmd_zip := exec.Command("curl", "-L", tmux_zip_url, "-o", "package.zip")
	err = tmux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmux_bin_url := "https://github.com/tmux/tmux/releases/download/3.4/tmux-3.4.bin"
	tmux_cmd_bin := exec.Command("curl", "-L", tmux_bin_url, "-o", "binary.bin")
	err = tmux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmux_src_url := "https://github.com/tmux/tmux/releases/download/3.4/tmux-3.4.src.tar.gz"
	tmux_cmd_src := exec.Command("curl", "-L", tmux_src_url, "-o", "source.tar.gz")
	err = tmux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmux_cmd_direct := exec.Command("./binary")
	err = tmux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: utf8proc")
	exec.Command("latte", "install", "utf8proc").Run()
}
