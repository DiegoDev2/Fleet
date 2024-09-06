package main

// Cweb - Literate documentation system for C, C++, and Java
// Homepage: https://cs.stanford.edu/~knuth/cweb.html

import (
	"fmt"
	
	"os/exec"
)

func installCweb() {
	// Método 1: Descargar y extraer .tar.gz
	cweb_tar_url := "https://github.com/ascherer/cweb/archive/refs/tags/cweb-4.12.tar.gz"
	cweb_cmd_tar := exec.Command("curl", "-L", cweb_tar_url, "-o", "package.tar.gz")
	err := cweb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cweb_zip_url := "https://github.com/ascherer/cweb/archive/refs/tags/cweb-4.12.zip"
	cweb_cmd_zip := exec.Command("curl", "-L", cweb_zip_url, "-o", "package.zip")
	err = cweb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cweb_bin_url := "https://github.com/ascherer/cweb/archive/refs/tags/cweb-4.12.bin"
	cweb_cmd_bin := exec.Command("curl", "-L", cweb_bin_url, "-o", "binary.bin")
	err = cweb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cweb_src_url := "https://github.com/ascherer/cweb/archive/refs/tags/cweb-4.12.src.tar.gz"
	cweb_cmd_src := exec.Command("curl", "-L", cweb_src_url, "-o", "source.tar.gz")
	err = cweb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cweb_cmd_direct := exec.Command("./binary")
	err = cweb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
