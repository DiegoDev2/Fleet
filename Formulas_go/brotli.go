package main

// Brotli - Generic-purpose lossless compression algorithm by Google
// Homepage: https://github.com/google/brotli

import (
	"fmt"
	
	"os/exec"
)

func installBrotli() {
	// Método 1: Descargar y extraer .tar.gz
	brotli_tar_url := "https://github.com/google/brotli/archive/refs/tags/v1.1.0.tar.gz"
	brotli_cmd_tar := exec.Command("curl", "-L", brotli_tar_url, "-o", "package.tar.gz")
	err := brotli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brotli_zip_url := "https://github.com/google/brotli/archive/refs/tags/v1.1.0.zip"
	brotli_cmd_zip := exec.Command("curl", "-L", brotli_zip_url, "-o", "package.zip")
	err = brotli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brotli_bin_url := "https://github.com/google/brotli/archive/refs/tags/v1.1.0.bin"
	brotli_cmd_bin := exec.Command("curl", "-L", brotli_bin_url, "-o", "binary.bin")
	err = brotli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brotli_src_url := "https://github.com/google/brotli/archive/refs/tags/v1.1.0.src.tar.gz"
	brotli_cmd_src := exec.Command("curl", "-L", brotli_src_url, "-o", "source.tar.gz")
	err = brotli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brotli_cmd_direct := exec.Command("./binary")
	err = brotli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
