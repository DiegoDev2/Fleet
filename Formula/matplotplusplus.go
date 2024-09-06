package main

// Matplotplusplus - C++ Graphics Library for Data Visualization
// Homepage: https://github.com/alandefreitas/matplotplusplus

import (
	"fmt"
	
	"os/exec"
)

func installMatplotplusplus() {
	// Método 1: Descargar y extraer .tar.gz
	matplotplusplus_tar_url := "https://github.com/alandefreitas/matplotplusplus/archive/refs/tags/v1.2.1.tar.gz"
	matplotplusplus_cmd_tar := exec.Command("curl", "-L", matplotplusplus_tar_url, "-o", "package.tar.gz")
	err := matplotplusplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	matplotplusplus_zip_url := "https://github.com/alandefreitas/matplotplusplus/archive/refs/tags/v1.2.1.zip"
	matplotplusplus_cmd_zip := exec.Command("curl", "-L", matplotplusplus_zip_url, "-o", "package.zip")
	err = matplotplusplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	matplotplusplus_bin_url := "https://github.com/alandefreitas/matplotplusplus/archive/refs/tags/v1.2.1.bin"
	matplotplusplus_cmd_bin := exec.Command("curl", "-L", matplotplusplus_bin_url, "-o", "binary.bin")
	err = matplotplusplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	matplotplusplus_src_url := "https://github.com/alandefreitas/matplotplusplus/archive/refs/tags/v1.2.1.src.tar.gz"
	matplotplusplus_cmd_src := exec.Command("curl", "-L", matplotplusplus_src_url, "-o", "source.tar.gz")
	err = matplotplusplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	matplotplusplus_cmd_direct := exec.Command("./binary")
	err = matplotplusplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gnuplot")
	exec.Command("latte", "install", "gnuplot").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: openexr")
	exec.Command("latte", "install", "openexr").Run()
}
