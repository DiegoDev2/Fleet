package main

// TmuxXpanes - Ultimate terminal divider powered by tmux
// Homepage: https://github.com/greymd/tmux-xpanes

import (
	"fmt"
	
	"os/exec"
)

func installTmuxXpanes() {
	// Método 1: Descargar y extraer .tar.gz
	tmuxxpanes_tar_url := "https://github.com/greymd/tmux-xpanes/archive/refs/tags/v4.2.0.tar.gz"
	tmuxxpanes_cmd_tar := exec.Command("curl", "-L", tmuxxpanes_tar_url, "-o", "package.tar.gz")
	err := tmuxxpanes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmuxxpanes_zip_url := "https://github.com/greymd/tmux-xpanes/archive/refs/tags/v4.2.0.zip"
	tmuxxpanes_cmd_zip := exec.Command("curl", "-L", tmuxxpanes_zip_url, "-o", "package.zip")
	err = tmuxxpanes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmuxxpanes_bin_url := "https://github.com/greymd/tmux-xpanes/archive/refs/tags/v4.2.0.bin"
	tmuxxpanes_cmd_bin := exec.Command("curl", "-L", tmuxxpanes_bin_url, "-o", "binary.bin")
	err = tmuxxpanes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmuxxpanes_src_url := "https://github.com/greymd/tmux-xpanes/archive/refs/tags/v4.2.0.src.tar.gz"
	tmuxxpanes_cmd_src := exec.Command("curl", "-L", tmuxxpanes_src_url, "-o", "source.tar.gz")
	err = tmuxxpanes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmuxxpanes_cmd_direct := exec.Command("./binary")
	err = tmuxxpanes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tmux")
exec.Command("latte", "install", "tmux")
}
