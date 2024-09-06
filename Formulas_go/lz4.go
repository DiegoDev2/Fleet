package main

// Lz4 - Extremely Fast Compression algorithm
// Homepage: https://lz4.github.io/lz4/

import (
	"fmt"
	
	"os/exec"
)

func installLz4() {
	// Método 1: Descargar y extraer .tar.gz
	lz4_tar_url := "https://github.com/lz4/lz4/archive/refs/tags/v1.10.0.tar.gz"
	lz4_cmd_tar := exec.Command("curl", "-L", lz4_tar_url, "-o", "package.tar.gz")
	err := lz4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lz4_zip_url := "https://github.com/lz4/lz4/archive/refs/tags/v1.10.0.zip"
	lz4_cmd_zip := exec.Command("curl", "-L", lz4_zip_url, "-o", "package.zip")
	err = lz4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lz4_bin_url := "https://github.com/lz4/lz4/archive/refs/tags/v1.10.0.bin"
	lz4_cmd_bin := exec.Command("curl", "-L", lz4_bin_url, "-o", "binary.bin")
	err = lz4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lz4_src_url := "https://github.com/lz4/lz4/archive/refs/tags/v1.10.0.src.tar.gz"
	lz4_cmd_src := exec.Command("curl", "-L", lz4_src_url, "-o", "source.tar.gz")
	err = lz4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lz4_cmd_direct := exec.Command("./binary")
	err = lz4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
