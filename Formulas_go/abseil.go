package main

// Abseil - C++ Common Libraries
// Homepage: https://abseil.io

import (
	"fmt"
	
	"os/exec"
)

func installAbseil() {
	// Método 1: Descargar y extraer .tar.gz
	abseil_tar_url := "https://github.com/abseil/abseil-cpp/archive/refs/tags/20240722.0.tar.gz"
	abseil_cmd_tar := exec.Command("curl", "-L", abseil_tar_url, "-o", "package.tar.gz")
	err := abseil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abseil_zip_url := "https://github.com/abseil/abseil-cpp/archive/refs/tags/20240722.0.zip"
	abseil_cmd_zip := exec.Command("curl", "-L", abseil_zip_url, "-o", "package.zip")
	err = abseil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abseil_bin_url := "https://github.com/abseil/abseil-cpp/archive/refs/tags/20240722.0.bin"
	abseil_cmd_bin := exec.Command("curl", "-L", abseil_bin_url, "-o", "binary.bin")
	err = abseil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abseil_src_url := "https://github.com/abseil/abseil-cpp/archive/refs/tags/20240722.0.src.tar.gz"
	abseil_cmd_src := exec.Command("curl", "-L", abseil_src_url, "-o", "source.tar.gz")
	err = abseil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abseil_cmd_direct := exec.Command("./binary")
	err = abseil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
}
