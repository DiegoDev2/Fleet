package main

// ElixirLs - Language Server and Debugger for Elixir
// Homepage: https://elixir-lsp.github.io/elixir-ls

import (
	"fmt"
	
	"os/exec"
)

func installElixirLs() {
	// Método 1: Descargar y extraer .tar.gz
	elixirls_tar_url := "https://github.com/elixir-lsp/elixir-ls/archive/refs/tags/v0.23.0.tar.gz"
	elixirls_cmd_tar := exec.Command("curl", "-L", elixirls_tar_url, "-o", "package.tar.gz")
	err := elixirls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elixirls_zip_url := "https://github.com/elixir-lsp/elixir-ls/archive/refs/tags/v0.23.0.zip"
	elixirls_cmd_zip := exec.Command("curl", "-L", elixirls_zip_url, "-o", "package.zip")
	err = elixirls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elixirls_bin_url := "https://github.com/elixir-lsp/elixir-ls/archive/refs/tags/v0.23.0.bin"
	elixirls_cmd_bin := exec.Command("curl", "-L", elixirls_bin_url, "-o", "binary.bin")
	err = elixirls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elixirls_src_url := "https://github.com/elixir-lsp/elixir-ls/archive/refs/tags/v0.23.0.src.tar.gz"
	elixirls_cmd_src := exec.Command("curl", "-L", elixirls_src_url, "-o", "source.tar.gz")
	err = elixirls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elixirls_cmd_direct := exec.Command("./binary")
	err = elixirls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: elixir")
	exec.Command("latte", "install", "elixir").Run()
}
