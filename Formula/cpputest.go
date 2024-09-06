package main

// Cpputest - C /C++ based unit xUnit test framework
// Homepage: https://www.cpputest.org/

import (
	"fmt"
	
	"os/exec"
)

func installCpputest() {
	// Método 1: Descargar y extraer .tar.gz
	cpputest_tar_url := "https://github.com/cpputest/cpputest/releases/download/v4.0/cpputest-4.0.tar.gz"
	cpputest_cmd_tar := exec.Command("curl", "-L", cpputest_tar_url, "-o", "package.tar.gz")
	err := cpputest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpputest_zip_url := "https://github.com/cpputest/cpputest/releases/download/v4.0/cpputest-4.0.zip"
	cpputest_cmd_zip := exec.Command("curl", "-L", cpputest_zip_url, "-o", "package.zip")
	err = cpputest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpputest_bin_url := "https://github.com/cpputest/cpputest/releases/download/v4.0/cpputest-4.0.bin"
	cpputest_cmd_bin := exec.Command("curl", "-L", cpputest_bin_url, "-o", "binary.bin")
	err = cpputest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpputest_src_url := "https://github.com/cpputest/cpputest/releases/download/v4.0/cpputest-4.0.src.tar.gz"
	cpputest_cmd_src := exec.Command("curl", "-L", cpputest_src_url, "-o", "source.tar.gz")
	err = cpputest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpputest_cmd_direct := exec.Command("./binary")
	err = cpputest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
