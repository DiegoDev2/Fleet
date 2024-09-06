package main

// GitOpen - Open GitHub webpages from a terminal
// Homepage: https://github.com/jeffreyiacono/git-open

import (
	"fmt"
	
	"os/exec"
)

func installGitOpen() {
	// Método 1: Descargar y extraer .tar.gz
	gitopen_tar_url := "https://github.com/jeffreyiacono/git-open/archive/refs/tags/v1.3.tar.gz"
	gitopen_cmd_tar := exec.Command("curl", "-L", gitopen_tar_url, "-o", "package.tar.gz")
	err := gitopen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitopen_zip_url := "https://github.com/jeffreyiacono/git-open/archive/refs/tags/v1.3.zip"
	gitopen_cmd_zip := exec.Command("curl", "-L", gitopen_zip_url, "-o", "package.zip")
	err = gitopen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitopen_bin_url := "https://github.com/jeffreyiacono/git-open/archive/refs/tags/v1.3.bin"
	gitopen_cmd_bin := exec.Command("curl", "-L", gitopen_bin_url, "-o", "binary.bin")
	err = gitopen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitopen_src_url := "https://github.com/jeffreyiacono/git-open/archive/refs/tags/v1.3.src.tar.gz"
	gitopen_cmd_src := exec.Command("curl", "-L", gitopen_src_url, "-o", "source.tar.gz")
	err = gitopen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitopen_cmd_direct := exec.Command("./binary")
	err = gitopen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
