package main

// BoostBuild - C++ build system
// Homepage: https://www.boost.org/build/

import (
	"fmt"
	
	"os/exec"
)

func installBoostBuild() {
	// Método 1: Descargar y extraer .tar.gz
	boostbuild_tar_url := "https://github.com/boostorg/build/archive/refs/tags/boost-1.86.0.tar.gz"
	boostbuild_cmd_tar := exec.Command("curl", "-L", boostbuild_tar_url, "-o", "package.tar.gz")
	err := boostbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boostbuild_zip_url := "https://github.com/boostorg/build/archive/refs/tags/boost-1.86.0.zip"
	boostbuild_cmd_zip := exec.Command("curl", "-L", boostbuild_zip_url, "-o", "package.zip")
	err = boostbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boostbuild_bin_url := "https://github.com/boostorg/build/archive/refs/tags/boost-1.86.0.bin"
	boostbuild_cmd_bin := exec.Command("curl", "-L", boostbuild_bin_url, "-o", "binary.bin")
	err = boostbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boostbuild_src_url := "https://github.com/boostorg/build/archive/refs/tags/boost-1.86.0.src.tar.gz"
	boostbuild_cmd_src := exec.Command("curl", "-L", boostbuild_src_url, "-o", "source.tar.gz")
	err = boostbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boostbuild_cmd_direct := exec.Command("./binary")
	err = boostbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
