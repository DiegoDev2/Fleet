package main

// LibunwindHeaders - C API for determining the call-chain of a program
// Homepage: https://opensource.apple.com/

import (
	"fmt"
	
	"os/exec"
)

func installLibunwindHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	libunwindheaders_tar_url := "https://github.com/apple-oss-distributions/libunwind/archive/refs/tags/libunwind-201.tar.gz"
	libunwindheaders_cmd_tar := exec.Command("curl", "-L", libunwindheaders_tar_url, "-o", "package.tar.gz")
	err := libunwindheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libunwindheaders_zip_url := "https://github.com/apple-oss-distributions/libunwind/archive/refs/tags/libunwind-201.zip"
	libunwindheaders_cmd_zip := exec.Command("curl", "-L", libunwindheaders_zip_url, "-o", "package.zip")
	err = libunwindheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libunwindheaders_bin_url := "https://github.com/apple-oss-distributions/libunwind/archive/refs/tags/libunwind-201.bin"
	libunwindheaders_cmd_bin := exec.Command("curl", "-L", libunwindheaders_bin_url, "-o", "binary.bin")
	err = libunwindheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libunwindheaders_src_url := "https://github.com/apple-oss-distributions/libunwind/archive/refs/tags/libunwind-201.src.tar.gz"
	libunwindheaders_cmd_src := exec.Command("curl", "-L", libunwindheaders_src_url, "-o", "source.tar.gz")
	err = libunwindheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libunwindheaders_cmd_direct := exec.Command("./binary")
	err = libunwindheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
