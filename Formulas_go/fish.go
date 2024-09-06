package main

// Fish - User-friendly command-line shell for UNIX-like operating systems
// Homepage: https://fishshell.com

import (
	"fmt"
	
	"os/exec"
)

func installFish() {
	// Método 1: Descargar y extraer .tar.gz
	fish_tar_url := "https://github.com/fish-shell/fish-shell/releases/download/3.7.1/fish-3.7.1.tar.xz"
	fish_cmd_tar := exec.Command("curl", "-L", fish_tar_url, "-o", "package.tar.gz")
	err := fish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fish_zip_url := "https://github.com/fish-shell/fish-shell/releases/download/3.7.1/fish-3.7.1.tar.xz"
	fish_cmd_zip := exec.Command("curl", "-L", fish_zip_url, "-o", "package.zip")
	err = fish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fish_bin_url := "https://github.com/fish-shell/fish-shell/releases/download/3.7.1/fish-3.7.1.tar.xz"
	fish_cmd_bin := exec.Command("curl", "-L", fish_bin_url, "-o", "binary.bin")
	err = fish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fish_src_url := "https://github.com/fish-shell/fish-shell/releases/download/3.7.1/fish-3.7.1.tar.xz"
	fish_cmd_src := exec.Command("curl", "-L", fish_src_url, "-o", "source.tar.gz")
	err = fish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fish_cmd_direct := exec.Command("./binary")
	err = fish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
