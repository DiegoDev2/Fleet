package main

// Historian - Command-line utility for managing shell history in a SQLite database
// Homepage: https://github.com/jcsalterego/historian

import (
	"fmt"
	
	"os/exec"
)

func installHistorian() {
	// Método 1: Descargar y extraer .tar.gz
	historian_tar_url := "https://github.com/jcsalterego/historian/archive/refs/tags/0.0.2.tar.gz"
	historian_cmd_tar := exec.Command("curl", "-L", historian_tar_url, "-o", "package.tar.gz")
	err := historian_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	historian_zip_url := "https://github.com/jcsalterego/historian/archive/refs/tags/0.0.2.zip"
	historian_cmd_zip := exec.Command("curl", "-L", historian_zip_url, "-o", "package.zip")
	err = historian_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	historian_bin_url := "https://github.com/jcsalterego/historian/archive/refs/tags/0.0.2.bin"
	historian_cmd_bin := exec.Command("curl", "-L", historian_bin_url, "-o", "binary.bin")
	err = historian_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	historian_src_url := "https://github.com/jcsalterego/historian/archive/refs/tags/0.0.2.src.tar.gz"
	historian_cmd_src := exec.Command("curl", "-L", historian_src_url, "-o", "source.tar.gz")
	err = historian_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	historian_cmd_direct := exec.Command("./binary")
	err = historian_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
