package main

// Openfast - NREL-supported OpenFAST whole-turbine simulation code
// Homepage: https://openfast.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installOpenfast() {
	// Método 1: Descargar y extraer .tar.gz
	openfast_tar_url := "https://github.com/openfast/openfast/archive/refs/tags/v3.5.3.tar.gz"
	openfast_cmd_tar := exec.Command("curl", "-L", openfast_tar_url, "-o", "package.tar.gz")
	err := openfast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openfast_zip_url := "https://github.com/openfast/openfast/archive/refs/tags/v3.5.3.zip"
	openfast_cmd_zip := exec.Command("curl", "-L", openfast_zip_url, "-o", "package.zip")
	err = openfast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openfast_bin_url := "https://github.com/openfast/openfast/archive/refs/tags/v3.5.3.bin"
	openfast_cmd_bin := exec.Command("curl", "-L", openfast_bin_url, "-o", "binary.bin")
	err = openfast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openfast_src_url := "https://github.com/openfast/openfast/archive/refs/tags/v3.5.3.src.tar.gz"
	openfast_cmd_src := exec.Command("curl", "-L", openfast_src_url, "-o", "source.tar.gz")
	err = openfast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openfast_cmd_direct := exec.Command("./binary")
	err = openfast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
