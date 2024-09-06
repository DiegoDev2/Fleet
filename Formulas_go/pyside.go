package main

// Pyside - Official Python bindings for Qt
// Homepage: https://wiki.qt.io/Qt_for_Python

import (
	"fmt"
	
	"os/exec"
)

func installPyside() {
	// Método 1: Descargar y extraer .tar.gz
	pyside_tar_url := "https://download.qt.io/official_releases/QtForPython/pyside6/PySide6-6.7.0-src/pyside-setup-everywhere-src-6.7.0.tar.xz"
	pyside_cmd_tar := exec.Command("curl", "-L", pyside_tar_url, "-o", "package.tar.gz")
	err := pyside_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyside_zip_url := "https://download.qt.io/official_releases/QtForPython/pyside6/PySide6-6.7.0-src/pyside-setup-everywhere-src-6.7.0.tar.xz"
	pyside_cmd_zip := exec.Command("curl", "-L", pyside_zip_url, "-o", "package.zip")
	err = pyside_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyside_bin_url := "https://download.qt.io/official_releases/QtForPython/pyside6/PySide6-6.7.0-src/pyside-setup-everywhere-src-6.7.0.tar.xz"
	pyside_cmd_bin := exec.Command("curl", "-L", pyside_bin_url, "-o", "binary.bin")
	err = pyside_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyside_src_url := "https://download.qt.io/official_releases/QtForPython/pyside6/PySide6-6.7.0-src/pyside-setup-everywhere-src-6.7.0.tar.xz"
	pyside_cmd_src := exec.Command("curl", "-L", pyside_src_url, "-o", "source.tar.gz")
	err = pyside_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyside_cmd_direct := exec.Command("./binary")
	err = pyside_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
