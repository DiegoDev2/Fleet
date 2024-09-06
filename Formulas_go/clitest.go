package main

// Clitest - Command-Line Tester
// Homepage: https://github.com/aureliojargas/clitest

import (
	"fmt"
	
	"os/exec"
)

func installClitest() {
	// Método 1: Descargar y extraer .tar.gz
	clitest_tar_url := "https://github.com/aureliojargas/clitest/archive/refs/tags/0.5.0.tar.gz"
	clitest_cmd_tar := exec.Command("curl", "-L", clitest_tar_url, "-o", "package.tar.gz")
	err := clitest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clitest_zip_url := "https://github.com/aureliojargas/clitest/archive/refs/tags/0.5.0.zip"
	clitest_cmd_zip := exec.Command("curl", "-L", clitest_zip_url, "-o", "package.zip")
	err = clitest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clitest_bin_url := "https://github.com/aureliojargas/clitest/archive/refs/tags/0.5.0.bin"
	clitest_cmd_bin := exec.Command("curl", "-L", clitest_bin_url, "-o", "binary.bin")
	err = clitest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clitest_src_url := "https://github.com/aureliojargas/clitest/archive/refs/tags/0.5.0.src.tar.gz"
	clitest_cmd_src := exec.Command("curl", "-L", clitest_src_url, "-o", "source.tar.gz")
	err = clitest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clitest_cmd_direct := exec.Command("./binary")
	err = clitest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
