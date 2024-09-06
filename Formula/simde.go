package main

// Simde - Implementations of SIMD intrinsics for systems which don't natively support them
// Homepage: https://github.com/simd-everywhere/simde

import (
	"fmt"
	
	"os/exec"
)

func installSimde() {
	// Método 1: Descargar y extraer .tar.gz
	simde_tar_url := "https://github.com/simd-everywhere/simde/archive/refs/tags/v0.8.2.tar.gz"
	simde_cmd_tar := exec.Command("curl", "-L", simde_tar_url, "-o", "package.tar.gz")
	err := simde_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simde_zip_url := "https://github.com/simd-everywhere/simde/archive/refs/tags/v0.8.2.zip"
	simde_cmd_zip := exec.Command("curl", "-L", simde_zip_url, "-o", "package.zip")
	err = simde_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simde_bin_url := "https://github.com/simd-everywhere/simde/archive/refs/tags/v0.8.2.bin"
	simde_cmd_bin := exec.Command("curl", "-L", simde_bin_url, "-o", "binary.bin")
	err = simde_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simde_src_url := "https://github.com/simd-everywhere/simde/archive/refs/tags/v0.8.2.src.tar.gz"
	simde_cmd_src := exec.Command("curl", "-L", simde_src_url, "-o", "source.tar.gz")
	err = simde_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simde_cmd_direct := exec.Command("./binary")
	err = simde_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
}
