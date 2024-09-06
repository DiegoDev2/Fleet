package main

// VulkanProfiles - Tools for Vulkan profiles
// Homepage: https://github.com/KhronosGroup/Vulkan-Profiles

import (
	"fmt"
	
	"os/exec"
)

func installVulkanProfiles() {
	// Método 1: Descargar y extraer .tar.gz
	vulkanprofiles_tar_url := "https://github.com/KhronosGroup/Vulkan-Profiles/archive/refs/tags/v1.3.295.tar.gz"
	vulkanprofiles_cmd_tar := exec.Command("curl", "-L", vulkanprofiles_tar_url, "-o", "package.tar.gz")
	err := vulkanprofiles_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulkanprofiles_zip_url := "https://github.com/KhronosGroup/Vulkan-Profiles/archive/refs/tags/v1.3.295.zip"
	vulkanprofiles_cmd_zip := exec.Command("curl", "-L", vulkanprofiles_zip_url, "-o", "package.zip")
	err = vulkanprofiles_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulkanprofiles_bin_url := "https://github.com/KhronosGroup/Vulkan-Profiles/archive/refs/tags/v1.3.295.bin"
	vulkanprofiles_cmd_bin := exec.Command("curl", "-L", vulkanprofiles_bin_url, "-o", "binary.bin")
	err = vulkanprofiles_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulkanprofiles_src_url := "https://github.com/KhronosGroup/Vulkan-Profiles/archive/refs/tags/v1.3.295.src.tar.gz"
	vulkanprofiles_cmd_src := exec.Command("curl", "-L", vulkanprofiles_src_url, "-o", "source.tar.gz")
	err = vulkanprofiles_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulkanprofiles_cmd_direct := exec.Command("./binary")
	err = vulkanprofiles_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: vulkan-tools")
	exec.Command("latte", "install", "vulkan-tools").Run()
	fmt.Println("Instalando dependencia: jsoncpp")
	exec.Command("latte", "install", "jsoncpp").Run()
	fmt.Println("Instalando dependencia: valijson")
	exec.Command("latte", "install", "valijson").Run()
	fmt.Println("Instalando dependencia: vulkan-headers")
	exec.Command("latte", "install", "vulkan-headers").Run()
	fmt.Println("Instalando dependencia: vulkan-loader")
	exec.Command("latte", "install", "vulkan-loader").Run()
	fmt.Println("Instalando dependencia: vulkan-utility-libraries")
	exec.Command("latte", "install", "vulkan-utility-libraries").Run()
	fmt.Println("Instalando dependencia: molten-vk")
	exec.Command("latte", "install", "molten-vk").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
