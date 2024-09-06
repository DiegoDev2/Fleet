package main

// ZshNavigationTools - Zsh curses-based tools, e.g. multi-word history searcher
// Homepage: https://github.com/z-shell/zsh-navigation-tools

import (
	"fmt"
	
	"os/exec"
)

func installZshNavigationTools() {
	// Método 1: Descargar y extraer .tar.gz
	zshnavigationtools_tar_url := "https://github.com/z-shell/zsh-navigation-tools/archive/refs/tags/v2.2.7.tar.gz"
	zshnavigationtools_cmd_tar := exec.Command("curl", "-L", zshnavigationtools_tar_url, "-o", "package.tar.gz")
	err := zshnavigationtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshnavigationtools_zip_url := "https://github.com/z-shell/zsh-navigation-tools/archive/refs/tags/v2.2.7.zip"
	zshnavigationtools_cmd_zip := exec.Command("curl", "-L", zshnavigationtools_zip_url, "-o", "package.zip")
	err = zshnavigationtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshnavigationtools_bin_url := "https://github.com/z-shell/zsh-navigation-tools/archive/refs/tags/v2.2.7.bin"
	zshnavigationtools_cmd_bin := exec.Command("curl", "-L", zshnavigationtools_bin_url, "-o", "binary.bin")
	err = zshnavigationtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshnavigationtools_src_url := "https://github.com/z-shell/zsh-navigation-tools/archive/refs/tags/v2.2.7.src.tar.gz"
	zshnavigationtools_cmd_src := exec.Command("curl", "-L", zshnavigationtools_src_url, "-o", "source.tar.gz")
	err = zshnavigationtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshnavigationtools_cmd_direct := exec.Command("./binary")
	err = zshnavigationtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
