package main

// Clac - Command-line, stack-based calculator with postfix notation
// Homepage: https://github.com/soveran/clac

import (
	"fmt"
	
	"os/exec"
)

func installClac() {
	// Método 1: Descargar y extraer .tar.gz
	clac_tar_url := "https://github.com/soveran/clac/archive/refs/tags/0.3.3.tar.gz"
	clac_cmd_tar := exec.Command("curl", "-L", clac_tar_url, "-o", "package.tar.gz")
	err := clac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clac_zip_url := "https://github.com/soveran/clac/archive/refs/tags/0.3.3.zip"
	clac_cmd_zip := exec.Command("curl", "-L", clac_zip_url, "-o", "package.zip")
	err = clac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clac_bin_url := "https://github.com/soveran/clac/archive/refs/tags/0.3.3.bin"
	clac_cmd_bin := exec.Command("curl", "-L", clac_bin_url, "-o", "binary.bin")
	err = clac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clac_src_url := "https://github.com/soveran/clac/archive/refs/tags/0.3.3.src.tar.gz"
	clac_cmd_src := exec.Command("curl", "-L", clac_src_url, "-o", "source.tar.gz")
	err = clac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clac_cmd_direct := exec.Command("./binary")
	err = clac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
