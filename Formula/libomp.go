package main

// Libomp - LLVM's OpenMP runtime library
// Homepage: https://openmp.llvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibomp() {
	// Método 1: Descargar y extraer .tar.gz
	libomp_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/openmp-18.1.8.src.tar.xz"
	libomp_cmd_tar := exec.Command("curl", "-L", libomp_tar_url, "-o", "package.tar.gz")
	err := libomp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libomp_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/openmp-18.1.8.src.tar.xz"
	libomp_cmd_zip := exec.Command("curl", "-L", libomp_zip_url, "-o", "package.zip")
	err = libomp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libomp_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/openmp-18.1.8.src.tar.xz"
	libomp_cmd_bin := exec.Command("curl", "-L", libomp_bin_url, "-o", "binary.bin")
	err = libomp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libomp_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/openmp-18.1.8.src.tar.xz"
	libomp_cmd_src := exec.Command("curl", "-L", libomp_src_url, "-o", "source.tar.gz")
	err = libomp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libomp_cmd_direct := exec.Command("./binary")
	err = libomp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: lit")
	exec.Command("latte", "install", "lit").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
