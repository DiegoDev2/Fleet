package main

// Redir - Port redirector
// Homepage: https://github.com/TracyWebTech/redir

import (
	"fmt"
	
	"os/exec"
)

func installRedir() {
	// Método 1: Descargar y extraer .tar.gz
	redir_tar_url := "https://github.com/TracyWebTech/redir/archive/refs/tags/2.2.1-9.tar.gz"
	redir_cmd_tar := exec.Command("curl", "-L", redir_tar_url, "-o", "package.tar.gz")
	err := redir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redir_zip_url := "https://github.com/TracyWebTech/redir/archive/refs/tags/2.2.1-9.zip"
	redir_cmd_zip := exec.Command("curl", "-L", redir_zip_url, "-o", "package.zip")
	err = redir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redir_bin_url := "https://github.com/TracyWebTech/redir/archive/refs/tags/2.2.1-9.bin"
	redir_cmd_bin := exec.Command("curl", "-L", redir_bin_url, "-o", "binary.bin")
	err = redir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redir_src_url := "https://github.com/TracyWebTech/redir/archive/refs/tags/2.2.1-9.src.tar.gz"
	redir_cmd_src := exec.Command("curl", "-L", redir_src_url, "-o", "source.tar.gz")
	err = redir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redir_cmd_direct := exec.Command("./binary")
	err = redir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
