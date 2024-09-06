package main

// ErlangLs - Erlang Language Server
// Homepage: https://erlang-ls.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installErlangLs() {
	// Método 1: Descargar y extraer .tar.gz
	erlangls_tar_url := "https://github.com/erlang-ls/erlang_ls/archive/refs/tags/0.52.0.tar.gz"
	erlangls_cmd_tar := exec.Command("curl", "-L", erlangls_tar_url, "-o", "package.tar.gz")
	err := erlangls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erlangls_zip_url := "https://github.com/erlang-ls/erlang_ls/archive/refs/tags/0.52.0.zip"
	erlangls_cmd_zip := exec.Command("curl", "-L", erlangls_zip_url, "-o", "package.zip")
	err = erlangls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erlangls_bin_url := "https://github.com/erlang-ls/erlang_ls/archive/refs/tags/0.52.0.bin"
	erlangls_cmd_bin := exec.Command("curl", "-L", erlangls_bin_url, "-o", "binary.bin")
	err = erlangls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erlangls_src_url := "https://github.com/erlang-ls/erlang_ls/archive/refs/tags/0.52.0.src.tar.gz"
	erlangls_cmd_src := exec.Command("curl", "-L", erlangls_src_url, "-o", "source.tar.gz")
	err = erlangls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erlangls_cmd_direct := exec.Command("./binary")
	err = erlangls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: erlang")
	exec.Command("latte", "install", "erlang").Run()
	fmt.Println("Instalando dependencia: rebar3")
	exec.Command("latte", "install", "rebar3").Run()
}
