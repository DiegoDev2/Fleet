package main

// Torchvision - Datasets, transforms, and models for computer vision
// Homepage: https://github.com/pytorch/vision

import (
	"fmt"
	
	"os/exec"
)

func installTorchvision() {
	// Método 1: Descargar y extraer .tar.gz
	torchvision_tar_url := "https://github.com/pytorch/vision/archive/refs/tags/v0.17.0.tar.gz"
	torchvision_cmd_tar := exec.Command("curl", "-L", torchvision_tar_url, "-o", "package.tar.gz")
	err := torchvision_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	torchvision_zip_url := "https://github.com/pytorch/vision/archive/refs/tags/v0.17.0.zip"
	torchvision_cmd_zip := exec.Command("curl", "-L", torchvision_zip_url, "-o", "package.zip")
	err = torchvision_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	torchvision_bin_url := "https://github.com/pytorch/vision/archive/refs/tags/v0.17.0.bin"
	torchvision_cmd_bin := exec.Command("curl", "-L", torchvision_bin_url, "-o", "binary.bin")
	err = torchvision_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	torchvision_src_url := "https://github.com/pytorch/vision/archive/refs/tags/v0.17.0.src.tar.gz"
	torchvision_cmd_src := exec.Command("curl", "-L", torchvision_src_url, "-o", "source.tar.gz")
	err = torchvision_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	torchvision_cmd_direct := exec.Command("./binary")
	err = torchvision_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: pytorch")
exec.Command("latte", "install", "pytorch")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
