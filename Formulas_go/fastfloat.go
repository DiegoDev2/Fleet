package main

// FastFloat - Fast and exact implementation of the C++ from_chars functions for number types
// Homepage: https://github.com/fastfloat/fast_float

import (
	"fmt"
	
	"os/exec"
)

func installFastFloat() {
	// Método 1: Descargar y extraer .tar.gz
	fastfloat_tar_url := "https://github.com/fastfloat/fast_float/archive/refs/tags/v6.1.5.tar.gz"
	fastfloat_cmd_tar := exec.Command("curl", "-L", fastfloat_tar_url, "-o", "package.tar.gz")
	err := fastfloat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastfloat_zip_url := "https://github.com/fastfloat/fast_float/archive/refs/tags/v6.1.5.zip"
	fastfloat_cmd_zip := exec.Command("curl", "-L", fastfloat_zip_url, "-o", "package.zip")
	err = fastfloat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastfloat_bin_url := "https://github.com/fastfloat/fast_float/archive/refs/tags/v6.1.5.bin"
	fastfloat_cmd_bin := exec.Command("curl", "-L", fastfloat_bin_url, "-o", "binary.bin")
	err = fastfloat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastfloat_src_url := "https://github.com/fastfloat/fast_float/archive/refs/tags/v6.1.5.src.tar.gz"
	fastfloat_cmd_src := exec.Command("curl", "-L", fastfloat_src_url, "-o", "source.tar.gz")
	err = fastfloat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastfloat_cmd_direct := exec.Command("./binary")
	err = fastfloat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
