package main

// Cpptest - Unit testing framework handling automated tests in C++
// Homepage: https://cpptest.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installCpptest() {
	// Método 1: Descargar y extraer .tar.gz
	cpptest_tar_url := "https://downloads.sourceforge.net/project/cpptest/cpptest/cpptest-2.0.0/cpptest-2.0.0.tar.bz2"
	cpptest_cmd_tar := exec.Command("curl", "-L", cpptest_tar_url, "-o", "package.tar.gz")
	err := cpptest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpptest_zip_url := "https://downloads.sourceforge.net/project/cpptest/cpptest/cpptest-2.0.0/cpptest-2.0.0.tar.bz2"
	cpptest_cmd_zip := exec.Command("curl", "-L", cpptest_zip_url, "-o", "package.zip")
	err = cpptest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpptest_bin_url := "https://downloads.sourceforge.net/project/cpptest/cpptest/cpptest-2.0.0/cpptest-2.0.0.tar.bz2"
	cpptest_cmd_bin := exec.Command("curl", "-L", cpptest_bin_url, "-o", "binary.bin")
	err = cpptest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpptest_src_url := "https://downloads.sourceforge.net/project/cpptest/cpptest/cpptest-2.0.0/cpptest-2.0.0.tar.bz2"
	cpptest_cmd_src := exec.Command("curl", "-L", cpptest_src_url, "-o", "source.tar.gz")
	err = cpptest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpptest_cmd_direct := exec.Command("./binary")
	err = cpptest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
