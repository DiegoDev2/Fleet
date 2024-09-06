package main

// Googletest - Google Testing and Mocking Framework
// Homepage: https://github.com/google/googletest

import (
	"fmt"
	
	"os/exec"
)

func installGoogletest() {
	// Método 1: Descargar y extraer .tar.gz
	googletest_tar_url := "https://github.com/google/googletest/archive/refs/tags/v1.15.2.tar.gz"
	googletest_cmd_tar := exec.Command("curl", "-L", googletest_tar_url, "-o", "package.tar.gz")
	err := googletest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	googletest_zip_url := "https://github.com/google/googletest/archive/refs/tags/v1.15.2.zip"
	googletest_cmd_zip := exec.Command("curl", "-L", googletest_zip_url, "-o", "package.zip")
	err = googletest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	googletest_bin_url := "https://github.com/google/googletest/archive/refs/tags/v1.15.2.bin"
	googletest_cmd_bin := exec.Command("curl", "-L", googletest_bin_url, "-o", "binary.bin")
	err = googletest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	googletest_src_url := "https://github.com/google/googletest/archive/refs/tags/v1.15.2.src.tar.gz"
	googletest_cmd_src := exec.Command("curl", "-L", googletest_src_url, "-o", "source.tar.gz")
	err = googletest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	googletest_cmd_direct := exec.Command("./binary")
	err = googletest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
