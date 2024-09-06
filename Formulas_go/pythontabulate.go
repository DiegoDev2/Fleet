package main

// PythonTabulate - Pretty-print tabular data in Python
// Homepage: https://github.com/astanin/python-tabulate

import (
	"fmt"
	
	"os/exec"
)

func installPythonTabulate() {
	// Método 1: Descargar y extraer .tar.gz
	pythontabulate_tar_url := "https://files.pythonhosted.org/packages/ec/fe/802052aecb21e3797b8f7902564ab6ea0d60ff8ca23952079064155d1ae1/tabulate-0.9.0.tar.gz"
	pythontabulate_cmd_tar := exec.Command("curl", "-L", pythontabulate_tar_url, "-o", "package.tar.gz")
	err := pythontabulate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythontabulate_zip_url := "https://files.pythonhosted.org/packages/ec/fe/802052aecb21e3797b8f7902564ab6ea0d60ff8ca23952079064155d1ae1/tabulate-0.9.0.zip"
	pythontabulate_cmd_zip := exec.Command("curl", "-L", pythontabulate_zip_url, "-o", "package.zip")
	err = pythontabulate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythontabulate_bin_url := "https://files.pythonhosted.org/packages/ec/fe/802052aecb21e3797b8f7902564ab6ea0d60ff8ca23952079064155d1ae1/tabulate-0.9.0.bin"
	pythontabulate_cmd_bin := exec.Command("curl", "-L", pythontabulate_bin_url, "-o", "binary.bin")
	err = pythontabulate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythontabulate_src_url := "https://files.pythonhosted.org/packages/ec/fe/802052aecb21e3797b8f7902564ab6ea0d60ff8ca23952079064155d1ae1/tabulate-0.9.0.src.tar.gz"
	pythontabulate_cmd_src := exec.Command("curl", "-L", pythontabulate_src_url, "-o", "source.tar.gz")
	err = pythontabulate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythontabulate_cmd_direct := exec.Command("./binary")
	err = pythontabulate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
