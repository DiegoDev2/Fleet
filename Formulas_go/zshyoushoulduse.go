package main

// ZshYouShouldUse - ZSH plugin that reminds you to use existing aliases for commands you just typed
// Homepage: https://github.com/MichaelAquilina/zsh-you-should-use

import (
	"fmt"
	
	"os/exec"
)

func installZshYouShouldUse() {
	// Método 1: Descargar y extraer .tar.gz
	zshyoushoulduse_tar_url := "https://github.com/MichaelAquilina/zsh-you-should-use/archive/refs/tags/1.9.0.tar.gz"
	zshyoushoulduse_cmd_tar := exec.Command("curl", "-L", zshyoushoulduse_tar_url, "-o", "package.tar.gz")
	err := zshyoushoulduse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshyoushoulduse_zip_url := "https://github.com/MichaelAquilina/zsh-you-should-use/archive/refs/tags/1.9.0.zip"
	zshyoushoulduse_cmd_zip := exec.Command("curl", "-L", zshyoushoulduse_zip_url, "-o", "package.zip")
	err = zshyoushoulduse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshyoushoulduse_bin_url := "https://github.com/MichaelAquilina/zsh-you-should-use/archive/refs/tags/1.9.0.bin"
	zshyoushoulduse_cmd_bin := exec.Command("curl", "-L", zshyoushoulduse_bin_url, "-o", "binary.bin")
	err = zshyoushoulduse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshyoushoulduse_src_url := "https://github.com/MichaelAquilina/zsh-you-should-use/archive/refs/tags/1.9.0.src.tar.gz"
	zshyoushoulduse_cmd_src := exec.Command("curl", "-L", zshyoushoulduse_src_url, "-o", "source.tar.gz")
	err = zshyoushoulduse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshyoushoulduse_cmd_direct := exec.Command("./binary")
	err = zshyoushoulduse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
