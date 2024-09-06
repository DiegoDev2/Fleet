package main

// Libxls - Read binary Excel files from C/C++
// Homepage: https://github.com/libxls/libxls

import (
	"fmt"
	
	"os/exec"
)

func installLibxls() {
	// Método 1: Descargar y extraer .tar.gz
	libxls_tar_url := "https://github.com/libxls/libxls/releases/download/v1.6.2/libxls-1.6.2.tar.gz"
	libxls_cmd_tar := exec.Command("curl", "-L", libxls_tar_url, "-o", "package.tar.gz")
	err := libxls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxls_zip_url := "https://github.com/libxls/libxls/releases/download/v1.6.2/libxls-1.6.2.zip"
	libxls_cmd_zip := exec.Command("curl", "-L", libxls_zip_url, "-o", "package.zip")
	err = libxls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxls_bin_url := "https://github.com/libxls/libxls/releases/download/v1.6.2/libxls-1.6.2.bin"
	libxls_cmd_bin := exec.Command("curl", "-L", libxls_bin_url, "-o", "binary.bin")
	err = libxls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxls_src_url := "https://github.com/libxls/libxls/releases/download/v1.6.2/libxls-1.6.2.src.tar.gz"
	libxls_cmd_src := exec.Command("curl", "-L", libxls_src_url, "-o", "source.tar.gz")
	err = libxls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxls_cmd_direct := exec.Command("./binary")
	err = libxls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
