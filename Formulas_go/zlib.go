package main

// Zlib - General-purpose lossless data-compression library
// Homepage: https://zlib.net/

import (
	"fmt"
	
	"os/exec"
)

func installZlib() {
	// Método 1: Descargar y extraer .tar.gz
	zlib_tar_url := "https://zlib.net/zlib-1.3.1.tar.gz"
	zlib_cmd_tar := exec.Command("curl", "-L", zlib_tar_url, "-o", "package.tar.gz")
	err := zlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zlib_zip_url := "https://zlib.net/zlib-1.3.1.zip"
	zlib_cmd_zip := exec.Command("curl", "-L", zlib_zip_url, "-o", "package.zip")
	err = zlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zlib_bin_url := "https://zlib.net/zlib-1.3.1.bin"
	zlib_cmd_bin := exec.Command("curl", "-L", zlib_bin_url, "-o", "binary.bin")
	err = zlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zlib_src_url := "https://zlib.net/zlib-1.3.1.src.tar.gz"
	zlib_cmd_src := exec.Command("curl", "-L", zlib_src_url, "-o", "source.tar.gz")
	err = zlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zlib_cmd_direct := exec.Command("./binary")
	err = zlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
