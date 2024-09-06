package main

// Pytorch - Tensors and dynamic neural networks
// Homepage: https://pytorch.org/

import (
	"fmt"
	
	"os/exec"
)

func installPytorch() {
	// Método 1: Descargar y extraer .tar.gz
	pytorch_tar_url := "https://github.com/pytorch/pytorch/releases/download/v2.2.0/pytorch-v2.2.0.tar.gz"
	pytorch_cmd_tar := exec.Command("curl", "-L", pytorch_tar_url, "-o", "package.tar.gz")
	err := pytorch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pytorch_zip_url := "https://github.com/pytorch/pytorch/releases/download/v2.2.0/pytorch-v2.2.0.zip"
	pytorch_cmd_zip := exec.Command("curl", "-L", pytorch_zip_url, "-o", "package.zip")
	err = pytorch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pytorch_bin_url := "https://github.com/pytorch/pytorch/releases/download/v2.2.0/pytorch-v2.2.0.bin"
	pytorch_cmd_bin := exec.Command("curl", "-L", pytorch_bin_url, "-o", "binary.bin")
	err = pytorch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pytorch_src_url := "https://github.com/pytorch/pytorch/releases/download/v2.2.0/pytorch-v2.2.0.src.tar.gz"
	pytorch_cmd_src := exec.Command("curl", "-L", pytorch_src_url, "-o", "source.tar.gz")
	err = pytorch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pytorch_cmd_direct := exec.Command("./binary")
	err = pytorch_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: sleef")
exec.Command("latte", "install", "sleef")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
