package main

// Qsv - Ultra-fast CSV data-wrangling toolkit
// Homepage: https://github.com/jqnatividad/qsv

import (
	"fmt"
	
	"os/exec"
)

func installQsv() {
	// Método 1: Descargar y extraer .tar.gz
	qsv_tar_url := "https://github.com/jqnatividad/qsv/archive/refs/tags/0.133.1.tar.gz"
	qsv_cmd_tar := exec.Command("curl", "-L", qsv_tar_url, "-o", "package.tar.gz")
	err := qsv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qsv_zip_url := "https://github.com/jqnatividad/qsv/archive/refs/tags/0.133.1.zip"
	qsv_cmd_zip := exec.Command("curl", "-L", qsv_zip_url, "-o", "package.zip")
	err = qsv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qsv_bin_url := "https://github.com/jqnatividad/qsv/archive/refs/tags/0.133.1.bin"
	qsv_cmd_bin := exec.Command("curl", "-L", qsv_bin_url, "-o", "binary.bin")
	err = qsv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qsv_src_url := "https://github.com/jqnatividad/qsv/archive/refs/tags/0.133.1.src.tar.gz"
	qsv_cmd_src := exec.Command("curl", "-L", qsv_src_url, "-o", "source.tar.gz")
	err = qsv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qsv_cmd_direct := exec.Command("./binary")
	err = qsv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
}
