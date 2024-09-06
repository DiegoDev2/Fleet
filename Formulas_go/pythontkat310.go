package main

// PythonTkAT310 - Python interface to Tcl/Tk
// Homepage: https://www.python.org/

import (
	"fmt"
	
	"os/exec"
)

func installPythonTkAT310() {
	// Método 1: Descargar y extraer .tar.gz
	pythontkat310_tar_url := "https://www.python.org/ftp/python/3.10.14/Python-3.10.14.tgz"
	pythontkat310_cmd_tar := exec.Command("curl", "-L", pythontkat310_tar_url, "-o", "package.tar.gz")
	err := pythontkat310_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythontkat310_zip_url := "https://www.python.org/ftp/python/3.10.14/Python-3.10.14.tgz"
	pythontkat310_cmd_zip := exec.Command("curl", "-L", pythontkat310_zip_url, "-o", "package.zip")
	err = pythontkat310_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythontkat310_bin_url := "https://www.python.org/ftp/python/3.10.14/Python-3.10.14.tgz"
	pythontkat310_cmd_bin := exec.Command("curl", "-L", pythontkat310_bin_url, "-o", "binary.bin")
	err = pythontkat310_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythontkat310_src_url := "https://www.python.org/ftp/python/3.10.14/Python-3.10.14.tgz"
	pythontkat310_cmd_src := exec.Command("curl", "-L", pythontkat310_src_url, "-o", "source.tar.gz")
	err = pythontkat310_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythontkat310_cmd_direct := exec.Command("./binary")
	err = pythontkat310_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.10")
exec.Command("latte", "install", "python@3.10")
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
}
