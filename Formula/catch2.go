package main

// Catch2 - Modern, C++-native, header-only, test framework
// Homepage: https://github.com/catchorg/Catch2

import (
	"fmt"
	
	"os/exec"
)

func installCatch2() {
	// Método 1: Descargar y extraer .tar.gz
	catch2_tar_url := "https://github.com/catchorg/Catch2/archive/refs/tags/v3.7.0.tar.gz"
	catch2_cmd_tar := exec.Command("curl", "-L", catch2_tar_url, "-o", "package.tar.gz")
	err := catch2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	catch2_zip_url := "https://github.com/catchorg/Catch2/archive/refs/tags/v3.7.0.zip"
	catch2_cmd_zip := exec.Command("curl", "-L", catch2_zip_url, "-o", "package.zip")
	err = catch2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	catch2_bin_url := "https://github.com/catchorg/Catch2/archive/refs/tags/v3.7.0.bin"
	catch2_cmd_bin := exec.Command("curl", "-L", catch2_bin_url, "-o", "binary.bin")
	err = catch2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	catch2_src_url := "https://github.com/catchorg/Catch2/archive/refs/tags/v3.7.0.src.tar.gz"
	catch2_cmd_src := exec.Command("curl", "-L", catch2_src_url, "-o", "source.tar.gz")
	err = catch2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	catch2_cmd_direct := exec.Command("./binary")
	err = catch2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
