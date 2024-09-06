package main

// Functionalplus - Functional Programming Library for C++
// Homepage: https://github.com/Dobiasd/FunctionalPlus

import (
	"fmt"
	
	"os/exec"
)

func installFunctionalplus() {
	// Método 1: Descargar y extraer .tar.gz
	functionalplus_tar_url := "https://github.com/Dobiasd/FunctionalPlus/archive/refs/tags/v0.2.25.tar.gz"
	functionalplus_cmd_tar := exec.Command("curl", "-L", functionalplus_tar_url, "-o", "package.tar.gz")
	err := functionalplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	functionalplus_zip_url := "https://github.com/Dobiasd/FunctionalPlus/archive/refs/tags/v0.2.25.zip"
	functionalplus_cmd_zip := exec.Command("curl", "-L", functionalplus_zip_url, "-o", "package.zip")
	err = functionalplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	functionalplus_bin_url := "https://github.com/Dobiasd/FunctionalPlus/archive/refs/tags/v0.2.25.bin"
	functionalplus_cmd_bin := exec.Command("curl", "-L", functionalplus_bin_url, "-o", "binary.bin")
	err = functionalplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	functionalplus_src_url := "https://github.com/Dobiasd/FunctionalPlus/archive/refs/tags/v0.2.25.src.tar.gz"
	functionalplus_cmd_src := exec.Command("curl", "-L", functionalplus_src_url, "-o", "source.tar.gz")
	err = functionalplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	functionalplus_cmd_direct := exec.Command("./binary")
	err = functionalplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
