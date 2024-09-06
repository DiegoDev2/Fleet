package main

// Clfft - FFT functions written in OpenCL
// Homepage: https://github.com/clMathLibraries/clFFT

import (
	"fmt"
	
	"os/exec"
)

func installClfft() {
	// Método 1: Descargar y extraer .tar.gz
	clfft_tar_url := "https://github.com/clMathLibraries/clFFT/archive/refs/tags/v2.12.2.tar.gz"
	clfft_cmd_tar := exec.Command("curl", "-L", clfft_tar_url, "-o", "package.tar.gz")
	err := clfft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clfft_zip_url := "https://github.com/clMathLibraries/clFFT/archive/refs/tags/v2.12.2.zip"
	clfft_cmd_zip := exec.Command("curl", "-L", clfft_zip_url, "-o", "package.zip")
	err = clfft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clfft_bin_url := "https://github.com/clMathLibraries/clFFT/archive/refs/tags/v2.12.2.bin"
	clfft_cmd_bin := exec.Command("curl", "-L", clfft_bin_url, "-o", "binary.bin")
	err = clfft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clfft_src_url := "https://github.com/clMathLibraries/clFFT/archive/refs/tags/v2.12.2.src.tar.gz"
	clfft_cmd_src := exec.Command("curl", "-L", clfft_src_url, "-o", "source.tar.gz")
	err = clfft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clfft_cmd_direct := exec.Command("./binary")
	err = clfft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
