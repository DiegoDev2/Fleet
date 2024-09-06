package main

// Tbb - Rich and complete approach to parallelism in C++
// Homepage: https://github.com/oneapi-src/oneTBB

import (
	"fmt"
	
	"os/exec"
)

func installTbb() {
	// Método 1: Descargar y extraer .tar.gz
	tbb_tar_url := "https://github.com/oneapi-src/oneTBB/archive/refs/tags/v2021.13.0.tar.gz"
	tbb_cmd_tar := exec.Command("curl", "-L", tbb_tar_url, "-o", "package.tar.gz")
	err := tbb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tbb_zip_url := "https://github.com/oneapi-src/oneTBB/archive/refs/tags/v2021.13.0.zip"
	tbb_cmd_zip := exec.Command("curl", "-L", tbb_zip_url, "-o", "package.zip")
	err = tbb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tbb_bin_url := "https://github.com/oneapi-src/oneTBB/archive/refs/tags/v2021.13.0.bin"
	tbb_cmd_bin := exec.Command("curl", "-L", tbb_bin_url, "-o", "binary.bin")
	err = tbb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tbb_src_url := "https://github.com/oneapi-src/oneTBB/archive/refs/tags/v2021.13.0.src.tar.gz"
	tbb_cmd_src := exec.Command("curl", "-L", tbb_src_url, "-o", "source.tar.gz")
	err = tbb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tbb_cmd_direct := exec.Command("./binary")
	err = tbb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: hwloc")
	exec.Command("latte", "install", "hwloc").Run()
}
