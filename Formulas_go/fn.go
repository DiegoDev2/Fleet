package main

// Fn - Command-line tool for the fn project
// Homepage: https://fnproject.io

import (
	"fmt"
	
	"os/exec"
)

func installFn() {
	// Método 1: Descargar y extraer .tar.gz
	fn_tar_url := "https://github.com/fnproject/cli/archive/refs/tags/0.6.34.tar.gz"
	fn_cmd_tar := exec.Command("curl", "-L", fn_tar_url, "-o", "package.tar.gz")
	err := fn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fn_zip_url := "https://github.com/fnproject/cli/archive/refs/tags/0.6.34.zip"
	fn_cmd_zip := exec.Command("curl", "-L", fn_zip_url, "-o", "package.zip")
	err = fn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fn_bin_url := "https://github.com/fnproject/cli/archive/refs/tags/0.6.34.bin"
	fn_cmd_bin := exec.Command("curl", "-L", fn_bin_url, "-o", "binary.bin")
	err = fn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fn_src_url := "https://github.com/fnproject/cli/archive/refs/tags/0.6.34.src.tar.gz"
	fn_cmd_src := exec.Command("curl", "-L", fn_src_url, "-o", "source.tar.gz")
	err = fn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fn_cmd_direct := exec.Command("./binary")
	err = fn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
