package main

// Jansson - C library for encoding, decoding, and manipulating JSON
// Homepage: https://digip.org/jansson/

import (
	"fmt"
	
	"os/exec"
)

func installJansson() {
	// Método 1: Descargar y extraer .tar.gz
	jansson_tar_url := "https://github.com/akheron/jansson/releases/download/v2.14/jansson-2.14.tar.gz"
	jansson_cmd_tar := exec.Command("curl", "-L", jansson_tar_url, "-o", "package.tar.gz")
	err := jansson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jansson_zip_url := "https://github.com/akheron/jansson/releases/download/v2.14/jansson-2.14.zip"
	jansson_cmd_zip := exec.Command("curl", "-L", jansson_zip_url, "-o", "package.zip")
	err = jansson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jansson_bin_url := "https://github.com/akheron/jansson/releases/download/v2.14/jansson-2.14.bin"
	jansson_cmd_bin := exec.Command("curl", "-L", jansson_bin_url, "-o", "binary.bin")
	err = jansson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jansson_src_url := "https://github.com/akheron/jansson/releases/download/v2.14/jansson-2.14.src.tar.gz"
	jansson_cmd_src := exec.Command("curl", "-L", jansson_src_url, "-o", "source.tar.gz")
	err = jansson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jansson_cmd_direct := exec.Command("./binary")
	err = jansson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
