package main

// SpirvHeaders - Headers for SPIR-V
// Homepage: https://github.com/KhronosGroup/SPIRV-Headers

import (
	"fmt"
	
	"os/exec"
)

func installSpirvHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	spirvheaders_tar_url := "https://github.com/KhronosGroup/SPIRV-Headers/archive/refs/tags/vulkan-sdk-1.3.290.0.tar.gz"
	spirvheaders_cmd_tar := exec.Command("curl", "-L", spirvheaders_tar_url, "-o", "package.tar.gz")
	err := spirvheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spirvheaders_zip_url := "https://github.com/KhronosGroup/SPIRV-Headers/archive/refs/tags/vulkan-sdk-1.3.290.0.zip"
	spirvheaders_cmd_zip := exec.Command("curl", "-L", spirvheaders_zip_url, "-o", "package.zip")
	err = spirvheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spirvheaders_bin_url := "https://github.com/KhronosGroup/SPIRV-Headers/archive/refs/tags/vulkan-sdk-1.3.290.0.bin"
	spirvheaders_cmd_bin := exec.Command("curl", "-L", spirvheaders_bin_url, "-o", "binary.bin")
	err = spirvheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spirvheaders_src_url := "https://github.com/KhronosGroup/SPIRV-Headers/archive/refs/tags/vulkan-sdk-1.3.290.0.src.tar.gz"
	spirvheaders_cmd_src := exec.Command("curl", "-L", spirvheaders_src_url, "-o", "source.tar.gz")
	err = spirvheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spirvheaders_cmd_direct := exec.Command("./binary")
	err = spirvheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
