package main

// Bzip3 - Better and stronger spiritual successor to BZip2
// Homepage: https://github.com/kspalaiologos/bzip3

import (
	"fmt"
	
	"os/exec"
)

func installBzip3() {
	// Método 1: Descargar y extraer .tar.gz
	bzip3_tar_url := "https://github.com/kspalaiologos/bzip3/releases/download/1.4.0/bzip3-1.4.0.tar.gz"
	bzip3_cmd_tar := exec.Command("curl", "-L", bzip3_tar_url, "-o", "package.tar.gz")
	err := bzip3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bzip3_zip_url := "https://github.com/kspalaiologos/bzip3/releases/download/1.4.0/bzip3-1.4.0.zip"
	bzip3_cmd_zip := exec.Command("curl", "-L", bzip3_zip_url, "-o", "package.zip")
	err = bzip3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bzip3_bin_url := "https://github.com/kspalaiologos/bzip3/releases/download/1.4.0/bzip3-1.4.0.bin"
	bzip3_cmd_bin := exec.Command("curl", "-L", bzip3_bin_url, "-o", "binary.bin")
	err = bzip3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bzip3_src_url := "https://github.com/kspalaiologos/bzip3/releases/download/1.4.0/bzip3-1.4.0.src.tar.gz"
	bzip3_cmd_src := exec.Command("curl", "-L", bzip3_src_url, "-o", "source.tar.gz")
	err = bzip3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bzip3_cmd_direct := exec.Command("./binary")
	err = bzip3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
