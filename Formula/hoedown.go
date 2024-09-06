package main

// Hoedown - Secure Markdown processing (a revived fork of Sundown)
// Homepage: https://github.com/hoedown/hoedown

import (
	"fmt"
	
	"os/exec"
)

func installHoedown() {
	// Método 1: Descargar y extraer .tar.gz
	hoedown_tar_url := "https://github.com/hoedown/hoedown/archive/refs/tags/3.0.7.tar.gz"
	hoedown_cmd_tar := exec.Command("curl", "-L", hoedown_tar_url, "-o", "package.tar.gz")
	err := hoedown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hoedown_zip_url := "https://github.com/hoedown/hoedown/archive/refs/tags/3.0.7.zip"
	hoedown_cmd_zip := exec.Command("curl", "-L", hoedown_zip_url, "-o", "package.zip")
	err = hoedown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hoedown_bin_url := "https://github.com/hoedown/hoedown/archive/refs/tags/3.0.7.bin"
	hoedown_cmd_bin := exec.Command("curl", "-L", hoedown_bin_url, "-o", "binary.bin")
	err = hoedown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hoedown_src_url := "https://github.com/hoedown/hoedown/archive/refs/tags/3.0.7.src.tar.gz"
	hoedown_cmd_src := exec.Command("curl", "-L", hoedown_src_url, "-o", "source.tar.gz")
	err = hoedown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hoedown_cmd_direct := exec.Command("./binary")
	err = hoedown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
