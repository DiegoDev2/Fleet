package main

// Z80dasm - Disassembler for the Zilog Z80 microprocessor and compatibles
// Homepage: https://www.tablix.org/~avian/blog/articles/z80dasm/

import (
	"fmt"
	
	"os/exec"
)

func installZ80dasm() {
	// Método 1: Descargar y extraer .tar.gz
	z80dasm_tar_url := "https://www.tablix.org/~avian/z80dasm/z80dasm-1.2.0.tar.gz"
	z80dasm_cmd_tar := exec.Command("curl", "-L", z80dasm_tar_url, "-o", "package.tar.gz")
	err := z80dasm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	z80dasm_zip_url := "https://www.tablix.org/~avian/z80dasm/z80dasm-1.2.0.zip"
	z80dasm_cmd_zip := exec.Command("curl", "-L", z80dasm_zip_url, "-o", "package.zip")
	err = z80dasm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	z80dasm_bin_url := "https://www.tablix.org/~avian/z80dasm/z80dasm-1.2.0.bin"
	z80dasm_cmd_bin := exec.Command("curl", "-L", z80dasm_bin_url, "-o", "binary.bin")
	err = z80dasm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	z80dasm_src_url := "https://www.tablix.org/~avian/z80dasm/z80dasm-1.2.0.src.tar.gz"
	z80dasm_cmd_src := exec.Command("curl", "-L", z80dasm_src_url, "-o", "source.tar.gz")
	err = z80dasm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	z80dasm_cmd_direct := exec.Command("./binary")
	err = z80dasm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
