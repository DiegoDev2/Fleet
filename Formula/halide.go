package main

// Halide - Language for fast, portable data-parallel computation
// Homepage: https://halide-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installHalide() {
	// Método 1: Descargar y extraer .tar.gz
	halide_tar_url := "https://github.com/halide/Halide/archive/refs/tags/v18.0.0.tar.gz"
	halide_cmd_tar := exec.Command("curl", "-L", halide_tar_url, "-o", "package.tar.gz")
	err := halide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	halide_zip_url := "https://github.com/halide/Halide/archive/refs/tags/v18.0.0.zip"
	halide_cmd_zip := exec.Command("curl", "-L", halide_zip_url, "-o", "package.zip")
	err = halide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	halide_bin_url := "https://github.com/halide/Halide/archive/refs/tags/v18.0.0.bin"
	halide_cmd_bin := exec.Command("curl", "-L", halide_bin_url, "-o", "binary.bin")
	err = halide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	halide_src_url := "https://github.com/halide/Halide/archive/refs/tags/v18.0.0.src.tar.gz"
	halide_cmd_src := exec.Command("curl", "-L", halide_src_url, "-o", "source.tar.gz")
	err = halide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	halide_cmd_direct := exec.Command("./binary")
	err = halide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pybind11")
	exec.Command("latte", "install", "pybind11").Run()
	fmt.Println("Instalando dependencia: flatbuffers")
	exec.Command("latte", "install", "flatbuffers").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
