package main

// Simdutf - Unicode conversion routines, fast
// Homepage: https://github.com/simdutf/simdutf

import (
	"fmt"
	
	"os/exec"
)

func installSimdutf() {
	// Método 1: Descargar y extraer .tar.gz
	simdutf_tar_url := "https://github.com/simdutf/simdutf/archive/refs/tags/v5.5.0.tar.gz"
	simdutf_cmd_tar := exec.Command("curl", "-L", simdutf_tar_url, "-o", "package.tar.gz")
	err := simdutf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simdutf_zip_url := "https://github.com/simdutf/simdutf/archive/refs/tags/v5.5.0.zip"
	simdutf_cmd_zip := exec.Command("curl", "-L", simdutf_zip_url, "-o", "package.zip")
	err = simdutf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simdutf_bin_url := "https://github.com/simdutf/simdutf/archive/refs/tags/v5.5.0.bin"
	simdutf_cmd_bin := exec.Command("curl", "-L", simdutf_bin_url, "-o", "binary.bin")
	err = simdutf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simdutf_src_url := "https://github.com/simdutf/simdutf/archive/refs/tags/v5.5.0.src.tar.gz"
	simdutf_cmd_src := exec.Command("curl", "-L", simdutf_src_url, "-o", "source.tar.gz")
	err = simdutf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simdutf_cmd_direct := exec.Command("./binary")
	err = simdutf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
}
