package main

// PythonTkAT311 - Python interface to Tcl/Tk
// Homepage: https://www.python.org/

import (
	"fmt"
	
	"os/exec"
)

func installPythonTkAT311() {
	// Método 1: Descargar y extraer .tar.gz
	pythontkat311_tar_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythontkat311_cmd_tar := exec.Command("curl", "-L", pythontkat311_tar_url, "-o", "package.tar.gz")
	err := pythontkat311_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythontkat311_zip_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythontkat311_cmd_zip := exec.Command("curl", "-L", pythontkat311_zip_url, "-o", "package.zip")
	err = pythontkat311_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythontkat311_bin_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythontkat311_cmd_bin := exec.Command("curl", "-L", pythontkat311_bin_url, "-o", "binary.bin")
	err = pythontkat311_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythontkat311_src_url := "https://www.python.org/ftp/python/3.11.9/Python-3.11.9.tgz"
	pythontkat311_cmd_src := exec.Command("curl", "-L", pythontkat311_src_url, "-o", "source.tar.gz")
	err = pythontkat311_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythontkat311_cmd_direct := exec.Command("./binary")
	err = pythontkat311_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
}
