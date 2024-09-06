package main

// Peg - Program to perform pattern matching on text
// Homepage: https://www.piumarta.com/software/peg/

import (
	"fmt"
	
	"os/exec"
)

func installPeg() {
	// Método 1: Descargar y extraer .tar.gz
	peg_tar_url := "https://www.piumarta.com/software/peg/peg-0.1.19.tar.gz"
	peg_cmd_tar := exec.Command("curl", "-L", peg_tar_url, "-o", "package.tar.gz")
	err := peg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	peg_zip_url := "https://www.piumarta.com/software/peg/peg-0.1.19.zip"
	peg_cmd_zip := exec.Command("curl", "-L", peg_zip_url, "-o", "package.zip")
	err = peg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	peg_bin_url := "https://www.piumarta.com/software/peg/peg-0.1.19.bin"
	peg_cmd_bin := exec.Command("curl", "-L", peg_bin_url, "-o", "binary.bin")
	err = peg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	peg_src_url := "https://www.piumarta.com/software/peg/peg-0.1.19.src.tar.gz"
	peg_cmd_src := exec.Command("curl", "-L", peg_src_url, "-o", "source.tar.gz")
	err = peg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	peg_cmd_direct := exec.Command("./binary")
	err = peg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
