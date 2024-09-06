package main

// Term - Open terminal in specified directory (and optionally run command)
// Homepage: https://github.com/liyanage/macosx-shell-scripts/blob/HEAD/term

import (
	"fmt"
	
	"os/exec"
)

func installTerm() {
	// Método 1: Descargar y extraer .tar.gz
	term_tar_url := "https://raw.githubusercontent.com/liyanage/macosx-shell-scripts/e29f7eaa1eb13d78056dec85dc517626ab1d93e3/term"
	term_cmd_tar := exec.Command("curl", "-L", term_tar_url, "-o", "package.tar.gz")
	err := term_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	term_zip_url := "https://raw.githubusercontent.com/liyanage/macosx-shell-scripts/e29f7eaa1eb13d78056dec85dc517626ab1d93e3/term"
	term_cmd_zip := exec.Command("curl", "-L", term_zip_url, "-o", "package.zip")
	err = term_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	term_bin_url := "https://raw.githubusercontent.com/liyanage/macosx-shell-scripts/e29f7eaa1eb13d78056dec85dc517626ab1d93e3/term"
	term_cmd_bin := exec.Command("curl", "-L", term_bin_url, "-o", "binary.bin")
	err = term_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	term_src_url := "https://raw.githubusercontent.com/liyanage/macosx-shell-scripts/e29f7eaa1eb13d78056dec85dc517626ab1d93e3/term"
	term_cmd_src := exec.Command("curl", "-L", term_src_url, "-o", "source.tar.gz")
	err = term_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	term_cmd_direct := exec.Command("./binary")
	err = term_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
