package main

// TtySolitaire - Ncurses-based klondike solitaire game
// Homepage: https://github.com/mpereira/tty-solitaire

import (
	"fmt"
	
	"os/exec"
)

func installTtySolitaire() {
	// Método 1: Descargar y extraer .tar.gz
	ttysolitaire_tar_url := "https://github.com/mpereira/tty-solitaire/archive/refs/tags/v1.3.1.tar.gz"
	ttysolitaire_cmd_tar := exec.Command("curl", "-L", ttysolitaire_tar_url, "-o", "package.tar.gz")
	err := ttysolitaire_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttysolitaire_zip_url := "https://github.com/mpereira/tty-solitaire/archive/refs/tags/v1.3.1.zip"
	ttysolitaire_cmd_zip := exec.Command("curl", "-L", ttysolitaire_zip_url, "-o", "package.zip")
	err = ttysolitaire_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttysolitaire_bin_url := "https://github.com/mpereira/tty-solitaire/archive/refs/tags/v1.3.1.bin"
	ttysolitaire_cmd_bin := exec.Command("curl", "-L", ttysolitaire_bin_url, "-o", "binary.bin")
	err = ttysolitaire_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttysolitaire_src_url := "https://github.com/mpereira/tty-solitaire/archive/refs/tags/v1.3.1.src.tar.gz"
	ttysolitaire_cmd_src := exec.Command("curl", "-L", ttysolitaire_src_url, "-o", "source.tar.gz")
	err = ttysolitaire_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttysolitaire_cmd_direct := exec.Command("./binary")
	err = ttysolitaire_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
