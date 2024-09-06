package main

// Pngpaste - Paste PNG into files
// Homepage: https://github.com/jcsalterego/pngpaste

import (
	"fmt"
	
	"os/exec"
)

func installPngpaste() {
	// Método 1: Descargar y extraer .tar.gz
	pngpaste_tar_url := "https://github.com/jcsalterego/pngpaste/archive/refs/tags/0.2.3.tar.gz"
	pngpaste_cmd_tar := exec.Command("curl", "-L", pngpaste_tar_url, "-o", "package.tar.gz")
	err := pngpaste_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pngpaste_zip_url := "https://github.com/jcsalterego/pngpaste/archive/refs/tags/0.2.3.zip"
	pngpaste_cmd_zip := exec.Command("curl", "-L", pngpaste_zip_url, "-o", "package.zip")
	err = pngpaste_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pngpaste_bin_url := "https://github.com/jcsalterego/pngpaste/archive/refs/tags/0.2.3.bin"
	pngpaste_cmd_bin := exec.Command("curl", "-L", pngpaste_bin_url, "-o", "binary.bin")
	err = pngpaste_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pngpaste_src_url := "https://github.com/jcsalterego/pngpaste/archive/refs/tags/0.2.3.src.tar.gz"
	pngpaste_cmd_src := exec.Command("curl", "-L", pngpaste_src_url, "-o", "source.tar.gz")
	err = pngpaste_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pngpaste_cmd_direct := exec.Command("./binary")
	err = pngpaste_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
