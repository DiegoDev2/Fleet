package main

// Bmake - Portable version of NetBSD make(1)
// Homepage: https://www.crufty.net/help/sjg/bmake.html

import (
	"fmt"
	
	"os/exec"
)

func installBmake() {
	// Método 1: Descargar y extraer .tar.gz
	bmake_tar_url := "https://www.crufty.net/ftp/pub/sjg/bmake-20240808.tar.gz"
	bmake_cmd_tar := exec.Command("curl", "-L", bmake_tar_url, "-o", "package.tar.gz")
	err := bmake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bmake_zip_url := "https://www.crufty.net/ftp/pub/sjg/bmake-20240808.zip"
	bmake_cmd_zip := exec.Command("curl", "-L", bmake_zip_url, "-o", "package.zip")
	err = bmake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bmake_bin_url := "https://www.crufty.net/ftp/pub/sjg/bmake-20240808.bin"
	bmake_cmd_bin := exec.Command("curl", "-L", bmake_bin_url, "-o", "binary.bin")
	err = bmake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bmake_src_url := "https://www.crufty.net/ftp/pub/sjg/bmake-20240808.src.tar.gz"
	bmake_cmd_src := exec.Command("curl", "-L", bmake_src_url, "-o", "source.tar.gz")
	err = bmake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bmake_cmd_direct := exec.Command("./binary")
	err = bmake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
