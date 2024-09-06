package main

// Unittest - C++ Unit Test Framework
// Homepage: https://unittest.red-bean.com/

import (
	"fmt"
	
	"os/exec"
)

func installUnittest() {
	// Método 1: Descargar y extraer .tar.gz
	unittest_tar_url := "https://unittest.red-bean.com/tar/unittest-0.50-62.tar.gz"
	unittest_cmd_tar := exec.Command("curl", "-L", unittest_tar_url, "-o", "package.tar.gz")
	err := unittest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unittest_zip_url := "https://unittest.red-bean.com/tar/unittest-0.50-62.zip"
	unittest_cmd_zip := exec.Command("curl", "-L", unittest_zip_url, "-o", "package.zip")
	err = unittest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unittest_bin_url := "https://unittest.red-bean.com/tar/unittest-0.50-62.bin"
	unittest_cmd_bin := exec.Command("curl", "-L", unittest_bin_url, "-o", "binary.bin")
	err = unittest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unittest_src_url := "https://unittest.red-bean.com/tar/unittest-0.50-62.src.tar.gz"
	unittest_cmd_src := exec.Command("curl", "-L", unittest_src_url, "-o", "source.tar.gz")
	err = unittest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unittest_cmd_direct := exec.Command("./binary")
	err = unittest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
