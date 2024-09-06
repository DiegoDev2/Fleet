package main

// Pyqt - Python bindings for v6 of Qt
// Homepage: https://www.riverbankcomputing.com/software/pyqt/intro

import (
	"fmt"
	
	"os/exec"
)

func installPyqt() {
	// Método 1: Descargar y extraer .tar.gz
	pyqt_tar_url := "https://files.pythonhosted.org/packages/d1/f9/b0c2ba758b14a7219e076138ea1e738c068bf388e64eee68f3df4fc96f5a/PyQt6-6.7.1.tar.gz"
	pyqt_cmd_tar := exec.Command("curl", "-L", pyqt_tar_url, "-o", "package.tar.gz")
	err := pyqt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyqt_zip_url := "https://files.pythonhosted.org/packages/d1/f9/b0c2ba758b14a7219e076138ea1e738c068bf388e64eee68f3df4fc96f5a/PyQt6-6.7.1.zip"
	pyqt_cmd_zip := exec.Command("curl", "-L", pyqt_zip_url, "-o", "package.zip")
	err = pyqt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyqt_bin_url := "https://files.pythonhosted.org/packages/d1/f9/b0c2ba758b14a7219e076138ea1e738c068bf388e64eee68f3df4fc96f5a/PyQt6-6.7.1.bin"
	pyqt_cmd_bin := exec.Command("curl", "-L", pyqt_bin_url, "-o", "binary.bin")
	err = pyqt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyqt_src_url := "https://files.pythonhosted.org/packages/d1/f9/b0c2ba758b14a7219e076138ea1e738c068bf388e64eee68f3df4fc96f5a/PyQt6-6.7.1.src.tar.gz"
	pyqt_cmd_src := exec.Command("curl", "-L", pyqt_src_url, "-o", "source.tar.gz")
	err = pyqt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyqt_cmd_direct := exec.Command("./binary")
	err = pyqt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyqt-builder")
	exec.Command("latte", "install", "pyqt-builder").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}
