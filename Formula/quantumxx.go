package main

// Quantumxx - Modern C++ quantum computing library
// Homepage: https://github.com/softwareQinc/qpp

import (
	"fmt"
	
	"os/exec"
)

func installQuantumxx() {
	// Método 1: Descargar y extraer .tar.gz
	quantumxx_tar_url := "https://github.com/softwareQinc/qpp/archive/refs/tags/v5.1.tar.gz"
	quantumxx_cmd_tar := exec.Command("curl", "-L", quantumxx_tar_url, "-o", "package.tar.gz")
	err := quantumxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quantumxx_zip_url := "https://github.com/softwareQinc/qpp/archive/refs/tags/v5.1.zip"
	quantumxx_cmd_zip := exec.Command("curl", "-L", quantumxx_zip_url, "-o", "package.zip")
	err = quantumxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quantumxx_bin_url := "https://github.com/softwareQinc/qpp/archive/refs/tags/v5.1.bin"
	quantumxx_cmd_bin := exec.Command("curl", "-L", quantumxx_bin_url, "-o", "binary.bin")
	err = quantumxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quantumxx_src_url := "https://github.com/softwareQinc/qpp/archive/refs/tags/v5.1.src.tar.gz"
	quantumxx_cmd_src := exec.Command("curl", "-L", quantumxx_src_url, "-o", "source.tar.gz")
	err = quantumxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quantumxx_cmd_direct := exec.Command("./binary")
	err = quantumxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
}
