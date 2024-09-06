package main

// VulkanUtilityLibraries - Utility Libraries for Vulkan
// Homepage: https://github.com/KhronosGroup/Vulkan-Utility-Libraries

import (
	"fmt"
	
	"os/exec"
)

func installVulkanUtilityLibraries() {
	// Método 1: Descargar y extraer .tar.gz
	vulkanutilitylibraries_tar_url := "https://github.com/KhronosGroup/Vulkan-Utility-Libraries/archive/refs/tags/v1.3.295.tar.gz"
	vulkanutilitylibraries_cmd_tar := exec.Command("curl", "-L", vulkanutilitylibraries_tar_url, "-o", "package.tar.gz")
	err := vulkanutilitylibraries_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulkanutilitylibraries_zip_url := "https://github.com/KhronosGroup/Vulkan-Utility-Libraries/archive/refs/tags/v1.3.295.zip"
	vulkanutilitylibraries_cmd_zip := exec.Command("curl", "-L", vulkanutilitylibraries_zip_url, "-o", "package.zip")
	err = vulkanutilitylibraries_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulkanutilitylibraries_bin_url := "https://github.com/KhronosGroup/Vulkan-Utility-Libraries/archive/refs/tags/v1.3.295.bin"
	vulkanutilitylibraries_cmd_bin := exec.Command("curl", "-L", vulkanutilitylibraries_bin_url, "-o", "binary.bin")
	err = vulkanutilitylibraries_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulkanutilitylibraries_src_url := "https://github.com/KhronosGroup/Vulkan-Utility-Libraries/archive/refs/tags/v1.3.295.src.tar.gz"
	vulkanutilitylibraries_cmd_src := exec.Command("curl", "-L", vulkanutilitylibraries_src_url, "-o", "source.tar.gz")
	err = vulkanutilitylibraries_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulkanutilitylibraries_cmd_direct := exec.Command("./binary")
	err = vulkanutilitylibraries_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: vulkan-headers")
	exec.Command("latte", "install", "vulkan-headers").Run()
}
