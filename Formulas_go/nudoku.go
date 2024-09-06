package main

// Nudoku - Ncurses based sudoku game
// Homepage: https://jubalh.github.io/nudoku/

import (
	"fmt"
	
	"os/exec"
)

func installNudoku() {
	// Método 1: Descargar y extraer .tar.gz
	nudoku_tar_url := "https://github.com/jubalh/nudoku/archive/refs/tags/5.0.0.tar.gz"
	nudoku_cmd_tar := exec.Command("curl", "-L", nudoku_tar_url, "-o", "package.tar.gz")
	err := nudoku_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nudoku_zip_url := "https://github.com/jubalh/nudoku/archive/refs/tags/5.0.0.zip"
	nudoku_cmd_zip := exec.Command("curl", "-L", nudoku_zip_url, "-o", "package.zip")
	err = nudoku_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nudoku_bin_url := "https://github.com/jubalh/nudoku/archive/refs/tags/5.0.0.bin"
	nudoku_cmd_bin := exec.Command("curl", "-L", nudoku_bin_url, "-o", "binary.bin")
	err = nudoku_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nudoku_src_url := "https://github.com/jubalh/nudoku/archive/refs/tags/5.0.0.src.tar.gz"
	nudoku_cmd_src := exec.Command("curl", "-L", nudoku_src_url, "-o", "source.tar.gz")
	err = nudoku_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nudoku_cmd_direct := exec.Command("./binary")
	err = nudoku_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
