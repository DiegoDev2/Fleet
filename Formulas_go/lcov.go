package main

// Lcov - Graphical front-end for GCC's coverage testing tool (gcov)
// Homepage: https://github.com/linux-test-project/lcov

import (
	"fmt"
	
	"os/exec"
)

func installLcov() {
	// Método 1: Descargar y extraer .tar.gz
	lcov_tar_url := "https://github.com/linux-test-project/lcov/releases/download/v2.1/lcov-2.1.tar.gz"
	lcov_cmd_tar := exec.Command("curl", "-L", lcov_tar_url, "-o", "package.tar.gz")
	err := lcov_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lcov_zip_url := "https://github.com/linux-test-project/lcov/releases/download/v2.1/lcov-2.1.zip"
	lcov_cmd_zip := exec.Command("curl", "-L", lcov_zip_url, "-o", "package.zip")
	err = lcov_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lcov_bin_url := "https://github.com/linux-test-project/lcov/releases/download/v2.1/lcov-2.1.bin"
	lcov_cmd_bin := exec.Command("curl", "-L", lcov_bin_url, "-o", "binary.bin")
	err = lcov_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lcov_src_url := "https://github.com/linux-test-project/lcov/releases/download/v2.1/lcov-2.1.src.tar.gz"
	lcov_cmd_src := exec.Command("curl", "-L", lcov_src_url, "-o", "source.tar.gz")
	err = lcov_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lcov_cmd_direct := exec.Command("./binary")
	err = lcov_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
