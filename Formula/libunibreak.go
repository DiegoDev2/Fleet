package main

// Libunibreak - Implementation of the Unicode line- and word-breaking algorithms
// Homepage: https://github.com/adah1972/libunibreak

import (
	"fmt"
	
	"os/exec"
)

func installLibunibreak() {
	// Método 1: Descargar y extraer .tar.gz
	libunibreak_tar_url := "https://github.com/adah1972/libunibreak/releases/download/libunibreak_6_1/libunibreak-6.1.tar.gz"
	libunibreak_cmd_tar := exec.Command("curl", "-L", libunibreak_tar_url, "-o", "package.tar.gz")
	err := libunibreak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libunibreak_zip_url := "https://github.com/adah1972/libunibreak/releases/download/libunibreak_6_1/libunibreak-6.1.zip"
	libunibreak_cmd_zip := exec.Command("curl", "-L", libunibreak_zip_url, "-o", "package.zip")
	err = libunibreak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libunibreak_bin_url := "https://github.com/adah1972/libunibreak/releases/download/libunibreak_6_1/libunibreak-6.1.bin"
	libunibreak_cmd_bin := exec.Command("curl", "-L", libunibreak_bin_url, "-o", "binary.bin")
	err = libunibreak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libunibreak_src_url := "https://github.com/adah1972/libunibreak/releases/download/libunibreak_6_1/libunibreak-6.1.src.tar.gz"
	libunibreak_cmd_src := exec.Command("curl", "-L", libunibreak_src_url, "-o", "source.tar.gz")
	err = libunibreak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libunibreak_cmd_direct := exec.Command("./binary")
	err = libunibreak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
