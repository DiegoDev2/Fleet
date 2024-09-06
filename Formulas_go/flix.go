package main

// Flix - Statically typed functional, imperative, and logic programming language
// Homepage: https://flix.dev/

import (
	"fmt"
	
	"os/exec"
)

func installFlix() {
	// Método 1: Descargar y extraer .tar.gz
	flix_tar_url := "https://github.com/flix/flix/archive/refs/tags/v0.50.0.tar.gz"
	flix_cmd_tar := exec.Command("curl", "-L", flix_tar_url, "-o", "package.tar.gz")
	err := flix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flix_zip_url := "https://github.com/flix/flix/archive/refs/tags/v0.50.0.zip"
	flix_cmd_zip := exec.Command("curl", "-L", flix_zip_url, "-o", "package.zip")
	err = flix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flix_bin_url := "https://github.com/flix/flix/archive/refs/tags/v0.50.0.bin"
	flix_cmd_bin := exec.Command("curl", "-L", flix_bin_url, "-o", "binary.bin")
	err = flix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flix_src_url := "https://github.com/flix/flix/archive/refs/tags/v0.50.0.src.tar.gz"
	flix_cmd_src := exec.Command("curl", "-L", flix_src_url, "-o", "source.tar.gz")
	err = flix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flix_cmd_direct := exec.Command("./binary")
	err = flix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
exec.Command("latte", "install", "gradle")
	fmt.Println("Instalando dependencia: scala")
exec.Command("latte", "install", "scala")
	fmt.Println("Instalando dependencia: openjdk@21")
exec.Command("latte", "install", "openjdk@21")
}
