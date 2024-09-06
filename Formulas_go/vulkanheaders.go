package main

// VulkanHeaders - Vulkan Header files and API registry
// Homepage: https://github.com/KhronosGroup/Vulkan-Headers

import (
	"fmt"
	
	"os/exec"
)

func installVulkanHeaders() {
	// Método 1: Descargar y extraer .tar.gz
	vulkanheaders_tar_url := "https://github.com/KhronosGroup/Vulkan-Headers/archive/refs/tags/v1.3.295.tar.gz"
	vulkanheaders_cmd_tar := exec.Command("curl", "-L", vulkanheaders_tar_url, "-o", "package.tar.gz")
	err := vulkanheaders_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulkanheaders_zip_url := "https://github.com/KhronosGroup/Vulkan-Headers/archive/refs/tags/v1.3.295.zip"
	vulkanheaders_cmd_zip := exec.Command("curl", "-L", vulkanheaders_zip_url, "-o", "package.zip")
	err = vulkanheaders_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulkanheaders_bin_url := "https://github.com/KhronosGroup/Vulkan-Headers/archive/refs/tags/v1.3.295.bin"
	vulkanheaders_cmd_bin := exec.Command("curl", "-L", vulkanheaders_bin_url, "-o", "binary.bin")
	err = vulkanheaders_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulkanheaders_src_url := "https://github.com/KhronosGroup/Vulkan-Headers/archive/refs/tags/v1.3.295.src.tar.gz"
	vulkanheaders_cmd_src := exec.Command("curl", "-L", vulkanheaders_src_url, "-o", "source.tar.gz")
	err = vulkanheaders_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulkanheaders_cmd_direct := exec.Command("./binary")
	err = vulkanheaders_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
