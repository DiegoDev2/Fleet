package main

// Nsnake - Classic snake game with textual interface
// Homepage: https://github.com/alexdantas/nSnake

import (
	"fmt"
	
	"os/exec"
)

func installNsnake() {
	// Método 1: Descargar y extraer .tar.gz
	nsnake_tar_url := "https://downloads.sourceforge.net/project/nsnake/GNU-Linux/nsnake-3.0.1.tar.gz"
	nsnake_cmd_tar := exec.Command("curl", "-L", nsnake_tar_url, "-o", "package.tar.gz")
	err := nsnake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nsnake_zip_url := "https://downloads.sourceforge.net/project/nsnake/GNU-Linux/nsnake-3.0.1.zip"
	nsnake_cmd_zip := exec.Command("curl", "-L", nsnake_zip_url, "-o", "package.zip")
	err = nsnake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nsnake_bin_url := "https://downloads.sourceforge.net/project/nsnake/GNU-Linux/nsnake-3.0.1.bin"
	nsnake_cmd_bin := exec.Command("curl", "-L", nsnake_bin_url, "-o", "binary.bin")
	err = nsnake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nsnake_src_url := "https://downloads.sourceforge.net/project/nsnake/GNU-Linux/nsnake-3.0.1.src.tar.gz"
	nsnake_cmd_src := exec.Command("curl", "-L", nsnake_src_url, "-o", "source.tar.gz")
	err = nsnake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nsnake_cmd_direct := exec.Command("./binary")
	err = nsnake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
