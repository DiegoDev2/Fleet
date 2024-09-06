package main

// Elvis - Erlang Style Reviewer
// Homepage: https://github.com/inaka/elvis

import (
	"fmt"
	
	"os/exec"
)

func installElvis() {
	// Método 1: Descargar y extraer .tar.gz
	elvis_tar_url := "https://github.com/inaka/elvis/archive/refs/tags/3.2.5.tar.gz"
	elvis_cmd_tar := exec.Command("curl", "-L", elvis_tar_url, "-o", "package.tar.gz")
	err := elvis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elvis_zip_url := "https://github.com/inaka/elvis/archive/refs/tags/3.2.5.zip"
	elvis_cmd_zip := exec.Command("curl", "-L", elvis_zip_url, "-o", "package.zip")
	err = elvis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elvis_bin_url := "https://github.com/inaka/elvis/archive/refs/tags/3.2.5.bin"
	elvis_cmd_bin := exec.Command("curl", "-L", elvis_bin_url, "-o", "binary.bin")
	err = elvis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elvis_src_url := "https://github.com/inaka/elvis/archive/refs/tags/3.2.5.src.tar.gz"
	elvis_cmd_src := exec.Command("curl", "-L", elvis_src_url, "-o", "source.tar.gz")
	err = elvis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elvis_cmd_direct := exec.Command("./binary")
	err = elvis_cmd_direct.Run()
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
