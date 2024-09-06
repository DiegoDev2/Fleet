package main

// ZshSyntaxHighlighting - Fish shell like syntax highlighting for zsh
// Homepage: https://github.com/zsh-users/zsh-syntax-highlighting

import (
	"fmt"
	
	"os/exec"
)

func installZshSyntaxHighlighting() {
	// Método 1: Descargar y extraer .tar.gz
	zshsyntaxhighlighting_tar_url := "https://github.com/zsh-users/zsh-syntax-highlighting/archive/refs/tags/0.8.0.tar.gz"
	zshsyntaxhighlighting_cmd_tar := exec.Command("curl", "-L", zshsyntaxhighlighting_tar_url, "-o", "package.tar.gz")
	err := zshsyntaxhighlighting_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshsyntaxhighlighting_zip_url := "https://github.com/zsh-users/zsh-syntax-highlighting/archive/refs/tags/0.8.0.zip"
	zshsyntaxhighlighting_cmd_zip := exec.Command("curl", "-L", zshsyntaxhighlighting_zip_url, "-o", "package.zip")
	err = zshsyntaxhighlighting_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshsyntaxhighlighting_bin_url := "https://github.com/zsh-users/zsh-syntax-highlighting/archive/refs/tags/0.8.0.bin"
	zshsyntaxhighlighting_cmd_bin := exec.Command("curl", "-L", zshsyntaxhighlighting_bin_url, "-o", "binary.bin")
	err = zshsyntaxhighlighting_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshsyntaxhighlighting_src_url := "https://github.com/zsh-users/zsh-syntax-highlighting/archive/refs/tags/0.8.0.src.tar.gz"
	zshsyntaxhighlighting_cmd_src := exec.Command("curl", "-L", zshsyntaxhighlighting_src_url, "-o", "source.tar.gz")
	err = zshsyntaxhighlighting_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshsyntaxhighlighting_cmd_direct := exec.Command("./binary")
	err = zshsyntaxhighlighting_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
