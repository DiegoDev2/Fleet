package main

// Fswatch - Monitor a directory for changes and run a shell command
// Homepage: https://github.com/emcrisostomo/fswatch

import (
	"fmt"
	
	"os/exec"
)

func installFswatch() {
	// Método 1: Descargar y extraer .tar.gz
	fswatch_tar_url := "https://github.com/emcrisostomo/fswatch/releases/download/1.17.1/fswatch-1.17.1.tar.gz"
	fswatch_cmd_tar := exec.Command("curl", "-L", fswatch_tar_url, "-o", "package.tar.gz")
	err := fswatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fswatch_zip_url := "https://github.com/emcrisostomo/fswatch/releases/download/1.17.1/fswatch-1.17.1.zip"
	fswatch_cmd_zip := exec.Command("curl", "-L", fswatch_zip_url, "-o", "package.zip")
	err = fswatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fswatch_bin_url := "https://github.com/emcrisostomo/fswatch/releases/download/1.17.1/fswatch-1.17.1.bin"
	fswatch_cmd_bin := exec.Command("curl", "-L", fswatch_bin_url, "-o", "binary.bin")
	err = fswatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fswatch_src_url := "https://github.com/emcrisostomo/fswatch/releases/download/1.17.1/fswatch-1.17.1.src.tar.gz"
	fswatch_cmd_src := exec.Command("curl", "-L", fswatch_src_url, "-o", "source.tar.gz")
	err = fswatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fswatch_cmd_direct := exec.Command("./binary")
	err = fswatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
