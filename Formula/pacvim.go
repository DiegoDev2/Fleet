package main

// Pacvim - Learn vim commands via a game
// Homepage: https://github.com/jmoon018/PacVim

import (
	"fmt"
	
	"os/exec"
)

func installPacvim() {
	// Método 1: Descargar y extraer .tar.gz
	pacvim_tar_url := "https://github.com/jmoon018/PacVim/archive/refs/tags/v1.1.1.tar.gz"
	pacvim_cmd_tar := exec.Command("curl", "-L", pacvim_tar_url, "-o", "package.tar.gz")
	err := pacvim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pacvim_zip_url := "https://github.com/jmoon018/PacVim/archive/refs/tags/v1.1.1.zip"
	pacvim_cmd_zip := exec.Command("curl", "-L", pacvim_zip_url, "-o", "package.zip")
	err = pacvim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pacvim_bin_url := "https://github.com/jmoon018/PacVim/archive/refs/tags/v1.1.1.bin"
	pacvim_cmd_bin := exec.Command("curl", "-L", pacvim_bin_url, "-o", "binary.bin")
	err = pacvim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pacvim_src_url := "https://github.com/jmoon018/PacVim/archive/refs/tags/v1.1.1.src.tar.gz"
	pacvim_cmd_src := exec.Command("curl", "-L", pacvim_src_url, "-o", "source.tar.gz")
	err = pacvim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pacvim_cmd_direct := exec.Command("./binary")
	err = pacvim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
