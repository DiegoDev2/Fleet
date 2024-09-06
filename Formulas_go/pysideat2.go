package main

// PysideAT2 - Official Python bindings for Qt
// Homepage: https://wiki.qt.io/Qt_for_Python

import (
	"fmt"
	
	"os/exec"
)

func installPysideAT2() {
	// Método 1: Descargar y extraer .tar.gz
	pysideat2_tar_url := "https://download.qt.io/official_releases/QtForPython/pyside2/PySide2-5.15.15-src/pyside-setup-opensource-src-5.15.15.tar.xz"
	pysideat2_cmd_tar := exec.Command("curl", "-L", pysideat2_tar_url, "-o", "package.tar.gz")
	err := pysideat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pysideat2_zip_url := "https://download.qt.io/official_releases/QtForPython/pyside2/PySide2-5.15.15-src/pyside-setup-opensource-src-5.15.15.tar.xz"
	pysideat2_cmd_zip := exec.Command("curl", "-L", pysideat2_zip_url, "-o", "package.zip")
	err = pysideat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pysideat2_bin_url := "https://download.qt.io/official_releases/QtForPython/pyside2/PySide2-5.15.15-src/pyside-setup-opensource-src-5.15.15.tar.xz"
	pysideat2_cmd_bin := exec.Command("curl", "-L", pysideat2_bin_url, "-o", "binary.bin")
	err = pysideat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pysideat2_src_url := "https://download.qt.io/official_releases/QtForPython/pyside2/PySide2-5.15.15-src/pyside-setup-opensource-src-5.15.15.tar.xz"
	pysideat2_cmd_src := exec.Command("curl", "-L", pysideat2_src_url, "-o", "source.tar.gz")
	err = pysideat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pysideat2_cmd_direct := exec.Command("./binary")
	err = pysideat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm@15")
exec.Command("latte", "install", "llvm@15")
	fmt.Println("Instalando dependencia: python@3.10")
exec.Command("latte", "install", "python@3.10")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
