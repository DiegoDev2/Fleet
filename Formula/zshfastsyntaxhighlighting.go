package main

// ZshFastSyntaxHighlighting - Feature-rich syntax highlighting for Zsh
// Homepage: https://github.com/zdharma-continuum/fast-syntax-highlighting

import (
	"fmt"
	
	"os/exec"
)

func installZshFastSyntaxHighlighting() {
	// Método 1: Descargar y extraer .tar.gz
	zshfastsyntaxhighlighting_tar_url := "https://github.com/zdharma-continuum/fast-syntax-highlighting/archive/refs/tags/v1.55.tar.gz"
	zshfastsyntaxhighlighting_cmd_tar := exec.Command("curl", "-L", zshfastsyntaxhighlighting_tar_url, "-o", "package.tar.gz")
	err := zshfastsyntaxhighlighting_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshfastsyntaxhighlighting_zip_url := "https://github.com/zdharma-continuum/fast-syntax-highlighting/archive/refs/tags/v1.55.zip"
	zshfastsyntaxhighlighting_cmd_zip := exec.Command("curl", "-L", zshfastsyntaxhighlighting_zip_url, "-o", "package.zip")
	err = zshfastsyntaxhighlighting_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshfastsyntaxhighlighting_bin_url := "https://github.com/zdharma-continuum/fast-syntax-highlighting/archive/refs/tags/v1.55.bin"
	zshfastsyntaxhighlighting_cmd_bin := exec.Command("curl", "-L", zshfastsyntaxhighlighting_bin_url, "-o", "binary.bin")
	err = zshfastsyntaxhighlighting_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshfastsyntaxhighlighting_src_url := "https://github.com/zdharma-continuum/fast-syntax-highlighting/archive/refs/tags/v1.55.src.tar.gz"
	zshfastsyntaxhighlighting_cmd_src := exec.Command("curl", "-L", zshfastsyntaxhighlighting_src_url, "-o", "source.tar.gz")
	err = zshfastsyntaxhighlighting_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshfastsyntaxhighlighting_cmd_direct := exec.Command("./binary")
	err = zshfastsyntaxhighlighting_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
