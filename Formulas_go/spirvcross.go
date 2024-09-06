package main

// SpirvCross - Performing reflection and disassembling SPIR-V
// Homepage: https://github.com/KhronosGroup/SPIRV-Cross

import (
	"fmt"
	
	"os/exec"
)

func installSpirvCross() {
	// Método 1: Descargar y extraer .tar.gz
	spirvcross_tar_url := "https://github.com/KhronosGroup/SPIRV-Cross/archive/refs/tags/vulkan-sdk-1.3.290.0.tar.gz"
	spirvcross_cmd_tar := exec.Command("curl", "-L", spirvcross_tar_url, "-o", "package.tar.gz")
	err := spirvcross_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spirvcross_zip_url := "https://github.com/KhronosGroup/SPIRV-Cross/archive/refs/tags/vulkan-sdk-1.3.290.0.zip"
	spirvcross_cmd_zip := exec.Command("curl", "-L", spirvcross_zip_url, "-o", "package.zip")
	err = spirvcross_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spirvcross_bin_url := "https://github.com/KhronosGroup/SPIRV-Cross/archive/refs/tags/vulkan-sdk-1.3.290.0.bin"
	spirvcross_cmd_bin := exec.Command("curl", "-L", spirvcross_bin_url, "-o", "binary.bin")
	err = spirvcross_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spirvcross_src_url := "https://github.com/KhronosGroup/SPIRV-Cross/archive/refs/tags/vulkan-sdk-1.3.290.0.src.tar.gz"
	spirvcross_cmd_src := exec.Command("curl", "-L", spirvcross_src_url, "-o", "source.tar.gz")
	err = spirvcross_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spirvcross_cmd_direct := exec.Command("./binary")
	err = spirvcross_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: glm")
exec.Command("latte", "install", "glm")
	fmt.Println("Instalando dependencia: glslang")
exec.Command("latte", "install", "glslang")
}
