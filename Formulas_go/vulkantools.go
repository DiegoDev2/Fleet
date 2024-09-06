package main

// VulkanTools - Vulkan utilities and tools
// Homepage: https://github.com/KhronosGroup/Vulkan-Tools

import (
	"fmt"
	
	"os/exec"
)

func installVulkanTools() {
	// Método 1: Descargar y extraer .tar.gz
	vulkantools_tar_url := "https://github.com/KhronosGroup/Vulkan-Tools/archive/refs/tags/v1.3.295.tar.gz"
	vulkantools_cmd_tar := exec.Command("curl", "-L", vulkantools_tar_url, "-o", "package.tar.gz")
	err := vulkantools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulkantools_zip_url := "https://github.com/KhronosGroup/Vulkan-Tools/archive/refs/tags/v1.3.295.zip"
	vulkantools_cmd_zip := exec.Command("curl", "-L", vulkantools_zip_url, "-o", "package.zip")
	err = vulkantools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulkantools_bin_url := "https://github.com/KhronosGroup/Vulkan-Tools/archive/refs/tags/v1.3.295.bin"
	vulkantools_cmd_bin := exec.Command("curl", "-L", vulkantools_bin_url, "-o", "binary.bin")
	err = vulkantools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulkantools_src_url := "https://github.com/KhronosGroup/Vulkan-Tools/archive/refs/tags/v1.3.295.src.tar.gz"
	vulkantools_cmd_src := exec.Command("curl", "-L", vulkantools_src_url, "-o", "source.tar.gz")
	err = vulkantools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulkantools_cmd_direct := exec.Command("./binary")
	err = vulkantools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vulkan-volk")
exec.Command("latte", "install", "vulkan-volk")
	fmt.Println("Instalando dependencia: glslang")
exec.Command("latte", "install", "glslang")
	fmt.Println("Instalando dependencia: vulkan-headers")
exec.Command("latte", "install", "vulkan-headers")
	fmt.Println("Instalando dependencia: vulkan-loader")
exec.Command("latte", "install", "vulkan-loader")
	fmt.Println("Instalando dependencia: molten-vk")
exec.Command("latte", "install", "molten-vk")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: libxkbfile")
exec.Command("latte", "install", "libxkbfile")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
	fmt.Println("Instalando dependencia: wayland-protocols")
exec.Command("latte", "install", "wayland-protocols")
}
