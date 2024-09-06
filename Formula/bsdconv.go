package main

// Bsdconv - Charset/encoding converter library
// Homepage: https://github.com/buganini/bsdconv

import (
	"fmt"
	
	"os/exec"
)

func installBsdconv() {
	// Método 1: Descargar y extraer .tar.gz
	bsdconv_tar_url := "https://github.com/buganini/bsdconv/archive/refs/tags/11.6.tar.gz"
	bsdconv_cmd_tar := exec.Command("curl", "-L", bsdconv_tar_url, "-o", "package.tar.gz")
	err := bsdconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bsdconv_zip_url := "https://github.com/buganini/bsdconv/archive/refs/tags/11.6.zip"
	bsdconv_cmd_zip := exec.Command("curl", "-L", bsdconv_zip_url, "-o", "package.zip")
	err = bsdconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bsdconv_bin_url := "https://github.com/buganini/bsdconv/archive/refs/tags/11.6.bin"
	bsdconv_cmd_bin := exec.Command("curl", "-L", bsdconv_bin_url, "-o", "binary.bin")
	err = bsdconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bsdconv_src_url := "https://github.com/buganini/bsdconv/archive/refs/tags/11.6.src.tar.gz"
	bsdconv_cmd_src := exec.Command("curl", "-L", bsdconv_src_url, "-o", "source.tar.gz")
	err = bsdconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bsdconv_cmd_direct := exec.Command("./binary")
	err = bsdconv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
