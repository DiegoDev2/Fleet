package main

// DoubleConversion - Binary-decimal and decimal-binary routines for IEEE doubles
// Homepage: https://github.com/google/double-conversion

import (
	"fmt"
	
	"os/exec"
)

func installDoubleConversion() {
	// Método 1: Descargar y extraer .tar.gz
	doubleconversion_tar_url := "https://github.com/google/double-conversion/archive/refs/tags/v3.3.0.tar.gz"
	doubleconversion_cmd_tar := exec.Command("curl", "-L", doubleconversion_tar_url, "-o", "package.tar.gz")
	err := doubleconversion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doubleconversion_zip_url := "https://github.com/google/double-conversion/archive/refs/tags/v3.3.0.zip"
	doubleconversion_cmd_zip := exec.Command("curl", "-L", doubleconversion_zip_url, "-o", "package.zip")
	err = doubleconversion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doubleconversion_bin_url := "https://github.com/google/double-conversion/archive/refs/tags/v3.3.0.bin"
	doubleconversion_cmd_bin := exec.Command("curl", "-L", doubleconversion_bin_url, "-o", "binary.bin")
	err = doubleconversion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doubleconversion_src_url := "https://github.com/google/double-conversion/archive/refs/tags/v3.3.0.src.tar.gz"
	doubleconversion_cmd_src := exec.Command("curl", "-L", doubleconversion_src_url, "-o", "source.tar.gz")
	err = doubleconversion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doubleconversion_cmd_direct := exec.Command("./binary")
	err = doubleconversion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
