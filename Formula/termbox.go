package main

// Termbox - Library for writing text-based user interfaces
// Homepage: https://github.com/termbox/termbox

import (
	"fmt"
	
	"os/exec"
)

func installTermbox() {
	// Método 1: Descargar y extraer .tar.gz
	termbox_tar_url := "https://github.com/termbox/termbox/archive/refs/tags/v1.1.4.tar.gz"
	termbox_cmd_tar := exec.Command("curl", "-L", termbox_tar_url, "-o", "package.tar.gz")
	err := termbox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	termbox_zip_url := "https://github.com/termbox/termbox/archive/refs/tags/v1.1.4.zip"
	termbox_cmd_zip := exec.Command("curl", "-L", termbox_zip_url, "-o", "package.zip")
	err = termbox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	termbox_bin_url := "https://github.com/termbox/termbox/archive/refs/tags/v1.1.4.bin"
	termbox_cmd_bin := exec.Command("curl", "-L", termbox_bin_url, "-o", "binary.bin")
	err = termbox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	termbox_src_url := "https://github.com/termbox/termbox/archive/refs/tags/v1.1.4.src.tar.gz"
	termbox_cmd_src := exec.Command("curl", "-L", termbox_src_url, "-o", "source.tar.gz")
	err = termbox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	termbox_cmd_direct := exec.Command("./binary")
	err = termbox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
