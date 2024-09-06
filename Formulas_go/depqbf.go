package main

// Depqbf - Solver for quantified boolean formulae (QBF)
// Homepage: https://lonsing.github.io/depqbf/

import (
	"fmt"
	
	"os/exec"
)

func installDepqbf() {
	// Método 1: Descargar y extraer .tar.gz
	depqbf_tar_url := "https://github.com/lonsing/depqbf/archive/refs/tags/version-6.03.tar.gz"
	depqbf_cmd_tar := exec.Command("curl", "-L", depqbf_tar_url, "-o", "package.tar.gz")
	err := depqbf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	depqbf_zip_url := "https://github.com/lonsing/depqbf/archive/refs/tags/version-6.03.zip"
	depqbf_cmd_zip := exec.Command("curl", "-L", depqbf_zip_url, "-o", "package.zip")
	err = depqbf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	depqbf_bin_url := "https://github.com/lonsing/depqbf/archive/refs/tags/version-6.03.bin"
	depqbf_cmd_bin := exec.Command("curl", "-L", depqbf_bin_url, "-o", "binary.bin")
	err = depqbf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	depqbf_src_url := "https://github.com/lonsing/depqbf/archive/refs/tags/version-6.03.src.tar.gz"
	depqbf_cmd_src := exec.Command("curl", "-L", depqbf_src_url, "-o", "source.tar.gz")
	err = depqbf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	depqbf_cmd_direct := exec.Command("./binary")
	err = depqbf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
