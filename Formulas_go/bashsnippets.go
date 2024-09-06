package main

// BashSnippets - Collection of small bash scripts for heavy terminal users
// Homepage: https://github.com/alexanderepstein/Bash-Snippets

import (
	"fmt"
	
	"os/exec"
)

func installBashSnippets() {
	// Método 1: Descargar y extraer .tar.gz
	bashsnippets_tar_url := "https://github.com/alexanderepstein/Bash-Snippets/archive/refs/tags/v1.23.0.tar.gz"
	bashsnippets_cmd_tar := exec.Command("curl", "-L", bashsnippets_tar_url, "-o", "package.tar.gz")
	err := bashsnippets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashsnippets_zip_url := "https://github.com/alexanderepstein/Bash-Snippets/archive/refs/tags/v1.23.0.zip"
	bashsnippets_cmd_zip := exec.Command("curl", "-L", bashsnippets_zip_url, "-o", "package.zip")
	err = bashsnippets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashsnippets_bin_url := "https://github.com/alexanderepstein/Bash-Snippets/archive/refs/tags/v1.23.0.bin"
	bashsnippets_cmd_bin := exec.Command("curl", "-L", bashsnippets_bin_url, "-o", "binary.bin")
	err = bashsnippets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashsnippets_src_url := "https://github.com/alexanderepstein/Bash-Snippets/archive/refs/tags/v1.23.0.src.tar.gz"
	bashsnippets_cmd_src := exec.Command("curl", "-L", bashsnippets_src_url, "-o", "source.tar.gz")
	err = bashsnippets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashsnippets_cmd_direct := exec.Command("./binary")
	err = bashsnippets_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
