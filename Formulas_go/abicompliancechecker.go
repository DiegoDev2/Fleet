package main

// AbiComplianceChecker - Tool for checking backward API/ABI compatibility of a C/C++ library
// Homepage: https://lvc.github.io/abi-compliance-checker/

import (
	"fmt"
	
	"os/exec"
)

func installAbiComplianceChecker() {
	// Método 1: Descargar y extraer .tar.gz
	abicompliancechecker_tar_url := "https://github.com/lvc/abi-compliance-checker/archive/refs/tags/2.3.tar.gz"
	abicompliancechecker_cmd_tar := exec.Command("curl", "-L", abicompliancechecker_tar_url, "-o", "package.tar.gz")
	err := abicompliancechecker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abicompliancechecker_zip_url := "https://github.com/lvc/abi-compliance-checker/archive/refs/tags/2.3.zip"
	abicompliancechecker_cmd_zip := exec.Command("curl", "-L", abicompliancechecker_zip_url, "-o", "package.zip")
	err = abicompliancechecker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abicompliancechecker_bin_url := "https://github.com/lvc/abi-compliance-checker/archive/refs/tags/2.3.bin"
	abicompliancechecker_cmd_bin := exec.Command("curl", "-L", abicompliancechecker_bin_url, "-o", "binary.bin")
	err = abicompliancechecker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abicompliancechecker_src_url := "https://github.com/lvc/abi-compliance-checker/archive/refs/tags/2.3.src.tar.gz"
	abicompliancechecker_cmd_src := exec.Command("curl", "-L", abicompliancechecker_src_url, "-o", "source.tar.gz")
	err = abicompliancechecker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abicompliancechecker_cmd_direct := exec.Command("./binary")
	err = abicompliancechecker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: universal-ctags")
exec.Command("latte", "install", "universal-ctags")
}
