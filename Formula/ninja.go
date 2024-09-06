package main

// Ninja - Small build system for use with gyp or CMake
// Homepage: https://ninja-build.org/

import (
	"fmt"
	
	"os/exec"
)

func installNinja() {
	// Método 1: Descargar y extraer .tar.gz
	ninja_tar_url := "https://github.com/ninja-build/ninja/archive/refs/tags/v1.12.1.tar.gz"
	ninja_cmd_tar := exec.Command("curl", "-L", ninja_tar_url, "-o", "package.tar.gz")
	err := ninja_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ninja_zip_url := "https://github.com/ninja-build/ninja/archive/refs/tags/v1.12.1.zip"
	ninja_cmd_zip := exec.Command("curl", "-L", ninja_zip_url, "-o", "package.zip")
	err = ninja_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ninja_bin_url := "https://github.com/ninja-build/ninja/archive/refs/tags/v1.12.1.bin"
	ninja_cmd_bin := exec.Command("curl", "-L", ninja_bin_url, "-o", "binary.bin")
	err = ninja_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ninja_src_url := "https://github.com/ninja-build/ninja/archive/refs/tags/v1.12.1.src.tar.gz"
	ninja_cmd_src := exec.Command("curl", "-L", ninja_src_url, "-o", "source.tar.gz")
	err = ninja_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ninja_cmd_direct := exec.Command("./binary")
	err = ninja_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
