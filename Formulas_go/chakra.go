package main

// Chakra - Core part of the JavaScript engine that powers Microsoft Edge
// Homepage: https://github.com/chakra-core/ChakraCore

import (
	"fmt"
	
	"os/exec"
)

func installChakra() {
	// Método 1: Descargar y extraer .tar.gz
	chakra_tar_url := "https://github.com/chakra-core/ChakraCore/archive/refs/tags/v1.11.24.tar.gz"
	chakra_cmd_tar := exec.Command("curl", "-L", chakra_tar_url, "-o", "package.tar.gz")
	err := chakra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chakra_zip_url := "https://github.com/chakra-core/ChakraCore/archive/refs/tags/v1.11.24.zip"
	chakra_cmd_zip := exec.Command("curl", "-L", chakra_zip_url, "-o", "package.zip")
	err = chakra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chakra_bin_url := "https://github.com/chakra-core/ChakraCore/archive/refs/tags/v1.11.24.bin"
	chakra_cmd_bin := exec.Command("curl", "-L", chakra_bin_url, "-o", "binary.bin")
	err = chakra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chakra_src_url := "https://github.com/chakra-core/ChakraCore/archive/refs/tags/v1.11.24.src.tar.gz"
	chakra_cmd_src := exec.Command("curl", "-L", chakra_src_url, "-o", "source.tar.gz")
	err = chakra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chakra_cmd_direct := exec.Command("./binary")
	err = chakra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: llvm@15")
exec.Command("latte", "install", "llvm@15")
}
