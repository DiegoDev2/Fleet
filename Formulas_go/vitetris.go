package main

// Vitetris - Terminal-based Tetris clone
// Homepage: https://www.victornils.net/tetris/

import (
	"fmt"
	
	"os/exec"
)

func installVitetris() {
	// Método 1: Descargar y extraer .tar.gz
	vitetris_tar_url := "https://github.com/vicgeralds/vitetris/archive/refs/tags/v0.59.1.tar.gz"
	vitetris_cmd_tar := exec.Command("curl", "-L", vitetris_tar_url, "-o", "package.tar.gz")
	err := vitetris_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vitetris_zip_url := "https://github.com/vicgeralds/vitetris/archive/refs/tags/v0.59.1.zip"
	vitetris_cmd_zip := exec.Command("curl", "-L", vitetris_zip_url, "-o", "package.zip")
	err = vitetris_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vitetris_bin_url := "https://github.com/vicgeralds/vitetris/archive/refs/tags/v0.59.1.bin"
	vitetris_cmd_bin := exec.Command("curl", "-L", vitetris_bin_url, "-o", "binary.bin")
	err = vitetris_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vitetris_src_url := "https://github.com/vicgeralds/vitetris/archive/refs/tags/v0.59.1.src.tar.gz"
	vitetris_cmd_src := exec.Command("curl", "-L", vitetris_src_url, "-o", "source.tar.gz")
	err = vitetris_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vitetris_cmd_direct := exec.Command("./binary")
	err = vitetris_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
