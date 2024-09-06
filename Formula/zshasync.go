package main

// ZshAsync - Perform tasks asynchronously without external tools
// Homepage: https://github.com/mafredri/zsh-async

import (
	"fmt"
	
	"os/exec"
)

func installZshAsync() {
	// Método 1: Descargar y extraer .tar.gz
	zshasync_tar_url := "https://github.com/mafredri/zsh-async/archive/refs/tags/v1.8.6.tar.gz"
	zshasync_cmd_tar := exec.Command("curl", "-L", zshasync_tar_url, "-o", "package.tar.gz")
	err := zshasync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshasync_zip_url := "https://github.com/mafredri/zsh-async/archive/refs/tags/v1.8.6.zip"
	zshasync_cmd_zip := exec.Command("curl", "-L", zshasync_zip_url, "-o", "package.zip")
	err = zshasync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshasync_bin_url := "https://github.com/mafredri/zsh-async/archive/refs/tags/v1.8.6.bin"
	zshasync_cmd_bin := exec.Command("curl", "-L", zshasync_bin_url, "-o", "binary.bin")
	err = zshasync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshasync_src_url := "https://github.com/mafredri/zsh-async/archive/refs/tags/v1.8.6.src.tar.gz"
	zshasync_cmd_src := exec.Command("curl", "-L", zshasync_src_url, "-o", "source.tar.gz")
	err = zshasync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshasync_cmd_direct := exec.Command("./binary")
	err = zshasync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
