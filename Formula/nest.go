package main

// Nest - Neural Simulation Tool (NEST) with Python3 bindings (PyNEST)
// Homepage: https://www.nest-simulator.org/

import (
	"fmt"
	
	"os/exec"
)

func installNest() {
	// Método 1: Descargar y extraer .tar.gz
	nest_tar_url := "https://github.com/nest/nest-simulator/archive/refs/tags/v3.8.tar.gz"
	nest_cmd_tar := exec.Command("curl", "-L", nest_tar_url, "-o", "package.tar.gz")
	err := nest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nest_zip_url := "https://github.com/nest/nest-simulator/archive/refs/tags/v3.8.zip"
	nest_cmd_zip := exec.Command("curl", "-L", nest_zip_url, "-o", "package.zip")
	err = nest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nest_bin_url := "https://github.com/nest/nest-simulator/archive/refs/tags/v3.8.bin"
	nest_cmd_bin := exec.Command("curl", "-L", nest_bin_url, "-o", "binary.bin")
	err = nest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nest_src_url := "https://github.com/nest/nest-simulator/archive/refs/tags/v3.8.src.tar.gz"
	nest_cmd_src := exec.Command("curl", "-L", nest_src_url, "-o", "source.tar.gz")
	err = nest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nest_cmd_direct := exec.Command("./binary")
	err = nest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: gsl")
	exec.Command("latte", "install", "gsl").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
}
