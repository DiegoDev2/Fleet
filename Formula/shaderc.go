package main

// Shaderc - Collection of tools, libraries, and tests for Vulkan shader compilation
// Homepage: https://github.com/google/shaderc

import (
	"fmt"
	
	"os/exec"
)

func installShaderc() {
	// Método 1: Descargar y extraer .tar.gz
	shaderc_tar_url := "https://github.com/google/shaderc/archive/refs/tags/v2024.1.tar.gz"
	shaderc_cmd_tar := exec.Command("curl", "-L", shaderc_tar_url, "-o", "package.tar.gz")
	err := shaderc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shaderc_zip_url := "https://github.com/google/shaderc/archive/refs/tags/v2024.1.zip"
	shaderc_cmd_zip := exec.Command("curl", "-L", shaderc_zip_url, "-o", "package.zip")
	err = shaderc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shaderc_bin_url := "https://github.com/google/shaderc/archive/refs/tags/v2024.1.bin"
	shaderc_cmd_bin := exec.Command("curl", "-L", shaderc_bin_url, "-o", "binary.bin")
	err = shaderc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shaderc_src_url := "https://github.com/google/shaderc/archive/refs/tags/v2024.1.src.tar.gz"
	shaderc_cmd_src := exec.Command("curl", "-L", shaderc_src_url, "-o", "source.tar.gz")
	err = shaderc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shaderc_cmd_direct := exec.Command("./binary")
	err = shaderc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
