package main

// Clblas - Library containing BLAS functions written in OpenCL
// Homepage: https://github.com/clMathLibraries/clBLAS

import (
	"fmt"
	
	"os/exec"
)

func installClblas() {
	// Método 1: Descargar y extraer .tar.gz
	clblas_tar_url := "https://github.com/clMathLibraries/clBLAS/archive/refs/tags/v2.12.tar.gz"
	clblas_cmd_tar := exec.Command("curl", "-L", clblas_tar_url, "-o", "package.tar.gz")
	err := clblas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clblas_zip_url := "https://github.com/clMathLibraries/clBLAS/archive/refs/tags/v2.12.zip"
	clblas_cmd_zip := exec.Command("curl", "-L", clblas_zip_url, "-o", "package.zip")
	err = clblas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clblas_bin_url := "https://github.com/clMathLibraries/clBLAS/archive/refs/tags/v2.12.bin"
	clblas_cmd_bin := exec.Command("curl", "-L", clblas_bin_url, "-o", "binary.bin")
	err = clblas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clblas_src_url := "https://github.com/clMathLibraries/clBLAS/archive/refs/tags/v2.12.src.tar.gz"
	clblas_cmd_src := exec.Command("curl", "-L", clblas_src_url, "-o", "source.tar.gz")
	err = clblas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clblas_cmd_direct := exec.Command("./binary")
	err = clblas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
