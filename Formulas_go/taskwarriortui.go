package main

// TaskwarriorTui - Terminal user interface for taskwarrior
// Homepage: https://github.com/kdheepak/taskwarrior-tui

import (
	"fmt"
	
	"os/exec"
)

func installTaskwarriorTui() {
	// Método 1: Descargar y extraer .tar.gz
	taskwarriortui_tar_url := "https://github.com/kdheepak/taskwarrior-tui/archive/refs/tags/v0.26.3.tar.gz"
	taskwarriortui_cmd_tar := exec.Command("curl", "-L", taskwarriortui_tar_url, "-o", "package.tar.gz")
	err := taskwarriortui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taskwarriortui_zip_url := "https://github.com/kdheepak/taskwarrior-tui/archive/refs/tags/v0.26.3.zip"
	taskwarriortui_cmd_zip := exec.Command("curl", "-L", taskwarriortui_zip_url, "-o", "package.zip")
	err = taskwarriortui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taskwarriortui_bin_url := "https://github.com/kdheepak/taskwarrior-tui/archive/refs/tags/v0.26.3.bin"
	taskwarriortui_cmd_bin := exec.Command("curl", "-L", taskwarriortui_bin_url, "-o", "binary.bin")
	err = taskwarriortui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taskwarriortui_src_url := "https://github.com/kdheepak/taskwarrior-tui/archive/refs/tags/v0.26.3.src.tar.gz"
	taskwarriortui_cmd_src := exec.Command("curl", "-L", taskwarriortui_src_url, "-o", "source.tar.gz")
	err = taskwarriortui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taskwarriortui_cmd_direct := exec.Command("./binary")
	err = taskwarriortui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: task")
exec.Command("latte", "install", "task")
}
