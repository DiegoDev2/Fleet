package main

// Py3cairo - Python 3 bindings for the Cairo graphics library
// Homepage: https://cairographics.org/pycairo/

import (
	"fmt"
	
	"os/exec"
)

func installPy3cairo() {
	// Método 1: Descargar y extraer .tar.gz
	py3cairo_tar_url := "https://github.com/pygobject/pycairo/releases/download/v1.26.1/pycairo-1.26.1.tar.gz"
	py3cairo_cmd_tar := exec.Command("curl", "-L", py3cairo_tar_url, "-o", "package.tar.gz")
	err := py3cairo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	py3cairo_zip_url := "https://github.com/pygobject/pycairo/releases/download/v1.26.1/pycairo-1.26.1.zip"
	py3cairo_cmd_zip := exec.Command("curl", "-L", py3cairo_zip_url, "-o", "package.zip")
	err = py3cairo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	py3cairo_bin_url := "https://github.com/pygobject/pycairo/releases/download/v1.26.1/pycairo-1.26.1.bin"
	py3cairo_cmd_bin := exec.Command("curl", "-L", py3cairo_bin_url, "-o", "binary.bin")
	err = py3cairo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	py3cairo_src_url := "https://github.com/pygobject/pycairo/releases/download/v1.26.1/pycairo-1.26.1.src.tar.gz"
	py3cairo_cmd_src := exec.Command("curl", "-L", py3cairo_src_url, "-o", "source.tar.gz")
	err = py3cairo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	py3cairo_cmd_direct := exec.Command("./binary")
	err = py3cairo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
}
