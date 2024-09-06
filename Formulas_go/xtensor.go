package main

// Xtensor - Multi-dimensional arrays with broadcasting and lazy computing
// Homepage: https://xtensor.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installXtensor() {
	// Método 1: Descargar y extraer .tar.gz
	xtensor_tar_url := "https://github.com/xtensor-stack/xtensor/archive/refs/tags/0.25.0.tar.gz"
	xtensor_cmd_tar := exec.Command("curl", "-L", xtensor_tar_url, "-o", "package.tar.gz")
	err := xtensor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xtensor_zip_url := "https://github.com/xtensor-stack/xtensor/archive/refs/tags/0.25.0.zip"
	xtensor_cmd_zip := exec.Command("curl", "-L", xtensor_zip_url, "-o", "package.zip")
	err = xtensor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xtensor_bin_url := "https://github.com/xtensor-stack/xtensor/archive/refs/tags/0.25.0.bin"
	xtensor_cmd_bin := exec.Command("curl", "-L", xtensor_bin_url, "-o", "binary.bin")
	err = xtensor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xtensor_src_url := "https://github.com/xtensor-stack/xtensor/archive/refs/tags/0.25.0.src.tar.gz"
	xtensor_cmd_src := exec.Command("curl", "-L", xtensor_src_url, "-o", "source.tar.gz")
	err = xtensor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xtensor_cmd_direct := exec.Command("./binary")
	err = xtensor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
