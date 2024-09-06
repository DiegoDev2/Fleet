package main

// Proper - QuickCheck-inspired property-based testing tool for Erlang
// Homepage: https://proper-testing.github.io

import (
	"fmt"
	
	"os/exec"
)

func installProper() {
	// Método 1: Descargar y extraer .tar.gz
	proper_tar_url := "https://github.com/proper-testing/proper/archive/refs/tags/v1.4.tar.gz"
	proper_cmd_tar := exec.Command("curl", "-L", proper_tar_url, "-o", "package.tar.gz")
	err := proper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proper_zip_url := "https://github.com/proper-testing/proper/archive/refs/tags/v1.4.zip"
	proper_cmd_zip := exec.Command("curl", "-L", proper_zip_url, "-o", "package.zip")
	err = proper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proper_bin_url := "https://github.com/proper-testing/proper/archive/refs/tags/v1.4.bin"
	proper_cmd_bin := exec.Command("curl", "-L", proper_bin_url, "-o", "binary.bin")
	err = proper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proper_src_url := "https://github.com/proper-testing/proper/archive/refs/tags/v1.4.src.tar.gz"
	proper_cmd_src := exec.Command("curl", "-L", proper_src_url, "-o", "source.tar.gz")
	err = proper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proper_cmd_direct := exec.Command("./binary")
	err = proper_cmd_direct.Run()
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
