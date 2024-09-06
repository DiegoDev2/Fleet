package main

// Pybind11 - Seamless operability between C++11 and Python
// Homepage: https://github.com/pybind/pybind11

import (
	"fmt"
	
	"os/exec"
)

func installPybind11() {
	// Método 1: Descargar y extraer .tar.gz
	pybind11_tar_url := "https://github.com/pybind/pybind11/archive/refs/tags/v2.13.5.tar.gz"
	pybind11_cmd_tar := exec.Command("curl", "-L", pybind11_tar_url, "-o", "package.tar.gz")
	err := pybind11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pybind11_zip_url := "https://github.com/pybind/pybind11/archive/refs/tags/v2.13.5.zip"
	pybind11_cmd_zip := exec.Command("curl", "-L", pybind11_zip_url, "-o", "package.zip")
	err = pybind11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pybind11_bin_url := "https://github.com/pybind/pybind11/archive/refs/tags/v2.13.5.bin"
	pybind11_cmd_bin := exec.Command("curl", "-L", pybind11_bin_url, "-o", "binary.bin")
	err = pybind11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pybind11_src_url := "https://github.com/pybind/pybind11/archive/refs/tags/v2.13.5.src.tar.gz"
	pybind11_cmd_src := exec.Command("curl", "-L", pybind11_src_url, "-o", "source.tar.gz")
	err = pybind11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pybind11_cmd_direct := exec.Command("./binary")
	err = pybind11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
