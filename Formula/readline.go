package main

// Readline - Library for command-line editing
// Homepage: https://tiswww.case.edu/php/chet/readline/rltop.html

import (
	"fmt"
	
	"os/exec"
)

func installReadline() {
	// Método 1: Descargar y extraer .tar.gz
	readline_tar_url := "https://ftp.gnu.org/gnu/readline/readline-8.2.tar.gz"
	readline_cmd_tar := exec.Command("curl", "-L", readline_tar_url, "-o", "package.tar.gz")
	err := readline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	readline_zip_url := "https://ftp.gnu.org/gnu/readline/readline-8.2.zip"
	readline_cmd_zip := exec.Command("curl", "-L", readline_zip_url, "-o", "package.zip")
	err = readline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	readline_bin_url := "https://ftp.gnu.org/gnu/readline/readline-8.2.bin"
	readline_cmd_bin := exec.Command("curl", "-L", readline_bin_url, "-o", "binary.bin")
	err = readline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	readline_src_url := "https://ftp.gnu.org/gnu/readline/readline-8.2.src.tar.gz"
	readline_cmd_src := exec.Command("curl", "-L", readline_src_url, "-o", "source.tar.gz")
	err = readline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	readline_cmd_direct := exec.Command("./binary")
	err = readline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
