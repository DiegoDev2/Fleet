package main

// VulkanExtensionlayer - Layer providing Vulkan features when native support is unavailable
// Homepage: https://github.com/KhronosGroup/Vulkan-ExtensionLayer

import (
	"fmt"
	
	"os/exec"
)

func installVulkanExtensionlayer() {
	// Método 1: Descargar y extraer .tar.gz
	vulkanextensionlayer_tar_url := "https://github.com/KhronosGroup/Vulkan-ExtensionLayer/archive/refs/tags/v1.3.295.tar.gz"
	vulkanextensionlayer_cmd_tar := exec.Command("curl", "-L", vulkanextensionlayer_tar_url, "-o", "package.tar.gz")
	err := vulkanextensionlayer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulkanextensionlayer_zip_url := "https://github.com/KhronosGroup/Vulkan-ExtensionLayer/archive/refs/tags/v1.3.295.zip"
	vulkanextensionlayer_cmd_zip := exec.Command("curl", "-L", vulkanextensionlayer_zip_url, "-o", "package.zip")
	err = vulkanextensionlayer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulkanextensionlayer_bin_url := "https://github.com/KhronosGroup/Vulkan-ExtensionLayer/archive/refs/tags/v1.3.295.bin"
	vulkanextensionlayer_cmd_bin := exec.Command("curl", "-L", vulkanextensionlayer_bin_url, "-o", "binary.bin")
	err = vulkanextensionlayer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulkanextensionlayer_src_url := "https://github.com/KhronosGroup/Vulkan-ExtensionLayer/archive/refs/tags/v1.3.295.src.tar.gz"
	vulkanextensionlayer_cmd_src := exec.Command("curl", "-L", vulkanextensionlayer_src_url, "-o", "source.tar.gz")
	err = vulkanextensionlayer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulkanextensionlayer_cmd_direct := exec.Command("./binary")
	err = vulkanextensionlayer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vulkan-loader")
exec.Command("latte", "install", "vulkan-loader")
	fmt.Println("Instalando dependencia: vulkan-tools")
exec.Command("latte", "install", "vulkan-tools")
	fmt.Println("Instalando dependencia: glslang")
exec.Command("latte", "install", "glslang")
	fmt.Println("Instalando dependencia: spirv-headers")
exec.Command("latte", "install", "spirv-headers")
	fmt.Println("Instalando dependencia: spirv-tools")
exec.Command("latte", "install", "spirv-tools")
	fmt.Println("Instalando dependencia: vulkan-headers")
exec.Command("latte", "install", "vulkan-headers")
	fmt.Println("Instalando dependencia: vulkan-utility-libraries")
exec.Command("latte", "install", "vulkan-utility-libraries")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
}
