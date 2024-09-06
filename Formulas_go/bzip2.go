package main

// Bzip2 - Freely available high-quality data compressor
// Homepage: https://sourceware.org/bzip2/

import (
	"fmt"
	
	"os/exec"
)

func installBzip2() {
	// Método 1: Descargar y extraer .tar.gz
	bzip2_tar_url := "https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz"
	bzip2_cmd_tar := exec.Command("curl", "-L", bzip2_tar_url, "-o", "package.tar.gz")
	err := bzip2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bzip2_zip_url := "https://sourceware.org/pub/bzip2/bzip2-1.0.8.zip"
	bzip2_cmd_zip := exec.Command("curl", "-L", bzip2_zip_url, "-o", "package.zip")
	err = bzip2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bzip2_bin_url := "https://sourceware.org/pub/bzip2/bzip2-1.0.8.bin"
	bzip2_cmd_bin := exec.Command("curl", "-L", bzip2_bin_url, "-o", "binary.bin")
	err = bzip2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bzip2_src_url := "https://sourceware.org/pub/bzip2/bzip2-1.0.8.src.tar.gz"
	bzip2_cmd_src := exec.Command("curl", "-L", bzip2_src_url, "-o", "source.tar.gz")
	err = bzip2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bzip2_cmd_direct := exec.Command("./binary")
	err = bzip2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
