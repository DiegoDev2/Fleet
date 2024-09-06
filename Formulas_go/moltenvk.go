package main

// MoltenVk - Implementation of the Vulkan graphics and compute API on top of Metal
// Homepage: https://github.com/KhronosGroup/MoltenVK

import (
	"fmt"
	
	"os/exec"
)

func installMoltenVk() {
	// Método 1: Descargar y extraer .tar.gz
	moltenvk_tar_url := "https://github.com/KhronosGroup/MoltenVK/archive/refs/tags/v1.2.10.tar.gz"
	moltenvk_cmd_tar := exec.Command("curl", "-L", moltenvk_tar_url, "-o", "package.tar.gz")
	err := moltenvk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moltenvk_zip_url := "https://github.com/KhronosGroup/MoltenVK/archive/refs/tags/v1.2.10.zip"
	moltenvk_cmd_zip := exec.Command("curl", "-L", moltenvk_zip_url, "-o", "package.zip")
	err = moltenvk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moltenvk_bin_url := "https://github.com/KhronosGroup/MoltenVK/archive/refs/tags/v1.2.10.bin"
	moltenvk_cmd_bin := exec.Command("curl", "-L", moltenvk_bin_url, "-o", "binary.bin")
	err = moltenvk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moltenvk_src_url := "https://github.com/KhronosGroup/MoltenVK/archive/refs/tags/v1.2.10.src.tar.gz"
	moltenvk_cmd_src := exec.Command("curl", "-L", moltenvk_src_url, "-o", "source.tar.gz")
	err = moltenvk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moltenvk_cmd_direct := exec.Command("./binary")
	err = moltenvk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
