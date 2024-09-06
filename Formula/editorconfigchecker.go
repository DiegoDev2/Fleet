package main

// EditorconfigChecker - Tool to verify that your files are in harmony with your .editorconfig
// Homepage: https://github.com/editorconfig-checker/editorconfig-checker

import (
	"fmt"
	
	"os/exec"
)

func installEditorconfigChecker() {
	// Método 1: Descargar y extraer .tar.gz
	editorconfigchecker_tar_url := "https://github.com/editorconfig-checker/editorconfig-checker/archive/refs/tags/v3.0.3.tar.gz"
	editorconfigchecker_cmd_tar := exec.Command("curl", "-L", editorconfigchecker_tar_url, "-o", "package.tar.gz")
	err := editorconfigchecker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	editorconfigchecker_zip_url := "https://github.com/editorconfig-checker/editorconfig-checker/archive/refs/tags/v3.0.3.zip"
	editorconfigchecker_cmd_zip := exec.Command("curl", "-L", editorconfigchecker_zip_url, "-o", "package.zip")
	err = editorconfigchecker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	editorconfigchecker_bin_url := "https://github.com/editorconfig-checker/editorconfig-checker/archive/refs/tags/v3.0.3.bin"
	editorconfigchecker_cmd_bin := exec.Command("curl", "-L", editorconfigchecker_bin_url, "-o", "binary.bin")
	err = editorconfigchecker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	editorconfigchecker_src_url := "https://github.com/editorconfig-checker/editorconfig-checker/archive/refs/tags/v3.0.3.src.tar.gz"
	editorconfigchecker_cmd_src := exec.Command("curl", "-L", editorconfigchecker_src_url, "-o", "source.tar.gz")
	err = editorconfigchecker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	editorconfigchecker_cmd_direct := exec.Command("./binary")
	err = editorconfigchecker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
