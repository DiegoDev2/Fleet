package main

// Taskell - Command-line Kanban board/task manager with support for Trello
// Homepage: https://taskell.app

import (
	"fmt"
	
	"os/exec"
)

func installTaskell() {
	// Método 1: Descargar y extraer .tar.gz
	taskell_tar_url := "https://github.com/smallhadroncollider/taskell/archive/refs/tags/1.11.4.tar.gz"
	taskell_cmd_tar := exec.Command("curl", "-L", taskell_tar_url, "-o", "package.tar.gz")
	err := taskell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taskell_zip_url := "https://github.com/smallhadroncollider/taskell/archive/refs/tags/1.11.4.zip"
	taskell_cmd_zip := exec.Command("curl", "-L", taskell_zip_url, "-o", "package.zip")
	err = taskell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taskell_bin_url := "https://github.com/smallhadroncollider/taskell/archive/refs/tags/1.11.4.bin"
	taskell_cmd_bin := exec.Command("curl", "-L", taskell_bin_url, "-o", "binary.bin")
	err = taskell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taskell_src_url := "https://github.com/smallhadroncollider/taskell/archive/refs/tags/1.11.4.src.tar.gz"
	taskell_cmd_src := exec.Command("curl", "-L", taskell_src_url, "-o", "source.tar.gz")
	err = taskell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taskell_cmd_direct := exec.Command("./binary")
	err = taskell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.2")
exec.Command("latte", "install", "ghc@9.2")
	fmt.Println("Instalando dependencia: hpack")
exec.Command("latte", "install", "hpack")
}
