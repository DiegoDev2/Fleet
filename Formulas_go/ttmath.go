package main

// Ttmath - Bignum library for C++
// Homepage: https://www.ttmath.org/

import (
	"fmt"
	
	"os/exec"
)

func installTtmath() {
	// Método 1: Descargar y extraer .tar.gz
	ttmath_tar_url := "https://downloads.sourceforge.net/project/ttmath/ttmath/ttmath-0.9.3/ttmath-0.9.3-src.tar.gz"
	ttmath_cmd_tar := exec.Command("curl", "-L", ttmath_tar_url, "-o", "package.tar.gz")
	err := ttmath_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttmath_zip_url := "https://downloads.sourceforge.net/project/ttmath/ttmath/ttmath-0.9.3/ttmath-0.9.3-src.zip"
	ttmath_cmd_zip := exec.Command("curl", "-L", ttmath_zip_url, "-o", "package.zip")
	err = ttmath_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttmath_bin_url := "https://downloads.sourceforge.net/project/ttmath/ttmath/ttmath-0.9.3/ttmath-0.9.3-src.bin"
	ttmath_cmd_bin := exec.Command("curl", "-L", ttmath_bin_url, "-o", "binary.bin")
	err = ttmath_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttmath_src_url := "https://downloads.sourceforge.net/project/ttmath/ttmath/ttmath-0.9.3/ttmath-0.9.3-src.src.tar.gz"
	ttmath_cmd_src := exec.Command("curl", "-L", ttmath_src_url, "-o", "source.tar.gz")
	err = ttmath_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttmath_cmd_direct := exec.Command("./binary")
	err = ttmath_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
