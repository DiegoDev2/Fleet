package main

// TodoTxt - Minimal, todo.txt-focused editor
// Homepage: http://todotxt.org/

import (
	"fmt"
	
	"os/exec"
)

func installTodoTxt() {
	// Método 1: Descargar y extraer .tar.gz
	todotxt_tar_url := "https://github.com/todotxt/todo.txt-cli/releases/download/v2.12.0/todo.txt_cli-2.12.0.tar.gz"
	todotxt_cmd_tar := exec.Command("curl", "-L", todotxt_tar_url, "-o", "package.tar.gz")
	err := todotxt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	todotxt_zip_url := "https://github.com/todotxt/todo.txt-cli/releases/download/v2.12.0/todo.txt_cli-2.12.0.zip"
	todotxt_cmd_zip := exec.Command("curl", "-L", todotxt_zip_url, "-o", "package.zip")
	err = todotxt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	todotxt_bin_url := "https://github.com/todotxt/todo.txt-cli/releases/download/v2.12.0/todo.txt_cli-2.12.0.bin"
	todotxt_cmd_bin := exec.Command("curl", "-L", todotxt_bin_url, "-o", "binary.bin")
	err = todotxt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	todotxt_src_url := "https://github.com/todotxt/todo.txt-cli/releases/download/v2.12.0/todo.txt_cli-2.12.0.src.tar.gz"
	todotxt_cmd_src := exec.Command("curl", "-L", todotxt_src_url, "-o", "source.tar.gz")
	err = todotxt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	todotxt_cmd_direct := exec.Command("./binary")
	err = todotxt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
