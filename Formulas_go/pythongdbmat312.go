package main

// PythonGdbmAT312 - Python interface to gdbm
// Homepage: https://www.python.org/

import (
	"fmt"
	
	"os/exec"
)

func installPythonGdbmAT312() {
	// Método 1: Descargar y extraer .tar.gz
	pythongdbmat312_tar_url := "https://www.python.org/ftp/python/3.12.5/Python-3.12.5.tgz"
	pythongdbmat312_cmd_tar := exec.Command("curl", "-L", pythongdbmat312_tar_url, "-o", "package.tar.gz")
	err := pythongdbmat312_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythongdbmat312_zip_url := "https://www.python.org/ftp/python/3.12.5/Python-3.12.5.tgz"
	pythongdbmat312_cmd_zip := exec.Command("curl", "-L", pythongdbmat312_zip_url, "-o", "package.zip")
	err = pythongdbmat312_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythongdbmat312_bin_url := "https://www.python.org/ftp/python/3.12.5/Python-3.12.5.tgz"
	pythongdbmat312_cmd_bin := exec.Command("curl", "-L", pythongdbmat312_bin_url, "-o", "binary.bin")
	err = pythongdbmat312_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythongdbmat312_src_url := "https://www.python.org/ftp/python/3.12.5/Python-3.12.5.tgz"
	pythongdbmat312_cmd_src := exec.Command("curl", "-L", pythongdbmat312_src_url, "-o", "source.tar.gz")
	err = pythongdbmat312_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythongdbmat312_cmd_direct := exec.Command("./binary")
	err = pythongdbmat312_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
