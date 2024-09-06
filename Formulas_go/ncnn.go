package main

// Ncnn - High-performance neural network inference framework
// Homepage: https://github.com/Tencent/ncnn

import (
	"fmt"
	
	"os/exec"
)

func installNcnn() {
	// Método 1: Descargar y extraer .tar.gz
	ncnn_tar_url := "https://github.com/Tencent/ncnn/archive/refs/tags/20240820.tar.gz"
	ncnn_cmd_tar := exec.Command("curl", "-L", ncnn_tar_url, "-o", "package.tar.gz")
	err := ncnn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncnn_zip_url := "https://github.com/Tencent/ncnn/archive/refs/tags/20240820.zip"
	ncnn_cmd_zip := exec.Command("curl", "-L", ncnn_zip_url, "-o", "package.zip")
	err = ncnn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncnn_bin_url := "https://github.com/Tencent/ncnn/archive/refs/tags/20240820.bin"
	ncnn_cmd_bin := exec.Command("curl", "-L", ncnn_bin_url, "-o", "binary.bin")
	err = ncnn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncnn_src_url := "https://github.com/Tencent/ncnn/archive/refs/tags/20240820.src.tar.gz"
	ncnn_cmd_src := exec.Command("curl", "-L", ncnn_src_url, "-o", "source.tar.gz")
	err = ncnn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncnn_cmd_direct := exec.Command("./binary")
	err = ncnn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: glslang")
exec.Command("latte", "install", "glslang")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: molten-vk")
exec.Command("latte", "install", "molten-vk")
	fmt.Println("Instalando dependencia: spirv-tools")
exec.Command("latte", "install", "spirv-tools")
	fmt.Println("Instalando dependencia: vulkan-tools")
exec.Command("latte", "install", "vulkan-tools")
}
