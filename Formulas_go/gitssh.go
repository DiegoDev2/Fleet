package main

// GitSsh - Proxy for serving git repositories over SSH
// Homepage: https://github.com/lemarsu/git-ssh

import (
	"fmt"
	
	"os/exec"
)

func installGitSsh() {
	// Método 1: Descargar y extraer .tar.gz
	gitssh_tar_url := "https://github.com/lemarsu/git-ssh/archive/refs/tags/v0.2.0.tar.gz"
	gitssh_cmd_tar := exec.Command("curl", "-L", gitssh_tar_url, "-o", "package.tar.gz")
	err := gitssh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitssh_zip_url := "https://github.com/lemarsu/git-ssh/archive/refs/tags/v0.2.0.zip"
	gitssh_cmd_zip := exec.Command("curl", "-L", gitssh_zip_url, "-o", "package.zip")
	err = gitssh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitssh_bin_url := "https://github.com/lemarsu/git-ssh/archive/refs/tags/v0.2.0.bin"
	gitssh_cmd_bin := exec.Command("curl", "-L", gitssh_bin_url, "-o", "binary.bin")
	err = gitssh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitssh_src_url := "https://github.com/lemarsu/git-ssh/archive/refs/tags/v0.2.0.src.tar.gz"
	gitssh_cmd_src := exec.Command("curl", "-L", gitssh_src_url, "-o", "source.tar.gz")
	err = gitssh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitssh_cmd_direct := exec.Command("./binary")
	err = gitssh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
