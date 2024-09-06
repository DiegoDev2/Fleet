package main

// Emp - CLI for Empire
// Homepage: https://github.com/remind101/empire

import (
	"fmt"
	
	"os/exec"
)

func installEmp() {
	// Método 1: Descargar y extraer .tar.gz
	emp_tar_url := "https://github.com/remind101/empire/archive/refs/tags/v0.13.0.tar.gz"
	emp_cmd_tar := exec.Command("curl", "-L", emp_tar_url, "-o", "package.tar.gz")
	err := emp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emp_zip_url := "https://github.com/remind101/empire/archive/refs/tags/v0.13.0.zip"
	emp_cmd_zip := exec.Command("curl", "-L", emp_zip_url, "-o", "package.zip")
	err = emp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emp_bin_url := "https://github.com/remind101/empire/archive/refs/tags/v0.13.0.bin"
	emp_cmd_bin := exec.Command("curl", "-L", emp_bin_url, "-o", "binary.bin")
	err = emp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emp_src_url := "https://github.com/remind101/empire/archive/refs/tags/v0.13.0.src.tar.gz"
	emp_cmd_src := exec.Command("curl", "-L", emp_src_url, "-o", "source.tar.gz")
	err = emp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emp_cmd_direct := exec.Command("./binary")
	err = emp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
