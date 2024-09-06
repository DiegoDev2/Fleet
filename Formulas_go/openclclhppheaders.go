package main

// OpenclClhppHeaders - C++ language header files for the OpenCL API
// Homepage: https://www.khronos.org/registry/OpenCL/

import (
	"fmt"
	
	"os/exec"
)

func installOpenclClhppHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	openclclhppheaders_tar_url := "https://github.com/KhronosGroup/OpenCL-CLHPP/archive/refs/tags/v2024.05.08.tar.gz"
	openclclhppheaders_cmd_tar := exec.Command("curl", "-L", openclclhppheaders_tar_url, "-o", "package.tar.gz")
	err := openclclhppheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openclclhppheaders_zip_url := "https://github.com/KhronosGroup/OpenCL-CLHPP/archive/refs/tags/v2024.05.08.zip"
	openclclhppheaders_cmd_zip := exec.Command("curl", "-L", openclclhppheaders_zip_url, "-o", "package.zip")
	err = openclclhppheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openclclhppheaders_bin_url := "https://github.com/KhronosGroup/OpenCL-CLHPP/archive/refs/tags/v2024.05.08.bin"
	openclclhppheaders_cmd_bin := exec.Command("curl", "-L", openclclhppheaders_bin_url, "-o", "binary.bin")
	err = openclclhppheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openclclhppheaders_src_url := "https://github.com/KhronosGroup/OpenCL-CLHPP/archive/refs/tags/v2024.05.08.src.tar.gz"
	openclclhppheaders_cmd_src := exec.Command("curl", "-L", openclclhppheaders_src_url, "-o", "source.tar.gz")
	err = openclclhppheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openclclhppheaders_cmd_direct := exec.Command("./binary")
	err = openclclhppheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
}
