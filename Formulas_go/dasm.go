package main

// Dasm - Macro assembler with support for several 8-bit microprocessors
// Homepage: https://dasm-assembler.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installDasm() {
	// Método 1: Descargar y extraer .tar.gz
	dasm_tar_url := "https://github.com/dasm-assembler/dasm/archive/refs/tags/2.20.14.1.tar.gz"
	dasm_cmd_tar := exec.Command("curl", "-L", dasm_tar_url, "-o", "package.tar.gz")
	err := dasm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dasm_zip_url := "https://github.com/dasm-assembler/dasm/archive/refs/tags/2.20.14.1.zip"
	dasm_cmd_zip := exec.Command("curl", "-L", dasm_zip_url, "-o", "package.zip")
	err = dasm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dasm_bin_url := "https://github.com/dasm-assembler/dasm/archive/refs/tags/2.20.14.1.bin"
	dasm_cmd_bin := exec.Command("curl", "-L", dasm_bin_url, "-o", "binary.bin")
	err = dasm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dasm_src_url := "https://github.com/dasm-assembler/dasm/archive/refs/tags/2.20.14.1.src.tar.gz"
	dasm_cmd_src := exec.Command("curl", "-L", dasm_src_url, "-o", "source.tar.gz")
	err = dasm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dasm_cmd_direct := exec.Command("./binary")
	err = dasm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
