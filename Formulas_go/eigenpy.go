package main

// Eigenpy - Python bindings of Eigen library with Numpy support
// Homepage: https://github.com/stack-of-tasks/eigenpy

import (
	"fmt"
	
	"os/exec"
)

func installEigenpy() {
	// Método 1: Descargar y extraer .tar.gz
	eigenpy_tar_url := "https://github.com/stack-of-tasks/eigenpy/releases/download/v3.9.0/eigenpy-3.9.0.tar.gz"
	eigenpy_cmd_tar := exec.Command("curl", "-L", eigenpy_tar_url, "-o", "package.tar.gz")
	err := eigenpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eigenpy_zip_url := "https://github.com/stack-of-tasks/eigenpy/releases/download/v3.9.0/eigenpy-3.9.0.zip"
	eigenpy_cmd_zip := exec.Command("curl", "-L", eigenpy_zip_url, "-o", "package.zip")
	err = eigenpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eigenpy_bin_url := "https://github.com/stack-of-tasks/eigenpy/releases/download/v3.9.0/eigenpy-3.9.0.bin"
	eigenpy_cmd_bin := exec.Command("curl", "-L", eigenpy_bin_url, "-o", "binary.bin")
	err = eigenpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eigenpy_src_url := "https://github.com/stack-of-tasks/eigenpy/releases/download/v3.9.0/eigenpy-3.9.0.src.tar.gz"
	eigenpy_cmd_src := exec.Command("curl", "-L", eigenpy_src_url, "-o", "source.tar.gz")
	err = eigenpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eigenpy_cmd_direct := exec.Command("./binary")
	err = eigenpy_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: boost-python3")
exec.Command("latte", "install", "boost-python3")
	fmt.Println("Instalando dependencia: eigen")
exec.Command("latte", "install", "eigen")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: scipy")
exec.Command("latte", "install", "scipy")
}
