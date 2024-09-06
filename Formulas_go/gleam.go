package main

// Gleam - Statically typed language for the Erlang VM
// Homepage: https://gleam.run

import (
	"fmt"
	
	"os/exec"
)

func installGleam() {
	// Método 1: Descargar y extraer .tar.gz
	gleam_tar_url := "https://github.com/gleam-lang/gleam/archive/refs/tags/v1.4.1.tar.gz"
	gleam_cmd_tar := exec.Command("curl", "-L", gleam_tar_url, "-o", "package.tar.gz")
	err := gleam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gleam_zip_url := "https://github.com/gleam-lang/gleam/archive/refs/tags/v1.4.1.zip"
	gleam_cmd_zip := exec.Command("curl", "-L", gleam_zip_url, "-o", "package.zip")
	err = gleam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gleam_bin_url := "https://github.com/gleam-lang/gleam/archive/refs/tags/v1.4.1.bin"
	gleam_cmd_bin := exec.Command("curl", "-L", gleam_bin_url, "-o", "binary.bin")
	err = gleam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gleam_src_url := "https://github.com/gleam-lang/gleam/archive/refs/tags/v1.4.1.src.tar.gz"
	gleam_cmd_src := exec.Command("curl", "-L", gleam_src_url, "-o", "source.tar.gz")
	err = gleam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gleam_cmd_direct := exec.Command("./binary")
	err = gleam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: erlang")
exec.Command("latte", "install", "erlang")
	fmt.Println("Instalando dependencia: rebar3")
exec.Command("latte", "install", "rebar3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
