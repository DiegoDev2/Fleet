package main

// Arrayfire - General purpose GPU library
// Homepage: https://arrayfire.com

import (
	"fmt"
	
	"os/exec"
)

func installArrayfire() {
	// Método 1: Descargar y extraer .tar.gz
	arrayfire_tar_url := "https://github.com/arrayfire/arrayfire/releases/download/v3.9.0/arrayfire-full-3.9.0.tar.bz2"
	arrayfire_cmd_tar := exec.Command("curl", "-L", arrayfire_tar_url, "-o", "package.tar.gz")
	err := arrayfire_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arrayfire_zip_url := "https://github.com/arrayfire/arrayfire/releases/download/v3.9.0/arrayfire-full-3.9.0.tar.bz2"
	arrayfire_cmd_zip := exec.Command("curl", "-L", arrayfire_zip_url, "-o", "package.zip")
	err = arrayfire_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arrayfire_bin_url := "https://github.com/arrayfire/arrayfire/releases/download/v3.9.0/arrayfire-full-3.9.0.tar.bz2"
	arrayfire_cmd_bin := exec.Command("curl", "-L", arrayfire_bin_url, "-o", "binary.bin")
	err = arrayfire_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arrayfire_src_url := "https://github.com/arrayfire/arrayfire/releases/download/v3.9.0/arrayfire-full-3.9.0.tar.bz2"
	arrayfire_cmd_src := exec.Command("curl", "-L", arrayfire_src_url, "-o", "source.tar.gz")
	err = arrayfire_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arrayfire_cmd_direct := exec.Command("./binary")
	err = arrayfire_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: freeimage")
exec.Command("latte", "install", "freeimage")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: spdlog")
exec.Command("latte", "install", "spdlog")
	fmt.Println("Instalando dependencia: opencl-headers")
exec.Command("latte", "install", "opencl-headers")
	fmt.Println("Instalando dependencia: opencl-icd-loader")
exec.Command("latte", "install", "opencl-icd-loader")
	fmt.Println("Instalando dependencia: pocl")
exec.Command("latte", "install", "pocl")
}
