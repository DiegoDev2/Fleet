package main

// Roundup - Unit testing tool
// Homepage: https://bmizerany.github.io/roundup

import (
	"fmt"
	
	"os/exec"
)

func installRoundup() {
	// Método 1: Descargar y extraer .tar.gz
	roundup_tar_url := "https://github.com/bmizerany/roundup/archive/refs/tags/v0.0.6.tar.gz"
	roundup_cmd_tar := exec.Command("curl", "-L", roundup_tar_url, "-o", "package.tar.gz")
	err := roundup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	roundup_zip_url := "https://github.com/bmizerany/roundup/archive/refs/tags/v0.0.6.zip"
	roundup_cmd_zip := exec.Command("curl", "-L", roundup_zip_url, "-o", "package.zip")
	err = roundup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	roundup_bin_url := "https://github.com/bmizerany/roundup/archive/refs/tags/v0.0.6.bin"
	roundup_cmd_bin := exec.Command("curl", "-L", roundup_bin_url, "-o", "binary.bin")
	err = roundup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	roundup_src_url := "https://github.com/bmizerany/roundup/archive/refs/tags/v0.0.6.src.tar.gz"
	roundup_cmd_src := exec.Command("curl", "-L", roundup_src_url, "-o", "source.tar.gz")
	err = roundup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	roundup_cmd_direct := exec.Command("./binary")
	err = roundup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
