package main

// Asciitex - Generate ASCII-art representations of mathematical equations
// Homepage: https://asciitex.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAsciitex() {
	// Método 1: Descargar y extraer .tar.gz
	asciitex_tar_url := "https://downloads.sourceforge.net/project/asciitex/asciiTeX-0.21.tar.gz"
	asciitex_cmd_tar := exec.Command("curl", "-L", asciitex_tar_url, "-o", "package.tar.gz")
	err := asciitex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asciitex_zip_url := "https://downloads.sourceforge.net/project/asciitex/asciiTeX-0.21.zip"
	asciitex_cmd_zip := exec.Command("curl", "-L", asciitex_zip_url, "-o", "package.zip")
	err = asciitex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asciitex_bin_url := "https://downloads.sourceforge.net/project/asciitex/asciiTeX-0.21.bin"
	asciitex_cmd_bin := exec.Command("curl", "-L", asciitex_bin_url, "-o", "binary.bin")
	err = asciitex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asciitex_src_url := "https://downloads.sourceforge.net/project/asciitex/asciiTeX-0.21.src.tar.gz"
	asciitex_cmd_src := exec.Command("curl", "-L", asciitex_src_url, "-o", "source.tar.gz")
	err = asciitex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asciitex_cmd_direct := exec.Command("./binary")
	err = asciitex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
