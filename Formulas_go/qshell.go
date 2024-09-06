package main

// Qshell - Shell Tools for Qiniu Cloud
// Homepage: https://github.com/qiniu/qshell

import (
	"fmt"
	
	"os/exec"
)

func installQshell() {
	// Método 1: Descargar y extraer .tar.gz
	qshell_tar_url := "https://github.com/qiniu/qshell/archive/refs/tags/v2.14.0.tar.gz"
	qshell_cmd_tar := exec.Command("curl", "-L", qshell_tar_url, "-o", "package.tar.gz")
	err := qshell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qshell_zip_url := "https://github.com/qiniu/qshell/archive/refs/tags/v2.14.0.zip"
	qshell_cmd_zip := exec.Command("curl", "-L", qshell_zip_url, "-o", "package.zip")
	err = qshell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qshell_bin_url := "https://github.com/qiniu/qshell/archive/refs/tags/v2.14.0.bin"
	qshell_cmd_bin := exec.Command("curl", "-L", qshell_bin_url, "-o", "binary.bin")
	err = qshell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qshell_src_url := "https://github.com/qiniu/qshell/archive/refs/tags/v2.14.0.src.tar.gz"
	qshell_cmd_src := exec.Command("curl", "-L", qshell_src_url, "-o", "source.tar.gz")
	err = qshell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qshell_cmd_direct := exec.Command("./binary")
	err = qshell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
