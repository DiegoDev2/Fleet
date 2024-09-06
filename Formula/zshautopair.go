package main

// ZshAutopair - Auto-close and delete matching delimiters in zsh
// Homepage: https://github.com/hlissner/zsh-autopair

import (
	"fmt"
	
	"os/exec"
)

func installZshAutopair() {
	// Método 1: Descargar y extraer .tar.gz
	zshautopair_tar_url := "https://github.com/hlissner/zsh-autopair/archive/refs/tags/v1.0.tar.gz"
	zshautopair_cmd_tar := exec.Command("curl", "-L", zshautopair_tar_url, "-o", "package.tar.gz")
	err := zshautopair_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshautopair_zip_url := "https://github.com/hlissner/zsh-autopair/archive/refs/tags/v1.0.zip"
	zshautopair_cmd_zip := exec.Command("curl", "-L", zshautopair_zip_url, "-o", "package.zip")
	err = zshautopair_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshautopair_bin_url := "https://github.com/hlissner/zsh-autopair/archive/refs/tags/v1.0.bin"
	zshautopair_cmd_bin := exec.Command("curl", "-L", zshautopair_bin_url, "-o", "binary.bin")
	err = zshautopair_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshautopair_src_url := "https://github.com/hlissner/zsh-autopair/archive/refs/tags/v1.0.src.tar.gz"
	zshautopair_cmd_src := exec.Command("curl", "-L", zshautopair_src_url, "-o", "source.tar.gz")
	err = zshautopair_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshautopair_cmd_direct := exec.Command("./binary")
	err = zshautopair_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
