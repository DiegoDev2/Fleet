package main

// Syntaxerl - Syntax checker for Erlang code and config files
// Homepage: https://github.com/ten0s/syntaxerl

import (
	"fmt"
	
	"os/exec"
)

func installSyntaxerl() {
	// Método 1: Descargar y extraer .tar.gz
	syntaxerl_tar_url := "https://github.com/ten0s/syntaxerl/archive/refs/tags/0.15.0.tar.gz"
	syntaxerl_cmd_tar := exec.Command("curl", "-L", syntaxerl_tar_url, "-o", "package.tar.gz")
	err := syntaxerl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	syntaxerl_zip_url := "https://github.com/ten0s/syntaxerl/archive/refs/tags/0.15.0.zip"
	syntaxerl_cmd_zip := exec.Command("curl", "-L", syntaxerl_zip_url, "-o", "package.zip")
	err = syntaxerl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	syntaxerl_bin_url := "https://github.com/ten0s/syntaxerl/archive/refs/tags/0.15.0.bin"
	syntaxerl_cmd_bin := exec.Command("curl", "-L", syntaxerl_bin_url, "-o", "binary.bin")
	err = syntaxerl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	syntaxerl_src_url := "https://github.com/ten0s/syntaxerl/archive/refs/tags/0.15.0.src.tar.gz"
	syntaxerl_cmd_src := exec.Command("curl", "-L", syntaxerl_src_url, "-o", "source.tar.gz")
	err = syntaxerl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	syntaxerl_cmd_direct := exec.Command("./binary")
	err = syntaxerl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rebar3")
	exec.Command("latte", "install", "rebar3").Run()
	fmt.Println("Instalando dependencia: erlang")
	exec.Command("latte", "install", "erlang").Run()
}
