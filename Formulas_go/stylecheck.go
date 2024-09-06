package main

// StyleCheck - Parses latex-formatted text in search of forbidden phrases
// Homepage: https://www.cs.umd.edu/~nspring/software/style-check-readme.html

import (
	"fmt"
	
	"os/exec"
)

func installStyleCheck() {
	// Método 1: Descargar y extraer .tar.gz
	stylecheck_tar_url := "https://www.cs.umd.edu/~nspring/software/style-check-0.14.tar.gz"
	stylecheck_cmd_tar := exec.Command("curl", "-L", stylecheck_tar_url, "-o", "package.tar.gz")
	err := stylecheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stylecheck_zip_url := "https://www.cs.umd.edu/~nspring/software/style-check-0.14.zip"
	stylecheck_cmd_zip := exec.Command("curl", "-L", stylecheck_zip_url, "-o", "package.zip")
	err = stylecheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stylecheck_bin_url := "https://www.cs.umd.edu/~nspring/software/style-check-0.14.bin"
	stylecheck_cmd_bin := exec.Command("curl", "-L", stylecheck_bin_url, "-o", "binary.bin")
	err = stylecheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stylecheck_src_url := "https://www.cs.umd.edu/~nspring/software/style-check-0.14.src.tar.gz"
	stylecheck_cmd_src := exec.Command("curl", "-L", stylecheck_src_url, "-o", "source.tar.gz")
	err = stylecheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stylecheck_cmd_direct := exec.Command("./binary")
	err = stylecheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
