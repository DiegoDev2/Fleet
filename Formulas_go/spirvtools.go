package main

// SpirvTools - API and commands for processing SPIR-V modules
// Homepage: https://github.com/KhronosGroup/SPIRV-Tools

import (
	"fmt"
	
	"os/exec"
)

func installSpirvTools() {
	// Método 1: Descargar y extraer .tar.gz
	spirvtools_tar_url := "https://github.com/KhronosGroup/SPIRV-Tools/archive/refs/tags/vulkan-sdk-1.3.290.0.tar.gz"
	spirvtools_cmd_tar := exec.Command("curl", "-L", spirvtools_tar_url, "-o", "package.tar.gz")
	err := spirvtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spirvtools_zip_url := "https://github.com/KhronosGroup/SPIRV-Tools/archive/refs/tags/vulkan-sdk-1.3.290.0.zip"
	spirvtools_cmd_zip := exec.Command("curl", "-L", spirvtools_zip_url, "-o", "package.zip")
	err = spirvtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spirvtools_bin_url := "https://github.com/KhronosGroup/SPIRV-Tools/archive/refs/tags/vulkan-sdk-1.3.290.0.bin"
	spirvtools_cmd_bin := exec.Command("curl", "-L", spirvtools_bin_url, "-o", "binary.bin")
	err = spirvtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spirvtools_src_url := "https://github.com/KhronosGroup/SPIRV-Tools/archive/refs/tags/vulkan-sdk-1.3.290.0.src.tar.gz"
	spirvtools_cmd_src := exec.Command("curl", "-L", spirvtools_src_url, "-o", "source.tar.gz")
	err = spirvtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spirvtools_cmd_direct := exec.Command("./binary")
	err = spirvtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
