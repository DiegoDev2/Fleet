package main

// Xsimd - Modern, portable C++ wrappers for SIMD intrinsics
// Homepage: https://xsimd.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installXsimd() {
	// Método 1: Descargar y extraer .tar.gz
	xsimd_tar_url := "https://github.com/xtensor-stack/xsimd/archive/refs/tags/13.0.0.tar.gz"
	xsimd_cmd_tar := exec.Command("curl", "-L", xsimd_tar_url, "-o", "package.tar.gz")
	err := xsimd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xsimd_zip_url := "https://github.com/xtensor-stack/xsimd/archive/refs/tags/13.0.0.zip"
	xsimd_cmd_zip := exec.Command("curl", "-L", xsimd_zip_url, "-o", "package.zip")
	err = xsimd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xsimd_bin_url := "https://github.com/xtensor-stack/xsimd/archive/refs/tags/13.0.0.bin"
	xsimd_cmd_bin := exec.Command("curl", "-L", xsimd_bin_url, "-o", "binary.bin")
	err = xsimd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xsimd_src_url := "https://github.com/xtensor-stack/xsimd/archive/refs/tags/13.0.0.src.tar.gz"
	xsimd_cmd_src := exec.Command("curl", "-L", xsimd_src_url, "-o", "source.tar.gz")
	err = xsimd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xsimd_cmd_direct := exec.Command("./binary")
	err = xsimd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
