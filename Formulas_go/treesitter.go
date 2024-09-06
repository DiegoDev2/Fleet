package main

// TreeSitter - Parser generator tool and incremental parsing library
// Homepage: https://tree-sitter.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installTreeSitter() {
	// Método 1: Descargar y extraer .tar.gz
	treesitter_tar_url := "https://github.com/tree-sitter/tree-sitter/archive/refs/tags/v0.23.0.tar.gz"
	treesitter_cmd_tar := exec.Command("curl", "-L", treesitter_tar_url, "-o", "package.tar.gz")
	err := treesitter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	treesitter_zip_url := "https://github.com/tree-sitter/tree-sitter/archive/refs/tags/v0.23.0.zip"
	treesitter_cmd_zip := exec.Command("curl", "-L", treesitter_zip_url, "-o", "package.zip")
	err = treesitter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	treesitter_bin_url := "https://github.com/tree-sitter/tree-sitter/archive/refs/tags/v0.23.0.bin"
	treesitter_cmd_bin := exec.Command("curl", "-L", treesitter_bin_url, "-o", "binary.bin")
	err = treesitter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	treesitter_src_url := "https://github.com/tree-sitter/tree-sitter/archive/refs/tags/v0.23.0.src.tar.gz"
	treesitter_cmd_src := exec.Command("curl", "-L", treesitter_src_url, "-o", "source.tar.gz")
	err = treesitter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	treesitter_cmd_direct := exec.Command("./binary")
	err = treesitter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
