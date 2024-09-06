package main

// Brainfuck - Interpreter for the brainfuck language
// Homepage: https://github.com/fabianishere/brainfuck

import (
	"fmt"
	
	"os/exec"
)

func installBrainfuck() {
	// Método 1: Descargar y extraer .tar.gz
	brainfuck_tar_url := "https://github.com/fabianishere/brainfuck/archive/refs/tags/2.7.3.tar.gz"
	brainfuck_cmd_tar := exec.Command("curl", "-L", brainfuck_tar_url, "-o", "package.tar.gz")
	err := brainfuck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brainfuck_zip_url := "https://github.com/fabianishere/brainfuck/archive/refs/tags/2.7.3.zip"
	brainfuck_cmd_zip := exec.Command("curl", "-L", brainfuck_zip_url, "-o", "package.zip")
	err = brainfuck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brainfuck_bin_url := "https://github.com/fabianishere/brainfuck/archive/refs/tags/2.7.3.bin"
	brainfuck_cmd_bin := exec.Command("curl", "-L", brainfuck_bin_url, "-o", "binary.bin")
	err = brainfuck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brainfuck_src_url := "https://github.com/fabianishere/brainfuck/archive/refs/tags/2.7.3.src.tar.gz"
	brainfuck_cmd_src := exec.Command("curl", "-L", brainfuck_src_url, "-o", "source.tar.gz")
	err = brainfuck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brainfuck_cmd_direct := exec.Command("./binary")
	err = brainfuck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
