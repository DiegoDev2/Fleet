package main

// PhoronixTestSuite - Open-source automated testing/benchmarking software
// Homepage: https://www.phoronix-test-suite.com/

import (
	"fmt"
	
	"os/exec"
)

func installPhoronixTestSuite() {
	// Método 1: Descargar y extraer .tar.gz
	phoronixtestsuite_tar_url := "https://github.com/phoronix-test-suite/phoronix-test-suite/archive/refs/tags/v10.8.4.tar.gz"
	phoronixtestsuite_cmd_tar := exec.Command("curl", "-L", phoronixtestsuite_tar_url, "-o", "package.tar.gz")
	err := phoronixtestsuite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phoronixtestsuite_zip_url := "https://github.com/phoronix-test-suite/phoronix-test-suite/archive/refs/tags/v10.8.4.zip"
	phoronixtestsuite_cmd_zip := exec.Command("curl", "-L", phoronixtestsuite_zip_url, "-o", "package.zip")
	err = phoronixtestsuite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phoronixtestsuite_bin_url := "https://github.com/phoronix-test-suite/phoronix-test-suite/archive/refs/tags/v10.8.4.bin"
	phoronixtestsuite_cmd_bin := exec.Command("curl", "-L", phoronixtestsuite_bin_url, "-o", "binary.bin")
	err = phoronixtestsuite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phoronixtestsuite_src_url := "https://github.com/phoronix-test-suite/phoronix-test-suite/archive/refs/tags/v10.8.4.src.tar.gz"
	phoronixtestsuite_cmd_src := exec.Command("curl", "-L", phoronixtestsuite_src_url, "-o", "source.tar.gz")
	err = phoronixtestsuite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phoronixtestsuite_cmd_direct := exec.Command("./binary")
	err = phoronixtestsuite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
