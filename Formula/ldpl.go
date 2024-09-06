package main

// Ldpl - Compiled programming language inspired by COBOL
// Homepage: https://www.ldpl-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installLdpl() {
	// Método 1: Descargar y extraer .tar.gz
	ldpl_tar_url := "https://github.com/Lartu/ldpl/archive/refs/tags/4.4.tar.gz"
	ldpl_cmd_tar := exec.Command("curl", "-L", ldpl_tar_url, "-o", "package.tar.gz")
	err := ldpl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ldpl_zip_url := "https://github.com/Lartu/ldpl/archive/refs/tags/4.4.zip"
	ldpl_cmd_zip := exec.Command("curl", "-L", ldpl_zip_url, "-o", "package.zip")
	err = ldpl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ldpl_bin_url := "https://github.com/Lartu/ldpl/archive/refs/tags/4.4.bin"
	ldpl_cmd_bin := exec.Command("curl", "-L", ldpl_bin_url, "-o", "binary.bin")
	err = ldpl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ldpl_src_url := "https://github.com/Lartu/ldpl/archive/refs/tags/4.4.src.tar.gz"
	ldpl_cmd_src := exec.Command("curl", "-L", ldpl_src_url, "-o", "source.tar.gz")
	err = ldpl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ldpl_cmd_direct := exec.Command("./binary")
	err = ldpl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
