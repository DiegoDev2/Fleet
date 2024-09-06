package main

// Datatype99 - Algebraic data types for C99
// Homepage: https://github.com/Hirrolot/datatype99

import (
	"fmt"
	
	"os/exec"
)

func installDatatype99() {
	// Método 1: Descargar y extraer .tar.gz
	datatype99_tar_url := "https://github.com/Hirrolot/datatype99/archive/refs/tags/v1.6.4.tar.gz"
	datatype99_cmd_tar := exec.Command("curl", "-L", datatype99_tar_url, "-o", "package.tar.gz")
	err := datatype99_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	datatype99_zip_url := "https://github.com/Hirrolot/datatype99/archive/refs/tags/v1.6.4.zip"
	datatype99_cmd_zip := exec.Command("curl", "-L", datatype99_zip_url, "-o", "package.zip")
	err = datatype99_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	datatype99_bin_url := "https://github.com/Hirrolot/datatype99/archive/refs/tags/v1.6.4.bin"
	datatype99_cmd_bin := exec.Command("curl", "-L", datatype99_bin_url, "-o", "binary.bin")
	err = datatype99_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	datatype99_src_url := "https://github.com/Hirrolot/datatype99/archive/refs/tags/v1.6.4.src.tar.gz"
	datatype99_cmd_src := exec.Command("curl", "-L", datatype99_src_url, "-o", "source.tar.gz")
	err = datatype99_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	datatype99_cmd_direct := exec.Command("./binary")
	err = datatype99_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: metalang99")
	exec.Command("latte", "install", "metalang99").Run()
}
