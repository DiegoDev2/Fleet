package main

// Z80asm - Assembler for the Zilog Z80 microprcessor and compatibles
// Homepage: https://www.nongnu.org/z80asm/

import (
	"fmt"
	
	"os/exec"
)

func installZ80asm() {
	// Método 1: Descargar y extraer .tar.gz
	z80asm_tar_url := "https://download.savannah.gnu.org/releases/z80asm/z80asm-1.8.tar.gz"
	z80asm_cmd_tar := exec.Command("curl", "-L", z80asm_tar_url, "-o", "package.tar.gz")
	err := z80asm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	z80asm_zip_url := "https://download.savannah.gnu.org/releases/z80asm/z80asm-1.8.zip"
	z80asm_cmd_zip := exec.Command("curl", "-L", z80asm_zip_url, "-o", "package.zip")
	err = z80asm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	z80asm_bin_url := "https://download.savannah.gnu.org/releases/z80asm/z80asm-1.8.bin"
	z80asm_cmd_bin := exec.Command("curl", "-L", z80asm_bin_url, "-o", "binary.bin")
	err = z80asm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	z80asm_src_url := "https://download.savannah.gnu.org/releases/z80asm/z80asm-1.8.src.tar.gz"
	z80asm_cmd_src := exec.Command("curl", "-L", z80asm_src_url, "-o", "source.tar.gz")
	err = z80asm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	z80asm_cmd_direct := exec.Command("./binary")
	err = z80asm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
