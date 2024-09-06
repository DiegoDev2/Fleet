package main

// Csmith - Generates random C programs conforming to the C99 standard
// Homepage: https://github.com/csmith-project/csmith

import (
	"fmt"
	
	"os/exec"
)

func installCsmith() {
	// Método 1: Descargar y extraer .tar.gz
	csmith_tar_url := "https://github.com/csmith-project/csmith/archive/refs/tags/csmith-2.3.0.tar.gz"
	csmith_cmd_tar := exec.Command("curl", "-L", csmith_tar_url, "-o", "package.tar.gz")
	err := csmith_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csmith_zip_url := "https://github.com/csmith-project/csmith/archive/refs/tags/csmith-2.3.0.zip"
	csmith_cmd_zip := exec.Command("curl", "-L", csmith_zip_url, "-o", "package.zip")
	err = csmith_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csmith_bin_url := "https://github.com/csmith-project/csmith/archive/refs/tags/csmith-2.3.0.bin"
	csmith_cmd_bin := exec.Command("curl", "-L", csmith_bin_url, "-o", "binary.bin")
	err = csmith_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csmith_src_url := "https://github.com/csmith-project/csmith/archive/refs/tags/csmith-2.3.0.src.tar.gz"
	csmith_cmd_src := exec.Command("curl", "-L", csmith_src_url, "-o", "source.tar.gz")
	err = csmith_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csmith_cmd_direct := exec.Command("./binary")
	err = csmith_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
