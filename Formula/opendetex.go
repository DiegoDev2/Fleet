package main

// Opendetex - Tool to strip TeX or LaTeX commands from documents
// Homepage: https://github.com/pkubowicz/opendetex

import (
	"fmt"
	
	"os/exec"
)

func installOpendetex() {
	// Método 1: Descargar y extraer .tar.gz
	opendetex_tar_url := "https://github.com/pkubowicz/opendetex/releases/download/v2.8.11/opendetex-2.8.11.tar.bz2"
	opendetex_cmd_tar := exec.Command("curl", "-L", opendetex_tar_url, "-o", "package.tar.gz")
	err := opendetex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opendetex_zip_url := "https://github.com/pkubowicz/opendetex/releases/download/v2.8.11/opendetex-2.8.11.tar.bz2"
	opendetex_cmd_zip := exec.Command("curl", "-L", opendetex_zip_url, "-o", "package.zip")
	err = opendetex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opendetex_bin_url := "https://github.com/pkubowicz/opendetex/releases/download/v2.8.11/opendetex-2.8.11.tar.bz2"
	opendetex_cmd_bin := exec.Command("curl", "-L", opendetex_bin_url, "-o", "binary.bin")
	err = opendetex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opendetex_src_url := "https://github.com/pkubowicz/opendetex/releases/download/v2.8.11/opendetex-2.8.11.tar.bz2"
	opendetex_cmd_src := exec.Command("curl", "-L", opendetex_src_url, "-o", "source.tar.gz")
	err = opendetex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opendetex_cmd_direct := exec.Command("./binary")
	err = opendetex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
