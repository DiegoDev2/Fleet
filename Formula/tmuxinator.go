package main

// Tmuxinator - Manage complex tmux sessions easily
// Homepage: https://github.com/tmuxinator/tmuxinator

import (
	"fmt"
	
	"os/exec"
)

func installTmuxinator() {
	// Método 1: Descargar y extraer .tar.gz
	tmuxinator_tar_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.tar.gz"
	tmuxinator_cmd_tar := exec.Command("curl", "-L", tmuxinator_tar_url, "-o", "package.tar.gz")
	err := tmuxinator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmuxinator_zip_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.zip"
	tmuxinator_cmd_zip := exec.Command("curl", "-L", tmuxinator_zip_url, "-o", "package.zip")
	err = tmuxinator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmuxinator_bin_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.bin"
	tmuxinator_cmd_bin := exec.Command("curl", "-L", tmuxinator_bin_url, "-o", "binary.bin")
	err = tmuxinator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmuxinator_src_url := "https://github.com/tmuxinator/tmuxinator/archive/refs/tags/v3.3.0.src.tar.gz"
	tmuxinator_cmd_src := exec.Command("curl", "-L", tmuxinator_src_url, "-o", "source.tar.gz")
	err = tmuxinator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmuxinator_cmd_direct := exec.Command("./binary")
	err = tmuxinator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
	fmt.Println("Instalando dependencia: tmux")
	exec.Command("latte", "install", "tmux").Run()
}
