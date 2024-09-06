package main

// Pylyzer - Fast static code analyzer & language server for Python
// Homepage: https://github.com/mtshiba/pylyzer

import (
	"fmt"
	
	"os/exec"
)

func installPylyzer() {
	// Método 1: Descargar y extraer .tar.gz
	pylyzer_tar_url := "https://github.com/mtshiba/pylyzer/archive/refs/tags/v0.0.61.tar.gz"
	pylyzer_cmd_tar := exec.Command("curl", "-L", pylyzer_tar_url, "-o", "package.tar.gz")
	err := pylyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pylyzer_zip_url := "https://github.com/mtshiba/pylyzer/archive/refs/tags/v0.0.61.zip"
	pylyzer_cmd_zip := exec.Command("curl", "-L", pylyzer_zip_url, "-o", "package.zip")
	err = pylyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pylyzer_bin_url := "https://github.com/mtshiba/pylyzer/archive/refs/tags/v0.0.61.bin"
	pylyzer_cmd_bin := exec.Command("curl", "-L", pylyzer_bin_url, "-o", "binary.bin")
	err = pylyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pylyzer_src_url := "https://github.com/mtshiba/pylyzer/archive/refs/tags/v0.0.61.src.tar.gz"
	pylyzer_cmd_src := exec.Command("curl", "-L", pylyzer_src_url, "-o", "source.tar.gz")
	err = pylyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pylyzer_cmd_direct := exec.Command("./binary")
	err = pylyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
