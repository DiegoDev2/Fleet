package main

// MagicEnum - Static reflection for enums (to string, from string, iteration) for modern C++
// Homepage: https://github.com/Neargye/magic_enum

import (
	"fmt"
	
	"os/exec"
)

func installMagicEnum() {
	// Método 1: Descargar y extraer .tar.gz
	magicenum_tar_url := "https://github.com/Neargye/magic_enum/archive/refs/tags/v0.9.6.tar.gz"
	magicenum_cmd_tar := exec.Command("curl", "-L", magicenum_tar_url, "-o", "package.tar.gz")
	err := magicenum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	magicenum_zip_url := "https://github.com/Neargye/magic_enum/archive/refs/tags/v0.9.6.zip"
	magicenum_cmd_zip := exec.Command("curl", "-L", magicenum_zip_url, "-o", "package.zip")
	err = magicenum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	magicenum_bin_url := "https://github.com/Neargye/magic_enum/archive/refs/tags/v0.9.6.bin"
	magicenum_cmd_bin := exec.Command("curl", "-L", magicenum_bin_url, "-o", "binary.bin")
	err = magicenum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	magicenum_src_url := "https://github.com/Neargye/magic_enum/archive/refs/tags/v0.9.6.src.tar.gz"
	magicenum_cmd_src := exec.Command("curl", "-L", magicenum_src_url, "-o", "source.tar.gz")
	err = magicenum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	magicenum_cmd_direct := exec.Command("./binary")
	err = magicenum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
