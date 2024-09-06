package main

// Fypp - Python powered Fortran preprocessor
// Homepage: https://fypp.readthedocs.io/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installFypp() {
	// Método 1: Descargar y extraer .tar.gz
	fypp_tar_url := "https://files.pythonhosted.org/packages/01/35/0e2dfffc90201f17436d3416f8d5c8b00e2187e410ec899bb62cf2cea59b/fypp-3.2.tar.gz"
	fypp_cmd_tar := exec.Command("curl", "-L", fypp_tar_url, "-o", "package.tar.gz")
	err := fypp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fypp_zip_url := "https://files.pythonhosted.org/packages/01/35/0e2dfffc90201f17436d3416f8d5c8b00e2187e410ec899bb62cf2cea59b/fypp-3.2.zip"
	fypp_cmd_zip := exec.Command("curl", "-L", fypp_zip_url, "-o", "package.zip")
	err = fypp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fypp_bin_url := "https://files.pythonhosted.org/packages/01/35/0e2dfffc90201f17436d3416f8d5c8b00e2187e410ec899bb62cf2cea59b/fypp-3.2.bin"
	fypp_cmd_bin := exec.Command("curl", "-L", fypp_bin_url, "-o", "binary.bin")
	err = fypp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fypp_src_url := "https://files.pythonhosted.org/packages/01/35/0e2dfffc90201f17436d3416f8d5c8b00e2187e410ec899bb62cf2cea59b/fypp-3.2.src.tar.gz"
	fypp_cmd_src := exec.Command("curl", "-L", fypp_src_url, "-o", "source.tar.gz")
	err = fypp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fypp_cmd_direct := exec.Command("./binary")
	err = fypp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
