package main

// Fselect - Find files with SQL-like queries
// Homepage: https://github.com/jhspetersson/fselect

import (
	"fmt"
	
	"os/exec"
)

func installFselect() {
	// Método 1: Descargar y extraer .tar.gz
	fselect_tar_url := "https://github.com/jhspetersson/fselect/archive/refs/tags/0.8.6.tar.gz"
	fselect_cmd_tar := exec.Command("curl", "-L", fselect_tar_url, "-o", "package.tar.gz")
	err := fselect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fselect_zip_url := "https://github.com/jhspetersson/fselect/archive/refs/tags/0.8.6.zip"
	fselect_cmd_zip := exec.Command("curl", "-L", fselect_zip_url, "-o", "package.zip")
	err = fselect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fselect_bin_url := "https://github.com/jhspetersson/fselect/archive/refs/tags/0.8.6.bin"
	fselect_cmd_bin := exec.Command("curl", "-L", fselect_bin_url, "-o", "binary.bin")
	err = fselect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fselect_src_url := "https://github.com/jhspetersson/fselect/archive/refs/tags/0.8.6.src.tar.gz"
	fselect_cmd_src := exec.Command("curl", "-L", fselect_src_url, "-o", "source.tar.gz")
	err = fselect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fselect_cmd_direct := exec.Command("./binary")
	err = fselect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
