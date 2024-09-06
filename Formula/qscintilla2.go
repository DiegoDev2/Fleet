package main

// Qscintilla2 - Port to Qt of the Scintilla editing component
// Homepage: https://www.riverbankcomputing.com/software/qscintilla/intro

import (
	"fmt"
	
	"os/exec"
)

func installQscintilla2() {
	// Método 1: Descargar y extraer .tar.gz
	qscintilla2_tar_url := "https://www.riverbankcomputing.com/static/Downloads/QScintilla/2.14.1/QScintilla_src-2.14.1.tar.gz"
	qscintilla2_cmd_tar := exec.Command("curl", "-L", qscintilla2_tar_url, "-o", "package.tar.gz")
	err := qscintilla2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qscintilla2_zip_url := "https://www.riverbankcomputing.com/static/Downloads/QScintilla/2.14.1/QScintilla_src-2.14.1.zip"
	qscintilla2_cmd_zip := exec.Command("curl", "-L", qscintilla2_zip_url, "-o", "package.zip")
	err = qscintilla2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qscintilla2_bin_url := "https://www.riverbankcomputing.com/static/Downloads/QScintilla/2.14.1/QScintilla_src-2.14.1.bin"
	qscintilla2_cmd_bin := exec.Command("curl", "-L", qscintilla2_bin_url, "-o", "binary.bin")
	err = qscintilla2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qscintilla2_src_url := "https://www.riverbankcomputing.com/static/Downloads/QScintilla/2.14.1/QScintilla_src-2.14.1.src.tar.gz"
	qscintilla2_cmd_src := exec.Command("curl", "-L", qscintilla2_src_url, "-o", "source.tar.gz")
	err = qscintilla2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qscintilla2_cmd_direct := exec.Command("./binary")
	err = qscintilla2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyqt-builder")
	exec.Command("latte", "install", "pyqt-builder").Run()
	fmt.Println("Instalando dependencia: pyqt")
	exec.Command("latte", "install", "pyqt").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}
