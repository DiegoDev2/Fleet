package main

// ZlibNg - Zlib replacement with optimizations for next generation systems
// Homepage: https://github.com/zlib-ng/zlib-ng

import (
	"fmt"
	
	"os/exec"
)

func installZlibNg() {
	// Método 1: Descargar y extraer .tar.gz
	zlibng_tar_url := "https://github.com/zlib-ng/zlib-ng/archive/refs/tags/2.2.1.tar.gz"
	zlibng_cmd_tar := exec.Command("curl", "-L", zlibng_tar_url, "-o", "package.tar.gz")
	err := zlibng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zlibng_zip_url := "https://github.com/zlib-ng/zlib-ng/archive/refs/tags/2.2.1.zip"
	zlibng_cmd_zip := exec.Command("curl", "-L", zlibng_zip_url, "-o", "package.zip")
	err = zlibng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zlibng_bin_url := "https://github.com/zlib-ng/zlib-ng/archive/refs/tags/2.2.1.bin"
	zlibng_cmd_bin := exec.Command("curl", "-L", zlibng_bin_url, "-o", "binary.bin")
	err = zlibng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zlibng_src_url := "https://github.com/zlib-ng/zlib-ng/archive/refs/tags/2.2.1.src.tar.gz"
	zlibng_cmd_src := exec.Command("curl", "-L", zlibng_src_url, "-o", "source.tar.gz")
	err = zlibng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zlibng_cmd_direct := exec.Command("./binary")
	err = zlibng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
