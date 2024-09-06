package main

// OpenclIcdLoader - OpenCL Installable Client Driver (ICD) Loader
// Homepage: https://www.khronos.org/registry/OpenCL/

import (
	"fmt"
	
	"os/exec"
)

func installOpenclIcdLoader() {
	// Método 1: Descargar y extraer .tar.gz
	openclicdloader_tar_url := "https://github.com/KhronosGroup/OpenCL-ICD-Loader/archive/refs/tags/v2024.05.08.tar.gz"
	openclicdloader_cmd_tar := exec.Command("curl", "-L", openclicdloader_tar_url, "-o", "package.tar.gz")
	err := openclicdloader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openclicdloader_zip_url := "https://github.com/KhronosGroup/OpenCL-ICD-Loader/archive/refs/tags/v2024.05.08.zip"
	openclicdloader_cmd_zip := exec.Command("curl", "-L", openclicdloader_zip_url, "-o", "package.zip")
	err = openclicdloader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openclicdloader_bin_url := "https://github.com/KhronosGroup/OpenCL-ICD-Loader/archive/refs/tags/v2024.05.08.bin"
	openclicdloader_cmd_bin := exec.Command("curl", "-L", openclicdloader_bin_url, "-o", "binary.bin")
	err = openclicdloader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openclicdloader_src_url := "https://github.com/KhronosGroup/OpenCL-ICD-Loader/archive/refs/tags/v2024.05.08.src.tar.gz"
	openclicdloader_cmd_src := exec.Command("curl", "-L", openclicdloader_src_url, "-o", "source.tar.gz")
	err = openclicdloader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openclicdloader_cmd_direct := exec.Command("./binary")
	err = openclicdloader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: opencl-headers")
	exec.Command("latte", "install", "opencl-headers").Run()
}
