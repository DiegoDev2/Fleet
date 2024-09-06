package main

// Rebar3 - Erlang build tool
// Homepage: https://github.com/erlang/rebar3

import (
	"fmt"
	
	"os/exec"
)

func installRebar3() {
	// Método 1: Descargar y extraer .tar.gz
	rebar3_tar_url := "https://github.com/erlang/rebar3/archive/refs/tags/3.24.0.tar.gz"
	rebar3_cmd_tar := exec.Command("curl", "-L", rebar3_tar_url, "-o", "package.tar.gz")
	err := rebar3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rebar3_zip_url := "https://github.com/erlang/rebar3/archive/refs/tags/3.24.0.zip"
	rebar3_cmd_zip := exec.Command("curl", "-L", rebar3_zip_url, "-o", "package.zip")
	err = rebar3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rebar3_bin_url := "https://github.com/erlang/rebar3/archive/refs/tags/3.24.0.bin"
	rebar3_cmd_bin := exec.Command("curl", "-L", rebar3_bin_url, "-o", "binary.bin")
	err = rebar3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rebar3_src_url := "https://github.com/erlang/rebar3/archive/refs/tags/3.24.0.src.tar.gz"
	rebar3_cmd_src := exec.Command("curl", "-L", rebar3_src_url, "-o", "source.tar.gz")
	err = rebar3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rebar3_cmd_direct := exec.Command("./binary")
	err = rebar3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: erlang")
	exec.Command("latte", "install", "erlang").Run()
}
