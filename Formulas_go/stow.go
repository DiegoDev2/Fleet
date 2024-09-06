package main

// Stow - Organize software neatly under a single directory tree (e.g. /usr/local)
// Homepage: https://www.gnu.org/software/stow/

import (
	"fmt"
	
	"os/exec"
)

func installStow() {
	// Método 1: Descargar y extraer .tar.gz
	stow_tar_url := "https://ftp.gnu.org/gnu/stow/stow-2.4.0.tar.gz"
	stow_cmd_tar := exec.Command("curl", "-L", stow_tar_url, "-o", "package.tar.gz")
	err := stow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stow_zip_url := "https://ftp.gnu.org/gnu/stow/stow-2.4.0.zip"
	stow_cmd_zip := exec.Command("curl", "-L", stow_zip_url, "-o", "package.zip")
	err = stow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stow_bin_url := "https://ftp.gnu.org/gnu/stow/stow-2.4.0.bin"
	stow_cmd_bin := exec.Command("curl", "-L", stow_bin_url, "-o", "binary.bin")
	err = stow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stow_src_url := "https://ftp.gnu.org/gnu/stow/stow-2.4.0.src.tar.gz"
	stow_cmd_src := exec.Command("curl", "-L", stow_src_url, "-o", "source.tar.gz")
	err = stow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stow_cmd_direct := exec.Command("./binary")
	err = stow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
