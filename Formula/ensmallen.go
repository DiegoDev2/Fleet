package main

// Ensmallen - Flexible C++ library for efficient mathematical optimization
// Homepage: https://ensmallen.org

import (
	"fmt"
	
	"os/exec"
)

func installEnsmallen() {
	// Método 1: Descargar y extraer .tar.gz
	ensmallen_tar_url := "https://github.com/mlpack/ensmallen/archive/refs/tags/2.21.1.tar.gz"
	ensmallen_cmd_tar := exec.Command("curl", "-L", ensmallen_tar_url, "-o", "package.tar.gz")
	err := ensmallen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ensmallen_zip_url := "https://github.com/mlpack/ensmallen/archive/refs/tags/2.21.1.zip"
	ensmallen_cmd_zip := exec.Command("curl", "-L", ensmallen_zip_url, "-o", "package.zip")
	err = ensmallen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ensmallen_bin_url := "https://github.com/mlpack/ensmallen/archive/refs/tags/2.21.1.bin"
	ensmallen_cmd_bin := exec.Command("curl", "-L", ensmallen_bin_url, "-o", "binary.bin")
	err = ensmallen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ensmallen_src_url := "https://github.com/mlpack/ensmallen/archive/refs/tags/2.21.1.src.tar.gz"
	ensmallen_cmd_src := exec.Command("curl", "-L", ensmallen_src_url, "-o", "source.tar.gz")
	err = ensmallen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ensmallen_cmd_direct := exec.Command("./binary")
	err = ensmallen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: armadillo")
	exec.Command("latte", "install", "armadillo").Run()
}
