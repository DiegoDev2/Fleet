package main

// Openvino - Open Visual Inference And Optimization toolkit for AI inference
// Homepage: https://docs.openvino.ai

import (
	"fmt"
	
	"os/exec"
)

func installOpenvino() {
	// Método 1: Descargar y extraer .tar.gz
	openvino_tar_url := "https://github.com/openvinotoolkit/openvino/archive/refs/tags/2024.3.0.tar.gz"
	openvino_cmd_tar := exec.Command("curl", "-L", openvino_tar_url, "-o", "package.tar.gz")
	err := openvino_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openvino_zip_url := "https://github.com/openvinotoolkit/openvino/archive/refs/tags/2024.3.0.zip"
	openvino_cmd_zip := exec.Command("curl", "-L", openvino_zip_url, "-o", "package.zip")
	err = openvino_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openvino_bin_url := "https://github.com/openvinotoolkit/openvino/archive/refs/tags/2024.3.0.bin"
	openvino_cmd_bin := exec.Command("curl", "-L", openvino_bin_url, "-o", "binary.bin")
	err = openvino_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openvino_src_url := "https://github.com/openvinotoolkit/openvino/archive/refs/tags/2024.3.0.src.tar.gz"
	openvino_cmd_src := exec.Command("curl", "-L", openvino_src_url, "-o", "source.tar.gz")
	err = openvino_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openvino_cmd_direct := exec.Command("./binary")
	err = openvino_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: flatbuffers")
exec.Command("latte", "install", "flatbuffers")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: protobuf@21")
exec.Command("latte", "install", "protobuf@21")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pugixml")
exec.Command("latte", "install", "pugixml")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: opencl-clhpp-headers")
exec.Command("latte", "install", "opencl-clhpp-headers")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: rapidjson")
exec.Command("latte", "install", "rapidjson")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: scons")
exec.Command("latte", "install", "scons")
	fmt.Println("Instalando dependencia: xbyak")
exec.Command("latte", "install", "xbyak")
}
