package main

// St - Statistics from the command-line
// Homepage: https://github.com/nferraz/st

import (
	"fmt"
	
	"os/exec"
)

func installSt() {
	// Método 1: Descargar y extraer .tar.gz
	st_tar_url := "https://github.com/nferraz/st/archive/refs/tags/v1.1.4.tar.gz"
	st_cmd_tar := exec.Command("curl", "-L", st_tar_url, "-o", "package.tar.gz")
	err := st_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	st_zip_url := "https://github.com/nferraz/st/archive/refs/tags/v1.1.4.zip"
	st_cmd_zip := exec.Command("curl", "-L", st_zip_url, "-o", "package.zip")
	err = st_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	st_bin_url := "https://github.com/nferraz/st/archive/refs/tags/v1.1.4.bin"
	st_cmd_bin := exec.Command("curl", "-L", st_bin_url, "-o", "binary.bin")
	err = st_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	st_src_url := "https://github.com/nferraz/st/archive/refs/tags/v1.1.4.src.tar.gz"
	st_cmd_src := exec.Command("curl", "-L", st_src_url, "-o", "source.tar.gz")
	err = st_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	st_cmd_direct := exec.Command("./binary")
	err = st_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
