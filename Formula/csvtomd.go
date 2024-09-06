package main

// Csvtomd - CSV to Markdown table converter
// Homepage: https://github.com/mplewis/csvtomd

import (
	"fmt"
	
	"os/exec"
)

func installCsvtomd() {
	// Método 1: Descargar y extraer .tar.gz
	csvtomd_tar_url := "https://files.pythonhosted.org/packages/9d/59/ea3c8b102f9c72e5d276a169f7f343432213441c39a6eac7a8f444c66681/csvtomd-0.3.0.tar.gz"
	csvtomd_cmd_tar := exec.Command("curl", "-L", csvtomd_tar_url, "-o", "package.tar.gz")
	err := csvtomd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csvtomd_zip_url := "https://files.pythonhosted.org/packages/9d/59/ea3c8b102f9c72e5d276a169f7f343432213441c39a6eac7a8f444c66681/csvtomd-0.3.0.zip"
	csvtomd_cmd_zip := exec.Command("curl", "-L", csvtomd_zip_url, "-o", "package.zip")
	err = csvtomd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csvtomd_bin_url := "https://files.pythonhosted.org/packages/9d/59/ea3c8b102f9c72e5d276a169f7f343432213441c39a6eac7a8f444c66681/csvtomd-0.3.0.bin"
	csvtomd_cmd_bin := exec.Command("curl", "-L", csvtomd_bin_url, "-o", "binary.bin")
	err = csvtomd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csvtomd_src_url := "https://files.pythonhosted.org/packages/9d/59/ea3c8b102f9c72e5d276a169f7f343432213441c39a6eac7a8f444c66681/csvtomd-0.3.0.src.tar.gz"
	csvtomd_cmd_src := exec.Command("curl", "-L", csvtomd_src_url, "-o", "source.tar.gz")
	err = csvtomd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csvtomd_cmd_direct := exec.Command("./binary")
	err = csvtomd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
