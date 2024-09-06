package main

// OpenclHeaders - C language header files for the OpenCL API
// Homepage: https://www.khronos.org/registry/OpenCL/

import (
	"fmt"
	
	"os/exec"
)

func installOpenclHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	openclheaders_tar_url := "https://github.com/KhronosGroup/OpenCL-Headers/archive/refs/tags/v2024.05.08.tar.gz"
	openclheaders_cmd_tar := exec.Command("curl", "-L", openclheaders_tar_url, "-o", "package.tar.gz")
	err := openclheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openclheaders_zip_url := "https://github.com/KhronosGroup/OpenCL-Headers/archive/refs/tags/v2024.05.08.zip"
	openclheaders_cmd_zip := exec.Command("curl", "-L", openclheaders_zip_url, "-o", "package.zip")
	err = openclheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openclheaders_bin_url := "https://github.com/KhronosGroup/OpenCL-Headers/archive/refs/tags/v2024.05.08.bin"
	openclheaders_cmd_bin := exec.Command("curl", "-L", openclheaders_bin_url, "-o", "binary.bin")
	err = openclheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openclheaders_src_url := "https://github.com/KhronosGroup/OpenCL-Headers/archive/refs/tags/v2024.05.08.src.tar.gz"
	openclheaders_cmd_src := exec.Command("curl", "-L", openclheaders_src_url, "-o", "source.tar.gz")
	err = openclheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openclheaders_cmd_direct := exec.Command("./binary")
	err = openclheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
