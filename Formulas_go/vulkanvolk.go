package main

// VulkanVolk - Meta loader for Vulkan API
// Homepage: https://github.com/zeux/volk

import (
	"fmt"
	
	"os/exec"
)

func installVulkanVolk() {
	// Método 1: Descargar y extraer .tar.gz
	vulkanvolk_tar_url := "https://github.com/zeux/volk/archive/refs/tags/vulkan-sdk-1.3.290.0.tar.gz"
	vulkanvolk_cmd_tar := exec.Command("curl", "-L", vulkanvolk_tar_url, "-o", "package.tar.gz")
	err := vulkanvolk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulkanvolk_zip_url := "https://github.com/zeux/volk/archive/refs/tags/vulkan-sdk-1.3.290.0.zip"
	vulkanvolk_cmd_zip := exec.Command("curl", "-L", vulkanvolk_zip_url, "-o", "package.zip")
	err = vulkanvolk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulkanvolk_bin_url := "https://github.com/zeux/volk/archive/refs/tags/vulkan-sdk-1.3.290.0.bin"
	vulkanvolk_cmd_bin := exec.Command("curl", "-L", vulkanvolk_bin_url, "-o", "binary.bin")
	err = vulkanvolk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulkanvolk_src_url := "https://github.com/zeux/volk/archive/refs/tags/vulkan-sdk-1.3.290.0.src.tar.gz"
	vulkanvolk_cmd_src := exec.Command("curl", "-L", vulkanvolk_src_url, "-o", "source.tar.gz")
	err = vulkanvolk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulkanvolk_cmd_direct := exec.Command("./binary")
	err = vulkanvolk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: vulkan-headers")
exec.Command("latte", "install", "vulkan-headers")
	fmt.Println("Instalando dependencia: vulkan-loader")
exec.Command("latte", "install", "vulkan-loader")
}
