package main

// PyqtAT5 - Python bindings for v5 of Qt
// Homepage: https://www.riverbankcomputing.com/software/pyqt/intro

import (
	"fmt"
	
	"os/exec"
)

func installPyqtAT5() {
	// Método 1: Descargar y extraer .tar.gz
	pyqtat5_tar_url := "https://files.pythonhosted.org/packages/4d/5d/b8b6e26956ec113ad3f566e02abd12ac3a56b103fcc7e0735e27ee4a1df3/PyQt5-5.15.10.tar.gz"
	pyqtat5_cmd_tar := exec.Command("curl", "-L", pyqtat5_tar_url, "-o", "package.tar.gz")
	err := pyqtat5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyqtat5_zip_url := "https://files.pythonhosted.org/packages/4d/5d/b8b6e26956ec113ad3f566e02abd12ac3a56b103fcc7e0735e27ee4a1df3/PyQt5-5.15.10.zip"
	pyqtat5_cmd_zip := exec.Command("curl", "-L", pyqtat5_zip_url, "-o", "package.zip")
	err = pyqtat5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyqtat5_bin_url := "https://files.pythonhosted.org/packages/4d/5d/b8b6e26956ec113ad3f566e02abd12ac3a56b103fcc7e0735e27ee4a1df3/PyQt5-5.15.10.bin"
	pyqtat5_cmd_bin := exec.Command("curl", "-L", pyqtat5_bin_url, "-o", "binary.bin")
	err = pyqtat5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyqtat5_src_url := "https://files.pythonhosted.org/packages/4d/5d/b8b6e26956ec113ad3f566e02abd12ac3a56b103fcc7e0735e27ee4a1df3/PyQt5-5.15.10.src.tar.gz"
	pyqtat5_cmd_src := exec.Command("curl", "-L", pyqtat5_src_url, "-o", "source.tar.gz")
	err = pyqtat5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyqtat5_cmd_direct := exec.Command("./binary")
	err = pyqtat5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyqt-builder")
exec.Command("latte", "install", "pyqt-builder")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
