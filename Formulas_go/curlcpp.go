package main

// Curlcpp - Object oriented C++ wrapper for CURL (libcurl)
// Homepage: https://josephp91.github.io/curlcpp

import (
	"fmt"
	
	"os/exec"
)

func installCurlcpp() {
	// Método 1: Descargar y extraer .tar.gz
	curlcpp_tar_url := "https://github.com/JosephP91/curlcpp/archive/refs/tags/3.1.tar.gz"
	curlcpp_cmd_tar := exec.Command("curl", "-L", curlcpp_tar_url, "-o", "package.tar.gz")
	err := curlcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	curlcpp_zip_url := "https://github.com/JosephP91/curlcpp/archive/refs/tags/3.1.zip"
	curlcpp_cmd_zip := exec.Command("curl", "-L", curlcpp_zip_url, "-o", "package.zip")
	err = curlcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	curlcpp_bin_url := "https://github.com/JosephP91/curlcpp/archive/refs/tags/3.1.bin"
	curlcpp_cmd_bin := exec.Command("curl", "-L", curlcpp_bin_url, "-o", "binary.bin")
	err = curlcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	curlcpp_src_url := "https://github.com/JosephP91/curlcpp/archive/refs/tags/3.1.src.tar.gz"
	curlcpp_cmd_src := exec.Command("curl", "-L", curlcpp_src_url, "-o", "source.tar.gz")
	err = curlcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	curlcpp_cmd_direct := exec.Command("./binary")
	err = curlcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
