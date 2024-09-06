package main

// Shunit2 - Unit testing framework for Bourne-based shell scripts
// Homepage: https://github.com/kward/shunit2

import (
	"fmt"
	
	"os/exec"
)

func installShunit2() {
	// Método 1: Descargar y extraer .tar.gz
	shunit2_tar_url := "https://github.com/kward/shunit2/archive/refs/tags/v2.1.8.tar.gz"
	shunit2_cmd_tar := exec.Command("curl", "-L", shunit2_tar_url, "-o", "package.tar.gz")
	err := shunit2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shunit2_zip_url := "https://github.com/kward/shunit2/archive/refs/tags/v2.1.8.zip"
	shunit2_cmd_zip := exec.Command("curl", "-L", shunit2_zip_url, "-o", "package.zip")
	err = shunit2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shunit2_bin_url := "https://github.com/kward/shunit2/archive/refs/tags/v2.1.8.bin"
	shunit2_cmd_bin := exec.Command("curl", "-L", shunit2_bin_url, "-o", "binary.bin")
	err = shunit2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shunit2_src_url := "https://github.com/kward/shunit2/archive/refs/tags/v2.1.8.src.tar.gz"
	shunit2_cmd_src := exec.Command("curl", "-L", shunit2_src_url, "-o", "source.tar.gz")
	err = shunit2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shunit2_cmd_direct := exec.Command("./binary")
	err = shunit2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
