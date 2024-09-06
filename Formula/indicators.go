package main

// Indicators - Activity indicators for modern C++
// Homepage: https://github.com/p-ranav/indicators

import (
	"fmt"
	
	"os/exec"
)

func installIndicators() {
	// Método 1: Descargar y extraer .tar.gz
	indicators_tar_url := "https://github.com/p-ranav/indicators/archive/refs/tags/v2.3.tar.gz"
	indicators_cmd_tar := exec.Command("curl", "-L", indicators_tar_url, "-o", "package.tar.gz")
	err := indicators_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	indicators_zip_url := "https://github.com/p-ranav/indicators/archive/refs/tags/v2.3.zip"
	indicators_cmd_zip := exec.Command("curl", "-L", indicators_zip_url, "-o", "package.zip")
	err = indicators_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	indicators_bin_url := "https://github.com/p-ranav/indicators/archive/refs/tags/v2.3.bin"
	indicators_cmd_bin := exec.Command("curl", "-L", indicators_bin_url, "-o", "binary.bin")
	err = indicators_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	indicators_src_url := "https://github.com/p-ranav/indicators/archive/refs/tags/v2.3.src.tar.gz"
	indicators_cmd_src := exec.Command("curl", "-L", indicators_src_url, "-o", "source.tar.gz")
	err = indicators_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	indicators_cmd_direct := exec.Command("./binary")
	err = indicators_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
