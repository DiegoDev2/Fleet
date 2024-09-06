package main

// Elixir - Functional metaprogramming aware language built on Erlang VM
// Homepage: https://elixir-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installElixir() {
	// Método 1: Descargar y extraer .tar.gz
	elixir_tar_url := "https://github.com/elixir-lang/elixir/archive/refs/tags/v1.17.2.tar.gz"
	elixir_cmd_tar := exec.Command("curl", "-L", elixir_tar_url, "-o", "package.tar.gz")
	err := elixir_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elixir_zip_url := "https://github.com/elixir-lang/elixir/archive/refs/tags/v1.17.2.zip"
	elixir_cmd_zip := exec.Command("curl", "-L", elixir_zip_url, "-o", "package.zip")
	err = elixir_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elixir_bin_url := "https://github.com/elixir-lang/elixir/archive/refs/tags/v1.17.2.bin"
	elixir_cmd_bin := exec.Command("curl", "-L", elixir_bin_url, "-o", "binary.bin")
	err = elixir_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elixir_src_url := "https://github.com/elixir-lang/elixir/archive/refs/tags/v1.17.2.src.tar.gz"
	elixir_cmd_src := exec.Command("curl", "-L", elixir_src_url, "-o", "source.tar.gz")
	err = elixir_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elixir_cmd_direct := exec.Command("./binary")
	err = elixir_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: erlang")
exec.Command("latte", "install", "erlang")
}
