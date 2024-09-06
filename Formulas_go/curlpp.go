package main

// Curlpp - C++ wrapper for libcURL
// Homepage: https://www.curlpp.org/

import (
	"fmt"
	
	"os/exec"
)

func installCurlpp() {
	// Método 1: Descargar y extraer .tar.gz
	curlpp_tar_url := "https://github.com/jpbarrette/curlpp/archive/refs/tags/v0.8.1.tar.gz"
	curlpp_cmd_tar := exec.Command("curl", "-L", curlpp_tar_url, "-o", "package.tar.gz")
	err := curlpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curlpp_zip_url := "https://github.com/jpbarrette/curlpp/archive/refs/tags/v0.8.1.zip"
	curlpp_cmd_zip := exec.Command("curl", "-L", curlpp_zip_url, "-o", "package.zip")
	err = curlpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curlpp_bin_url := "https://github.com/jpbarrette/curlpp/archive/refs/tags/v0.8.1.bin"
	curlpp_cmd_bin := exec.Command("curl", "-L", curlpp_bin_url, "-o", "binary.bin")
	err = curlpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curlpp_src_url := "https://github.com/jpbarrette/curlpp/archive/refs/tags/v0.8.1.src.tar.gz"
	curlpp_cmd_src := exec.Command("curl", "-L", curlpp_src_url, "-o", "source.tar.gz")
	err = curlpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curlpp_cmd_direct := exec.Command("./binary")
	err = curlpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
